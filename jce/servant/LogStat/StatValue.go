// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 1.6.0 by WSRD Tencent.
// Generated from `PropertyPlusF.jce'
// **********************************************************************

package LogStat
import "reflect"
import "code.com/tars/goframework/jce_parser/gojce"

type StatValue struct {
    Policy string
    Value float32
    Count int64
}

func (_obj *StatValue) resetDefault() {
    _obj.Policy = ""
    _obj.Value = 0
    _obj.Count = 0
}

func (_obj *StatValue) WriteTo(_os gojce.JceOutputStream) error {
    var _err error
    if _err = _os.Write(reflect.ValueOf(&_obj.Policy), 0); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.Value), 1); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.Count), 2); _err != nil {
        return _err
    }
    return nil
}

func (_obj *StatValue) ReadFrom(_is gojce.JceInputStream) error {
    var _err error
    var _i interface{}
    _obj.resetDefault()
    _i, _err = _is.Read(reflect.TypeOf(_obj.Policy), 0, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Policy = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.Value), 1, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Value = _i.(float32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.Count), 2, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Count = _i.(int64)
    }
    return nil
}

func (_obj *StatValue) Display(_ds gojce.JceDisplayer) {
    _ds.Display(reflect.ValueOf(&_obj.Policy), "policy")
    _ds.Display(reflect.ValueOf(&_obj.Value), "value")
    _ds.Display(reflect.ValueOf(&_obj.Count), "count")
}

func (_obj *StatValue) WriteJson(_en gojce.JceJsonEncoder) ([]byte, error) {
    var _err error
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Policy), "policy")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Value), "value")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Count), "count")
    if _err != nil {
        return nil, _err
    }
    return _en.ToBytes(), nil
}

func (_obj *StatValue) ReadJson(_de gojce.JceJsonDecoder) error {
    return _de.DecodeJSON(reflect.ValueOf(_obj))
}

