package servant

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"code.com/tars/goframework/tars/servant/warn"

	base "code.com/tars/goframework/jce/servant/taf"
	"code.com/tars/goframework/jce/taf"
	"code.com/tars/goframework/kissgo/appzaplog"
	"code.com/tars/goframework/kissgo/appzaplog/zap"
	pbtaf "code.com/tars/goframework/tars/servant/protocol"
	"code.com/tars/goframework/tars/tarsserver"
)

type ObjectProxy struct {
	manager  *EndpointManager
	comm     *Communicator
	queueLen int32
}

func NewObjectProxy(comm *Communicator, objName string) *ObjectProxy {
	return &ObjectProxy{
		comm:    comm,
		manager: NewEndpointManager(objName, comm),
	}
}

func (obj *ObjectProxy) Invoke(ctx context.Context, msg *Message, timeout time.Duration) error {
	adp := obj.manager.SelectAdapterProxy(msg)
	msg.Adp = adp
	if adp == nil {
		msg.Resp = &taf.ResponsePacket{
			IRet:        base.JCEADAPTERNULL,
			SResultDesc: "no adapter proxy selected",
		}
		warn.ServerAlarm(fmt.Sprintf("no adapter proxy selected, ip:%s, obj:%s, interface:%s", warn.GetIp(), msg.Req.SServantName, msg.Req.SFuncName))
		return NoAdapterErr
	}
	if atomic.LoadInt32(&obj.queueLen) > 10000 {
		msg.Resp = &taf.ResponsePacket{
			IRet:        base.JCESERVEROVERLOAD,
			SResultDesc: "invoke queue is full",
		}
		return OverloadErr
	}
	atomic.AddInt32(&obj.queueLen, 1)
	readCh := make(chan *taf.ResponsePacket, 1)
	//todo move to TarsClientProtocol
	adp.resp.Store(msg.Req.IRequestId, readCh)
	ctx, cancle := context.WithTimeout(ctx, timeout)
	defer func() {
		checkPanic()
		atomic.AddInt32(&obj.queueLen, -1)
		adp.resp.Delete(msg.Req.IRequestId)
		close(readCh)
		cancle()
	}()

	_, err := adp.circutbreaker.Execute(func() (interface{}, error) {
		if err := adp.Send(msg.Req); err != nil {
			appzaplog.Error("Send msg failed", zap.Error(err),
				zap.String("ipport", adp.point.IpPort()),
				zap.Int32("RequestId", msg.Req.IRequestId),
				zap.String("servant", msg.Req.SServantName),
				zap.String("func", msg.Req.SFuncName))
			return nil, SendErr
		}

		select {
		case <-ctx.Done():
			appzaplog.Warn("req timeout",
				zap.String("ipport", adp.point.IpPort()),
				zap.Int32("RequestId", msg.Req.IRequestId),
				zap.String("servant", msg.Req.SServantName),
				zap.String("func", msg.Req.SFuncName))

			msg.Resp = &taf.ResponsePacket{
				IRet:        base.JCEINVOKETIMEOUT,
				SResultDesc: "req timeout",
			}
			return nil, ReqTimeoutErr
		case msg.Resp = <-readCh:
			appzaplog.Debug("recv msg succ ", zap.Int32("RequestId", msg.Req.IRequestId))
		}
		return nil, nil
	})

	return err
}

func (obj *ObjectProxy) PbInvoke(msg *PbMessage) error {
	//now := time.Now()
	adp := obj.manager.SelectAdapterProxy(msg)
	msg.Adp = adp
	if adp == nil {
		msg.Resp = &pbtaf.ResponsePacket{
			IRet:        base.JCEADAPTERNULL,
			SResultDesc: "no adapter Proxy selected",
		}
		warn.ServerAlarm(fmt.Sprintf("no adapter proxy selected, ip:%s, obj:%s, interface:%s", warn.GetIp(), msg.Req.SServantName, msg.Req.SFuncName))
		return NoAdapterErr
	}
	if atomic.LoadInt32(&obj.queueLen) > 10000 {
		msg.Resp = &pbtaf.ResponsePacket{
			IRet:        base.JCESERVEROVERLOAD,
			SResultDesc: "invoke queue is full",
		}
		return OverloadErr
	}
	atomic.AddInt32(&obj.queueLen, 1)
	readCh := make(chan *pbtaf.ResponsePacket, 1)
	adp.resp.Store(msg.Req.IRequestId, readCh)
	defer func() {
		checkPanic()
		atomic.AddInt32(&obj.queueLen, -1)
		adp.resp.Delete(msg.Req.IRequestId)
		close(readCh)
	}()
	//TODO adp active check
	_, err := adp.circutbreaker.Execute(func() (interface{}, error) {
		if err := adp.PbSend(msg.Req); err != nil {
			//adp.failAdd()
			appzaplog.Error("Send msg failed", zap.Error(err))
			if err == tarsserver.NetDialTimeoutErr && obj.manager != nil {
				// need to refresh endpoint cache with ratelimit
				//obj.manager.findAndSetObj(startFrameWorkComm().sd)
			}
			return nil, SendErr
		}
		//httpmetrics.DefReport(msg.Req.SFuncName, 0, now)
		select {
		//TODO USE TIMEOUT
		case <-time.After(3 * time.Second):
			appzaplog.Warn("req timeout", zap.Int32("RequestId", msg.Req.IRequestId))

			msg.Resp = &pbtaf.ResponsePacket{
				IRet:        base.JCEINVOKETIMEOUT,
				SResultDesc: "req timeout",
			}
			//adp.failAdd()
			return nil, ReqTimeoutErr
		case msg.Resp = <-readCh:
			appzaplog.Debug("recv msg succ ", zap.Int32("RequestId", msg.Req.IRequestId))
		}
		return nil, nil
	})
	return err
}

func (obj *ObjectProxy) GetAvailableProxys() map[string]*AdapterProxy {
	return obj.manager.GetAvailableProxys()
}

type ObjectProxyFactory struct {
	objs map[string]*ObjectProxy
	comm *Communicator
	om   *sync.Mutex
}

//func NewObjectProxyFactory(comm *Communicator) *ObjectProxyFactory {
//	o := &ObjectProxyFactory{
//		om:   new(sync.Mutex),
//		comm: comm,
//		objs: make(map[string]*ObjectProxy),
//	}
//	return o
//}
//
//func (o *ObjectProxyFactory) GetObjectProxy(objName string) *ObjectProxy {
//	if obj, ok := o.objs[objName]; ok {
//		return obj
//	}
//	obj := NewObjectProxy(o.comm, objName)
//
//	o.om.Lock()
//	defer o.om.Unlock()
//	o.objs[objName] = obj
//	return obj
//}

func (obj *ObjectProxy) ProxyInvoke(ctx context.Context, msg *Message) error {
	adp := obj.manager.SelectAdapterProxy(msg)

	if msg.Adp != nil {
		adp = msg.Adp
	} else {
		msg.Adp = adp
	}

	if adp == nil {
		msg.Resp = &taf.ResponsePacket{
			IRet:        base.JCEADAPTERNULL,
			SResultDesc: "no adapter proxy selected",
		}
		warn.ServerAlarm(fmt.Sprintf("no adapter proxy selected, ip:%s, obj:%s, interface:%s", warn.GetIp(), msg.Req.SServantName, msg.Req.SFuncName))
		return NoAdapterErr
	}
	if atomic.LoadInt32(&obj.queueLen) > 10000 {
		msg.Resp = &taf.ResponsePacket{
			IRet:        base.JCESERVEROVERLOAD,
			SResultDesc: "invoke queue is full",
		}
		return OverloadErr
	}
	atomic.AddInt32(&obj.queueLen, 1)
	readCh := make(chan *taf.ResponsePacket, 1)
	//todo move to TarsClientProtocol
	adp.resp.Store(msg.Req.IRequestId, readCh)
	ctx, cancle := context.WithTimeout(ctx, takeclientsynctimeinvoketimeout())
	defer func() {
		checkPanic()
		atomic.AddInt32(&obj.queueLen, -1)
		adp.resp.Delete(msg.Req.IRequestId)
		close(readCh)
		cancle()
	}()

	// 如果是老逻辑服务则不走熔断插件
	if msg.Req.SServantName == "breath.videonode.NodeTarsObj" {
		appzaplog.Debug("no run circutbreaker", zap.String("obj", msg.Req.SServantName))
		_, err := func() (interface{}, error) {
			if err := adp.Send(msg.Req); err != nil {
				appzaplog.Error("Send msg failed", zap.Error(err))
				return nil, SendErr
			}

			select {
			case <-ctx.Done():
				appzaplog.Warn("req timeout", zap.Int32("RequestId", msg.Req.IRequestId))

				msg.Resp = &taf.ResponsePacket{
					IRet:        base.JCEINVOKETIMEOUT,
					SResultDesc: "req timeout",
				}
				return nil, ReqTimeoutErr
			case msg.Resp = <-readCh:
				appzaplog.Debug("recv msg succ ", zap.Int32("RequestId", msg.Req.IRequestId))
			}
			return nil, nil
		}()
		return err
	} else {
		_, err := adp.circutbreaker.Execute(func() (interface{}, error) {
			if err := adp.Send(msg.Req); err != nil {
				appzaplog.Error("Send msg failed", zap.Error(err))
				return nil, SendErr
			}

			select {
			case <-ctx.Done():
				appzaplog.Warn("req timeout", zap.Int32("RequestId", msg.Req.IRequestId))

				msg.Resp = &taf.ResponsePacket{
					IRet:        base.JCEINVOKETIMEOUT,
					SResultDesc: "req timeout",
				}
				return nil, ReqTimeoutErr
			case msg.Resp = <-readCh:
				appzaplog.Debug("recv msg succ ", zap.Int32("RequestId", msg.Req.IRequestId))
			}
			return nil, nil
		})
		return err
	}
}
