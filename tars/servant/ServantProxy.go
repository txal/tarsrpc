package servant

import (
	"context"
	"errors"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"code.com/tars/goframework/jce/taf"
	"code.com/tars/goframework/kissgo/appzaplog"
	"code.com/tars/goframework/kissgo/appzaplog/zap"
	pbtaf "code.com/tars/goframework/tars/servant/protocol"
	"code.com/tars/goframework/tars/util/current"
)

type ServantProxy struct {
	sid     int32
	name    string // appname.servername.objname
	comm    *Communicator
	obj     *ObjectProxy
	timeout int
}

func NewServantProxy(comm *Communicator, objName string) *ServantProxy {
	s := &ServantProxy{
		comm:    comm,
		timeout: comm.Client.syncInvokeTimeout,
	}
	pos := strings.Index(objName, "@")
	if pos > 0 {
		s.name = objName[0:pos]
	} else {
		s.name = objName
	}
	s.obj = NewObjectProxy(comm, objName)
	return s
}

func (s *ServantProxy) Taf_invoke(ctx context.Context, ctype byte, sFuncName string,
	buf []byte) (*taf.ResponsePacket, error) {
	//TODO 重置sid，防止溢出
	atomic.CompareAndSwapInt32(&s.sid, 1<<31-1, 1)
	ctxmap, _ := FromOutgoingContext(ctx)

	req := taf.RequestPacket{
		IVersion:     1,
		CPacketType:  ctype,
		IRequestId:   atomic.AddInt32(&s.sid, 1),
		SServantName: s.name,
		SFuncName:    sFuncName,
		SBuffer:      buf,
		ITimeout:     int32(s.timeout),
		Context:      ctxmap,
	}

	msg := &Message{Req: &req, Ser: s, Obj: s.obj}
	msg.Init()

	if key, ok := ctxmap[CONTEXTCONSISTHASHKEY]; ok {
		msg.setConsistHashCode(key)
	}

	invokeTimeout := time.Duration(s.timeout) * time.Millisecond
	ok, to, isTimeout := current.GetClientTimeout(ctx)
	if ok && isTimeout {
		invokeTimeout = time.Duration(to) * time.Millisecond
	}

	defer func() {
		msg.End()
		ReportStat(msg)
	}()
	appzaplog.Debug("Taf_invoke", zap.String("sFuncName", sFuncName),
		zap.String("obj", s.name), zap.Int32("IRequestId", req.IRequestId))
	if err := s.obj.Invoke(ctx, msg, invokeTimeout); err != nil {
		appzaplog.Error("Invoke error", zap.String("ServantName", s.name),
			zap.String("FuncName", sFuncName), zap.Int32("IRequestId", req.IRequestId), zap.Error(err))
		return nil, err
	}

	if msg.Resp != nil {
		if errstr, ok := msg.Resp.Status[STATUSERRSTR]; ok {
			return msg.Resp, errors.New(errstr)
		}
	}
	return msg.Resp, nil
}

func (s *ServantProxy) Pb_invoke(ctx context.Context, ctype byte, sFuncName string,
	buf []byte, status map[string]string, context map[string]string) (*pbtaf.ResponsePacket, error) {
	//TODO 重置sid，防止溢出
	atomic.CompareAndSwapInt32(&s.sid, 1<<31-1, 1)
	req := pbtaf.RequestPacket{
		IVersion:     1,
		CPacketType:  pbtaf.RequestPacket_PacketType(ctype),
		IRequestId:   atomic.AddInt32(&s.sid, 1),
		SServantName: s.name,
		SFuncName:    sFuncName,
		SBuffer:      buf,
		ITimeout:     3000,
		Context:      context,
		Status:       status,
	}

	msg := &PbMessage{Req: &req, Ser: s, Obj: s.obj}
	msg.Init()

	defer func() {
		msg.End()
		ReportStat(msg)
	}()
	appzaplog.Debug("Pb_invoke", zap.String("sFuncName", sFuncName),
		zap.String("obj", s.name), zap.Int32("IRequestId", req.IRequestId))
	if err := s.obj.PbInvoke(msg); err != nil {
		appzaplog.Error("Pb_invoke error", zap.String("ServantName", s.name),
			zap.String("FuncName", sFuncName),
			zap.Int32("IRequestId", req.IRequestId),
			zap.Error(err))
		return nil, err
	}

	return msg.Resp, nil
}

type ServantProxyFactory struct {
	lk   sync.RWMutex
	objs map[string]*ServantProxy
	comm *Communicator
}

func NewServantProxyFactory(comm *Communicator) *ServantProxyFactory {
	return &ServantProxyFactory{
		comm: comm,
		objs: make(map[string]*ServantProxy),
	}
}

func (o *ServantProxyFactory) getServantProxy(objName string) *ServantProxy {
	proxy := o.getProxy(objName)
	if proxy != nil {
		return proxy
	}
	return o.createProxy(objName)
}

func (o *ServantProxyFactory) getProxy(objName string) *ServantProxy {
	o.lk.RLock()
	defer o.lk.RUnlock()
	if obj, ok := o.objs[objName]; ok {
		return obj
	}
	return nil
}

func (o *ServantProxyFactory) createProxy(objName string) *ServantProxy {
	o.lk.Lock()
	defer o.lk.Unlock()
	if obj, ok := o.objs[objName]; ok {
		return obj
	}
	obj := NewServantProxy(o.comm, objName)
	o.objs[objName] = obj
	return obj
}

func (s *ServantProxy) GetAdapterProxys() map[string]*AdapterProxy {
	return s.obj.GetAvailableProxys()
}

func (s *ServantProxy) GetProxyEndPoints() []string {
	ipPorts := []string{}
	adps := s.obj.GetAvailableProxys()
	for ipPort, _ := range adps {
		ipPorts = append(ipPorts, ipPort)
	}
	return ipPorts
}

func (s *ServantProxy) Proxy_invoke(ctx context.Context, ctype byte, sFuncName string,
	buf []byte, ipPort string) (*taf.ResponsePacket, error) {

	adps := s.GetAdapterProxys()
	adp := adps[ipPort]

	//TODO 重置sid，防止溢出
	atomic.CompareAndSwapInt32(&s.sid, 1<<31-1, 1)
	ctxmap, _ := FromOutgoingContext(ctx)

	//appzaplog.Debug("ctxmap info",zap.Any("ctxmap",ctxmap))
	req := taf.RequestPacket{
		IVersion:     1,
		CPacketType:  ctype,
		IRequestId:   atomic.AddInt32(&s.sid, 1),
		SServantName: s.name,
		SFuncName:    sFuncName,
		SBuffer:      buf,
		ITimeout:     3000,
		Context:      ctxmap,
		//Status:       statusmap,
	}

	msg := &Message{Req: &req, Ser: s, Obj: s.obj, Adp: adp}
	if key, ok := ctxmap[CONTEXTCONSISTHASHKEY]; ok {
		//appzaplog.Debug("consisthashkey set",zap.String("consisthashkey",key))
		msg.setConsistHashCode(key)
	}
	msg.Init()

	defer func() {
		msg.End()
		ReportStat(msg)
	}()
	appzaplog.Debug("Proxy_invoke", zap.String("sFuncName", sFuncName),
		zap.String("obj", s.name), zap.Int32("IRequestId", req.IRequestId))
	if err := s.obj.ProxyInvoke(ctx, msg); err != nil {
		appzaplog.Error("Proxy_invoke error", zap.String("ServantName", s.name),
			zap.String("FuncName", sFuncName), zap.Int32("IRequestId", req.IRequestId), zap.Error(err))
		return nil, err
	}

	if msg.Resp != nil {
		if errstr, ok := msg.Resp.Status[STATUSERRSTR]; ok {
			return msg.Resp, errors.New(errstr)
		}
	}
	return msg.Resp, nil
}
