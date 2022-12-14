// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 1.6.0 by WSRD Tencent.
// Generated from `Config.jce'
// **********************************************************************

package taf
import "reflect"
import "code.com/tars/goframework/jce_parser/gojce"

type ConfigInfo struct {
    Appname string
    Servername string
    Filename string
    BAppOnly bool
    Host string
    Setdivision string
    Containername string
}

func (_obj *ConfigInfo) resetDefault() {
    _obj.Appname = ""
    _obj.Servername = ""
    _obj.Filename = ""
    _obj.BAppOnly = false
    _obj.Host = ""
    _obj.Setdivision = ""
    _obj.Containername = ""
}

func (_obj *ConfigInfo) WriteTo(_os gojce.JceOutputStream) error {
    var _err error
    if _err = _os.Write(reflect.ValueOf(&_obj.Appname), 0); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.Servername), 1); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.Filename), 2); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.BAppOnly), 3); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.Host), 4); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.Setdivision), 5); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.Containername), 6); _err != nil {
        return _err
    }
    return nil
}

func (_obj *ConfigInfo) ReadFrom(_is gojce.JceInputStream) error {
    var _err error
    var _i interface{}
    _obj.resetDefault()
    _i, _err = _is.Read(reflect.TypeOf(_obj.Appname), 0, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Appname = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.Servername), 1, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Servername = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.Filename), 2, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Filename = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.BAppOnly), 3, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.BAppOnly = _i.(bool)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.Host), 4, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Host = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.Setdivision), 5, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Setdivision = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.Containername), 6, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Containername = _i.(string)
    }
    return nil
}

func (_obj *ConfigInfo) Display(_ds gojce.JceDisplayer) {
    _ds.Display(reflect.ValueOf(&_obj.Appname), "appname")
    _ds.Display(reflect.ValueOf(&_obj.Servername), "servername")
    _ds.Display(reflect.ValueOf(&_obj.Filename), "filename")
    _ds.Display(reflect.ValueOf(&_obj.BAppOnly), "bAppOnly")
    _ds.Display(reflect.ValueOf(&_obj.Host), "host")
    _ds.Display(reflect.ValueOf(&_obj.Setdivision), "setdivision")
    _ds.Display(reflect.ValueOf(&_obj.Containername), "containername")
}

func (_obj *ConfigInfo) WriteJson(_en gojce.JceJsonEncoder) ([]byte, error) {
    var _err error
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Appname), "appname")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Servername), "servername")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Filename), "filename")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.BAppOnly), "bAppOnly")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Host), "host")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Setdivision), "setdivision")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Containername), "containername")
    if _err != nil {
        return nil, _err
    }
    return _en.ToBytes(), nil
}

func (_obj *ConfigInfo) ReadJson(_de gojce.JceJsonDecoder) error {
    return _de.DecodeJSON(reflect.ValueOf(_obj))
}

