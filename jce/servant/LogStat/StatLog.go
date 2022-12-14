// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 1.6.0 by WSRD Tencent.
// Generated from `PropertyPlusF.jce'
// **********************************************************************

package LogStat
import "reflect"
import "code.com/tars/goframework/jce_parser/gojce"

type StatLog struct {
    Logname string
    Content []string
    Date string
    Flag string
    Extend int32
}

func (_obj *StatLog) resetDefault() {
    _obj.Logname = ""
    _obj.Date = ""
    _obj.Flag = ""
    _obj.Extend = 0
}

func (_obj *StatLog) WriteTo(_os gojce.JceOutputStream) error {
    var _err error
    if _err = _os.Write(reflect.ValueOf(&_obj.Logname), 0); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.Content), 1); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.Date), 2); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.Flag), 3); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.Extend), 4); _err != nil {
        return _err
    }
    return nil
}

func (_obj *StatLog) ReadFrom(_is gojce.JceInputStream) error {
    var _err error
    var _i interface{}
    _obj.resetDefault()
    _i, _err = _is.Read(reflect.TypeOf(_obj.Logname), 0, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Logname = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.Content), 1, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Content = _i.([]string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.Date), 2, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Date = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.Flag), 3, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Flag = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.Extend), 4, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Extend = _i.(int32)
    }
    return nil
}

func (_obj *StatLog) Display(_ds gojce.JceDisplayer) {
    _ds.Display(reflect.ValueOf(&_obj.Logname), "logname")
    _ds.Display(reflect.ValueOf(&_obj.Content), "content")
    _ds.Display(reflect.ValueOf(&_obj.Date), "date")
    _ds.Display(reflect.ValueOf(&_obj.Flag), "flag")
    _ds.Display(reflect.ValueOf(&_obj.Extend), "extend")
}

func (_obj *StatLog) WriteJson(_en gojce.JceJsonEncoder) ([]byte, error) {
    var _err error
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Logname), "logname")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Content), "content")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Date), "date")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Flag), "flag")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Extend), "extend")
    if _err != nil {
        return nil, _err
    }
    return _en.ToBytes(), nil
}

func (_obj *StatLog) ReadJson(_de gojce.JceJsonDecoder) error {
    return _de.DecodeJSON(reflect.ValueOf(_obj))
}

