package servant

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"os"
	"runtime"
	"time"

	"tarsrpc/jce/taf"
	"tarsrpc/jce_parser/gojce"
	"tarsrpc/kissgo/appzaplog"
	"tarsrpc/kissgo/appzaplog/zap"
	"tarsrpc/tars/servant/warn"
)

type Dispatcher interface {
	Dispatch(context.Context, interface{}, *taf.RequestPacket) (*taf.ResponsePacket, error)
}

type JceProtocol struct {
	dispatcher Dispatcher
	serverImp  interface{}
}

func NewJceProtocol(dispatcher Dispatcher, imp interface{}) *JceProtocol {
	s := &JceProtocol{dispatcher: dispatcher, serverImp: imp}
	return s
}

const panicWarnTpl = `<font color="warning">panic</font>
>time:<font color="comment">%s</font>
>env:<font color="comment">%s</font>
>server:<font color="comment">%s</font>
>function:<font color="comment">%s</font>
>ip:<font color="comment">%s</font>
>panic:<font color="comment">%v</font>
>stack:<font color="comment">%s</font>`

func (s *JceProtocol) doDispatch(ctx context.Context, reqPackage *taf.RequestPacket) (rspPackage *taf.ResponsePacket, err error) {
	defer func() {
		if r := recover(); r != nil {
			env := os.Getenv("YUNGAME_ENV")
			_ = warn.ServerAlarm(fmt.Sprintf(panicWarnTpl, time.Now().Format("2006-01-02 15:04:05"), env,
				reqPackage.SServantName, reqPackage.SFuncName, warn.GetIp(), r, getStackInfo()))
			appzaplog.DPanic("doDispatch handler panic", zap.Any("panic", r))
			err = HandlerPanicTarErr
		}
	}()
	ctx = NewOutgoingContext(ctx, reqPackage.Context)

	rspPackage, err = s.dispatcher.Dispatch(ctx, s.serverImp, reqPackage)
	return
}

// getStackInfo 打印Panic堆栈信息
func getStackInfo() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return fmt.Sprintf("%s", buf[:n])
}

func (s *JceProtocol) Invoke(ctx context.Context, req []byte) (rsp []byte, err error) {
	defer checkPanic()
	var (
		reqPackage taf.RequestPacket
		rspPackage *taf.ResponsePacket
		//now  = time.Now()
	)

	is := gojce.NewInputStream(req)
	if err = reqPackage.ReadFrom(is); err != nil {
		//this will close the connection
		appzaplog.Error("Invoke ReadFrom reqPackage failed", zap.Error(err))
		return
	}

	appzaplog.Debug("[+]Invoke",
		zap.String("SFuncName", reqPackage.SFuncName),
		zap.String("SServantName", reqPackage.SServantName),
		zap.Int32("IRequestId", reqPackage.IRequestId))

	for {
		if callDisabled(reqPackage.SServantName + "." + reqPackage.SFuncName) {
			err = RPCCallDisabledTarErr
			break
		}
		rspPackage, err = s.doDispatch(ctx, &reqPackage)
		break
	}

	// TarError report here
	//if err == nil{
	//	httpmetrics.DefReport(reqPackage.SFuncName,0,now,httpmetrics.DefaultSuccessFun)
	//}else if merr,ok := err.(TarError);ok{
	//	httpmetrics.DefReport(reqPackage.SFuncName,merr.Code,now,httpmetrics.DefaultSuccessFun)
	//}

	// return err to client
	if err != nil {
		rspPackage = &taf.ResponsePacket{
			IVersion:   1,
			IRequestId: reqPackage.IRequestId,
			SBuffer:    nil,
			Context:    reqPackage.Context,
			Status:     map[string]string{STATUSERRSTR: err.Error()},
		}
		err = nil
	}
	os := gojce.NewOutputStream()
	rspPackage.WriteTo(os)
	bs := os.ToBytes()
	sbuf := bytes.NewBuffer(nil)
	sbuf.Write(make([]byte, 4))
	sbuf.Write(bs)
	binary.BigEndian.PutUint32(sbuf.Bytes(), uint32(sbuf.Len()))
	rsp = sbuf.Bytes()
	appzaplog.Debug("[-]Invoke",
		zap.String("SFuncName", reqPackage.SFuncName),
		zap.String("SServantName", reqPackage.SServantName),
		zap.Int32("IRequestId", reqPackage.IRequestId))
	return
}

func (s *JceProtocol) ParsePackage(buff []byte) (int, int) {
	return TafRequest(buff)
}

func (s *JceProtocol) InvokeTimeout(pkg []byte) (resp []byte, err error) {
	//TODO ERROR PACKAGE
	defer checkPanic()
	var (
		reqPackage taf.RequestPacket
		rspPackage *taf.ResponsePacket
	)
	appzaplog.Error("[+]InvokeTimeout", zap.Binary("pkg", pkg))
	is := gojce.NewInputStream(pkg)
	if err = reqPackage.ReadFrom(is); err != nil {
		//this will close the connection
		appzaplog.Error("InvokeTimeout ReadFrom reqPackage failed", zap.Error(err))
		return
	}
	rspPackage = &taf.ResponsePacket{
		IVersion:   1,
		IRequestId: reqPackage.IRequestId,
		SBuffer:    nil,
		Context:    reqPackage.Context,
		Status:     map[string]string{STATUSERRSTR: "rpc timeout"},
	}
	os := gojce.NewOutputStream()
	rspPackage.WriteTo(os)
	bs := os.ToBytes()
	sbuf := bytes.NewBuffer(nil)
	sbuf.Write(make([]byte, 4))
	sbuf.Write(bs)
	len := sbuf.Len()
	binary.BigEndian.PutUint32(sbuf.Bytes(), uint32(len))
	return sbuf.Bytes(), nil
	//payload := []byte("timeout")
	//ret := make([]byte, 4+len(payload))
	//binary.BigEndian.PutUint32(pkg[:4], uint32(len(ret)))
	//copy(pkg[4:], payload)
	//return ret, nil
}
