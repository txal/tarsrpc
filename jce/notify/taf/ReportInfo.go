// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 1.6.0 by WSRD Tencent.
// Generated from `Notify.jce'
// **********************************************************************

package taf
import "reflect"
import "code.com/tars/goframework/jce_parser/gojce"

type ReportInfo struct {
    EType int
    SApp string
    SSet string
    SContainer string
    SServer string
    SMessage string
    SThreadId string
    ELevel int
}

func (_obj *ReportInfo) resetDefault() {
    _obj.SApp = ""
    _obj.SSet = ""
    _obj.SContainer = ""
    _obj.SServer = ""
    _obj.SMessage = ""
    _obj.SThreadId = ""
}

func (_obj *ReportInfo) WriteTo(_os gojce.JceOutputStream) error {
    var _err error
    if _err = _os.Write(reflect.ValueOf(&_obj.EType), 1); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.SApp), 2); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.SSet), 3); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.SContainer), 4); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.SServer), 5); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.SMessage), 6); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.SThreadId), 7); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.ELevel), 8); _err != nil {
        return _err
    }
    return nil
}

func (_obj *ReportInfo) ReadFrom(_is gojce.JceInputStream) error {
    var _err error
    var _i interface{}
    _obj.resetDefault()
    _i, _err = _is.Read(reflect.TypeOf(_obj.EType), 1, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.EType = _i.(int)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.SApp), 2, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.SApp = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.SSet), 3, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.SSet = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.SContainer), 4, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.SContainer = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.SServer), 5, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.SServer = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.SMessage), 6, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.SMessage = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.SThreadId), 7, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.SThreadId = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.ELevel), 8, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.ELevel = _i.(int)
    }
    return nil
}

func (_obj *ReportInfo) Display(_ds gojce.JceDisplayer) {
    _ds.Display(reflect.ValueOf(&_obj.EType), "eType")
    _ds.Display(reflect.ValueOf(&_obj.SApp), "sApp")
    _ds.Display(reflect.ValueOf(&_obj.SSet), "sSet")
    _ds.Display(reflect.ValueOf(&_obj.SContainer), "sContainer")
    _ds.Display(reflect.ValueOf(&_obj.SServer), "sServer")
    _ds.Display(reflect.ValueOf(&_obj.SMessage), "sMessage")
    _ds.Display(reflect.ValueOf(&_obj.SThreadId), "sThreadId")
    _ds.Display(reflect.ValueOf(&_obj.ELevel), "eLevel")
}

func (_obj *ReportInfo) WriteJson(_en gojce.JceJsonEncoder) ([]byte, error) {
    var _err error
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.EType), "eType")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.SApp), "sApp")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.SSet), "sSet")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.SContainer), "sContainer")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.SServer), "sServer")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.SMessage), "sMessage")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.SThreadId), "sThreadId")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.ELevel), "eLevel")
    if _err != nil {
        return nil, _err
    }
    return _en.ToBytes(), nil
}

func (_obj *ReportInfo) ReadJson(_de gojce.JceJsonDecoder) error {
    return _de.DecodeJSON(reflect.ValueOf(_obj))
}

