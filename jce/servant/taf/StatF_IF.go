// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 1.6.0 by WSRD Tencent.
// Generated from `StatF.jce'
// **********************************************************************

package taf

import (
	context "context"
	"errors"
	"reflect"
	"tarsrpc/jce/taf"
	"tarsrpc/jce_parser/gojce"
	m "tarsrpc/tars/servant/model"
)

type StatF struct {
	s m.Servant
}

func (_obj *StatF) ReportMicMsg(msg map[StatMicMsgHead]StatMicMsgBody, bFromClient bool, _opt ...map[string]string) (_ret int32, _err error) {
	_oe := gojce.NewOutputStream()
	_oe.Write(reflect.ValueOf(&msg), 1)
	_oe.Write(reflect.ValueOf(&bFromClient), 2)
	var (
		_resp *taf.ResponsePacket
		err   error
	)
	_resp, err = _obj.s.Taf_invoke(context.TODO(), 0, "reportMicMsg", _oe.ToBytes())
	if err != nil {
		return _ret, err
	}
	_is := gojce.NewInputStream(_resp.SBuffer)
	r0, err := _is.Read(reflect.TypeOf(_ret), 0, true)
	if err != nil {
		return _ret, err
	}
	return r0.(int32), nil
}
func (_obj *StatF) ReportProxyMicMsg(msg map[StatMicMsgHead]StatMicMsgBody, proxyInfo ProxyInfo, _opt ...map[string]string) (_ret int32, _err error) {
	_oe := gojce.NewOutputStream()
	_oe.Write(reflect.ValueOf(&msg), 1)
	_oe.Write(reflect.ValueOf(&proxyInfo), 2)
	var (
		_resp *taf.ResponsePacket
		err   error
	)
	_resp, err = _obj.s.Taf_invoke(context.TODO(), 0, "reportProxyMicMsg", _oe.ToBytes())
	if err != nil {
		return _ret, err
	}
	_is := gojce.NewInputStream(_resp.SBuffer)
	r0, err := _is.Read(reflect.TypeOf(_ret), 0, true)
	if err != nil {
		return _ret, err
	}
	return r0.(int32), nil
}
func (_obj *StatF) ReportSampleMsg(msg []StatSampleMsg, _opt ...map[string]string) (_ret int32, _err error) {
	_oe := gojce.NewOutputStream()
	_oe.Write(reflect.ValueOf(&msg), 1)
	var (
		_resp *taf.ResponsePacket
		err   error
	)
	_resp, err = _obj.s.Taf_invoke(context.TODO(), 0, "reportSampleMsg", _oe.ToBytes())
	if err != nil {
		return _ret, err
	}
	_is := gojce.NewInputStream(_resp.SBuffer)
	r0, err := _is.Read(reflect.TypeOf(_ret), 0, true)
	if err != nil {
		return _ret, err
	}
	return r0.(int32), nil
}
func (_obj *StatF) SetServant(s m.Servant) {
	_obj.s = s
}

type _impStatF interface {
	ReportMicMsg(msg map[StatMicMsgHead]StatMicMsgBody, bFromClient bool) (int32, error)
	ReportProxyMicMsg(msg map[StatMicMsgHead]StatMicMsgBody, proxyInfo ProxyInfo) (int32, error)
	ReportSampleMsg(msg []StatSampleMsg) (int32, error)
}

func (_obj *StatF) Dispatch(ctx context.Context, _val interface{}, req *taf.RequestPacket) (*taf.ResponsePacket, error) {
	parms := gojce.NewInputStream(req.SBuffer)
	oe := gojce.NewOutputStream()
	_imp := _val.(_impStatF)
	switch req.SFuncName {
	case "reportMicMsg":
		var p_0 map[StatMicMsgHead]StatMicMsgBody
		t_0, err := parms.Read(reflect.TypeOf(p_0), 1, true)
		if err != nil {
			return nil, err
		}
		var p_1 bool
		t_1, err := parms.Read(reflect.TypeOf(p_1), 2, true)
		if err != nil {
			return nil, err
		}
		_ret, err := _imp.ReportMicMsg(t_0.(map[StatMicMsgHead]StatMicMsgBody), t_1.(bool))
		if err != nil {
			return nil, err
		}
		oe.Write(reflect.ValueOf(&_ret), 0)
	case "reportProxyMicMsg":
		var p_0 map[StatMicMsgHead]StatMicMsgBody
		t_0, err := parms.Read(reflect.TypeOf(p_0), 1, true)
		if err != nil {
			return nil, err
		}
		var p_1 ProxyInfo
		t_1, err := parms.Read(reflect.TypeOf(p_1), 2, true)
		if err != nil {
			return nil, err
		}
		_ret, err := _imp.ReportProxyMicMsg(t_0.(map[StatMicMsgHead]StatMicMsgBody), t_1.(ProxyInfo))
		if err != nil {
			return nil, err
		}
		oe.Write(reflect.ValueOf(&_ret), 0)
	case "reportSampleMsg":
		var p_0 []StatSampleMsg
		t_0, err := parms.Read(reflect.TypeOf(p_0), 1, true)
		if err != nil {
			return nil, err
		}
		_ret, err := _imp.ReportSampleMsg(t_0.([]StatSampleMsg))
		if err != nil {
			return nil, err
		}
		oe.Write(reflect.ValueOf(&_ret), 0)
	default:
		return nil, errors.New("func mismatch")
	}
	var status map[string]string
	return &taf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      oe.ToBytes(),
		Status:       status,
		SResultDesc:  "",
		Context:      req.Context,
	}, nil
}
