// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 1.6.0 by WSRD Tencent.
// Generated from `QueryF.jce'
// **********************************************************************

package taf
import (
    "code.com/tars/goframework/jce/taf"
    "code.com/tars/goframework/jce_parser/gojce"
    "reflect"
    m "code.com/tars/goframework/tars/servant/model"
    "errors"
    context "context"
    )
type QueryF struct {
    s m.Servant
}
func (_obj *QueryF) FindObjectById(id string,_opt ...map[string]string )(_ret []EndpointF,_err error){
    _oe := gojce.NewOutputStream()
    _oe.Write(reflect.ValueOf(&id), 1)
    var (
        _resp *taf.ResponsePacket
        err error
    )
    _resp,err = _obj.s.Taf_invoke(context.TODO(),0,"findObjectById", _oe.ToBytes())
    if err != nil {
        return _ret,err
    }
    _is := gojce.NewInputStream(_resp.SBuffer)
    r0, err := _is.Read(reflect.TypeOf(_ret), 0 ,true)
    if err!=nil {
        return _ret,err
    }
    return r0.([]EndpointF),nil
}
func (_obj *QueryF) FindObjectById4All(id string,activeEp *[]EndpointF,inactiveEp *[]EndpointF,_opt ...map[string]string )(_ret int32,_err error){
    _oe := gojce.NewOutputStream()
    _oe.Write(reflect.ValueOf(&id), 1)
    var (
        _resp *taf.ResponsePacket
        err error
    )
    _resp,err = _obj.s.Taf_invoke(context.TODO(),0,"findObjectById4All", _oe.ToBytes())
    if err != nil {
        return _ret,err
    }
    _is := gojce.NewInputStream(_resp.SBuffer)
    r0, err := _is.Read(reflect.TypeOf(_ret), 0 ,true)
    if err!=nil {
        return _ret,err
    }
    r_2, err := _is.Read(reflect.TypeOf(*activeEp),2,true)
    if err!=nil {
        return _ret,err
    }
    tmp_2:=r_2.([]EndpointF)
    *activeEp=tmp_2
    r_3, err := _is.Read(reflect.TypeOf(*inactiveEp),3,true)
    if err!=nil {
        return _ret,err
    }
    tmp_3:=r_3.([]EndpointF)
    *inactiveEp=tmp_3
    return r0.(int32),nil
}
func (_obj *QueryF) FindObjectById4Any(id string,activeEp *[]EndpointF,inactiveEp *[]EndpointF,_opt ...map[string]string )(_ret int32,_err error){
    _oe := gojce.NewOutputStream()
    _oe.Write(reflect.ValueOf(&id), 1)
    var (
        _resp *taf.ResponsePacket
        err error
    )
    _resp,err = _obj.s.Taf_invoke(context.TODO(),0,"findObjectById4Any", _oe.ToBytes())
    if err != nil {
        return _ret,err
    }
    _is := gojce.NewInputStream(_resp.SBuffer)
    r0, err := _is.Read(reflect.TypeOf(_ret), 0 ,true)
    if err!=nil {
        return _ret,err
    }
    r_2, err := _is.Read(reflect.TypeOf(*activeEp),2,true)
    if err!=nil {
        return _ret,err
    }
    tmp_2:=r_2.([]EndpointF)
    *activeEp=tmp_2
    r_3, err := _is.Read(reflect.TypeOf(*inactiveEp),3,true)
    if err!=nil {
        return _ret,err
    }
    tmp_3:=r_3.([]EndpointF)
    *inactiveEp=tmp_3
    return r0.(int32),nil
}
func (_obj *QueryF) FindObjectByIdInSameGroup(id string,activeEp *[]EndpointF,inactiveEp *[]EndpointF,_opt ...map[string]string )(_ret int32,_err error){
    _oe := gojce.NewOutputStream()
    _oe.Write(reflect.ValueOf(&id), 1)
    var (
        _resp *taf.ResponsePacket
        err error
    )
    _resp,err = _obj.s.Taf_invoke(context.TODO(),0,"findObjectByIdInSameGroup", _oe.ToBytes())
    if err != nil {
        return _ret,err
    }
    _is := gojce.NewInputStream(_resp.SBuffer)
    r0, err := _is.Read(reflect.TypeOf(_ret), 0 ,true)
    if err!=nil {
        return _ret,err
    }
    r_2, err := _is.Read(reflect.TypeOf(*activeEp),2,true)
    if err!=nil {
        return _ret,err
    }
    tmp_2:=r_2.([]EndpointF)
    *activeEp=tmp_2
    r_3, err := _is.Read(reflect.TypeOf(*inactiveEp),3,true)
    if err!=nil {
        return _ret,err
    }
    tmp_3:=r_3.([]EndpointF)
    *inactiveEp=tmp_3
    return r0.(int32),nil
}
func (_obj *QueryF) FindObjectByIdInSameSet(id string,setId string,activeEp *[]EndpointF,inactiveEp *[]EndpointF,_opt ...map[string]string )(_ret int32,_err error){
    _oe := gojce.NewOutputStream()
    _oe.Write(reflect.ValueOf(&id), 1)
    _oe.Write(reflect.ValueOf(&setId), 2)
    var (
        _resp *taf.ResponsePacket
        err error
    )
    _resp,err = _obj.s.Taf_invoke(context.TODO(),0,"findObjectByIdInSameSet", _oe.ToBytes())
    if err != nil {
        return _ret,err
    }
    _is := gojce.NewInputStream(_resp.SBuffer)
    r0, err := _is.Read(reflect.TypeOf(_ret), 0 ,true)
    if err!=nil {
        return _ret,err
    }
    r_3, err := _is.Read(reflect.TypeOf(*activeEp),3,true)
    if err!=nil {
        return _ret,err
    }
    tmp_3:=r_3.([]EndpointF)
    *activeEp=tmp_3
    r_4, err := _is.Read(reflect.TypeOf(*inactiveEp),4,true)
    if err!=nil {
        return _ret,err
    }
    tmp_4:=r_4.([]EndpointF)
    *inactiveEp=tmp_4
    return r0.(int32),nil
}
func (_obj *QueryF) FindObjectByIdInSameStation(id string,sStation string,activeEp *[]EndpointF,inactiveEp *[]EndpointF,_opt ...map[string]string )(_ret int32,_err error){
    _oe := gojce.NewOutputStream()
    _oe.Write(reflect.ValueOf(&id), 1)
    _oe.Write(reflect.ValueOf(&sStation), 2)
    var (
        _resp *taf.ResponsePacket
        err error
    )
    _resp,err = _obj.s.Taf_invoke(context.TODO(),0,"findObjectByIdInSameStation", _oe.ToBytes())
    if err != nil {
        return _ret,err
    }
    _is := gojce.NewInputStream(_resp.SBuffer)
    r0, err := _is.Read(reflect.TypeOf(_ret), 0 ,true)
    if err!=nil {
        return _ret,err
    }
    r_3, err := _is.Read(reflect.TypeOf(*activeEp),3,true)
    if err!=nil {
        return _ret,err
    }
    tmp_3:=r_3.([]EndpointF)
    *activeEp=tmp_3
    r_4, err := _is.Read(reflect.TypeOf(*inactiveEp),4,true)
    if err!=nil {
        return _ret,err
    }
    tmp_4:=r_4.([]EndpointF)
    *inactiveEp=tmp_4
    return r0.(int32),nil
}
func(_obj *QueryF) SetServant(s m.Servant){
    _obj.s = s
}
type _impQueryF interface {
    FindObjectById(id string) ([]EndpointF,error)
    FindObjectById4All(id string,activeEp *[]EndpointF,inactiveEp *[]EndpointF) (int32,error)
    FindObjectById4Any(id string,activeEp *[]EndpointF,inactiveEp *[]EndpointF) (int32,error)
    FindObjectByIdInSameGroup(id string,activeEp *[]EndpointF,inactiveEp *[]EndpointF) (int32,error)
    FindObjectByIdInSameSet(id string,setId string,activeEp *[]EndpointF,inactiveEp *[]EndpointF) (int32,error)
    FindObjectByIdInSameStation(id string,sStation string,activeEp *[]EndpointF,inactiveEp *[]EndpointF) (int32,error)
}
func(_obj *QueryF) Dispatch(ctx context.Context,_val interface{}, req *taf.RequestPacket) (*taf.ResponsePacket,error){
    parms := gojce.NewInputStream(req.SBuffer)
    oe := gojce.NewOutputStream()
    _imp := _val.(_impQueryF)
    switch req.SFuncName {
        case "findObjectById":
            var p_0 string
            t_0,err := parms.Read(reflect.TypeOf(p_0),1,true)
            if err != nil{
                return nil,err
            }
            _ret,err := _imp.FindObjectById(t_0.(string))
            if err != nil{
                return nil,err
            }
            oe.Write(reflect.ValueOf(&_ret), 0)
        case "findObjectById4All":
            var p_0 string
            t_0,err := parms.Read(reflect.TypeOf(p_0),1,true)
            if err != nil{
                return nil,err
            }
            var o_1 []EndpointF
            var o_2 []EndpointF
            _ret,err := _imp.FindObjectById4All(t_0.(string),&o_1,&o_2)
            if err != nil{
                return nil,err
            }
            oe.Write(reflect.ValueOf(&_ret), 0)
            oe.Write(reflect.ValueOf(&o_1),2)
            oe.Write(reflect.ValueOf(&o_2),3)
        case "findObjectById4Any":
            var p_0 string
            t_0,err := parms.Read(reflect.TypeOf(p_0),1,true)
            if err != nil{
                return nil,err
            }
            var o_1 []EndpointF
            var o_2 []EndpointF
            _ret,err := _imp.FindObjectById4Any(t_0.(string),&o_1,&o_2)
            if err != nil{
                return nil,err
            }
            oe.Write(reflect.ValueOf(&_ret), 0)
            oe.Write(reflect.ValueOf(&o_1),2)
            oe.Write(reflect.ValueOf(&o_2),3)
        case "findObjectByIdInSameGroup":
            var p_0 string
            t_0,err := parms.Read(reflect.TypeOf(p_0),1,true)
            if err != nil{
                return nil,err
            }
            var o_1 []EndpointF
            var o_2 []EndpointF
            _ret,err := _imp.FindObjectByIdInSameGroup(t_0.(string),&o_1,&o_2)
            if err != nil{
                return nil,err
            }
            oe.Write(reflect.ValueOf(&_ret), 0)
            oe.Write(reflect.ValueOf(&o_1),2)
            oe.Write(reflect.ValueOf(&o_2),3)
        case "findObjectByIdInSameSet":
            var p_0 string
            t_0,err := parms.Read(reflect.TypeOf(p_0),1,true)
            if err != nil{
                return nil,err
            }
            var p_1 string
            t_1,err := parms.Read(reflect.TypeOf(p_1),2,true)
            if err != nil{
                return nil,err
            }
            var o_2 []EndpointF
            var o_3 []EndpointF
            _ret,err := _imp.FindObjectByIdInSameSet(t_0.(string),t_1.(string),&o_2,&o_3)
            if err != nil{
                return nil,err
            }
            oe.Write(reflect.ValueOf(&_ret), 0)
            oe.Write(reflect.ValueOf(&o_2),3)
            oe.Write(reflect.ValueOf(&o_3),4)
        case "findObjectByIdInSameStation":
            var p_0 string
            t_0,err := parms.Read(reflect.TypeOf(p_0),1,true)
            if err != nil{
                return nil,err
            }
            var p_1 string
            t_1,err := parms.Read(reflect.TypeOf(p_1),2,true)
            if err != nil{
                return nil,err
            }
            var o_2 []EndpointF
            var o_3 []EndpointF
            _ret,err := _imp.FindObjectByIdInSameStation(t_0.(string),t_1.(string),&o_2,&o_3)
            if err != nil{
                return nil,err
            }
            oe.Write(reflect.ValueOf(&_ret), 0)
            oe.Write(reflect.ValueOf(&o_2),3)
            oe.Write(reflect.ValueOf(&o_3),4)
        default:
            return nil,errors.New("func mismatch")
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
    },nil
}
