// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 3.0.0.29 by WSRD Tencent.
// Generated from `./module/new.jce'
// **********************************************************************

package vvideo
import "reflect"
import "yytars/jce_parser/gojce"

type FileListPlayInfo struct {
    M_vid string
    M_title string
    M_type int32
    M_status int32
    M_duration int32
    M_uin int64
    M_vuid int64
    M_deliverMethod int32
    M_testMap map[string]int32
    M_list []FileListInfo
    M_itestint int32
    M_itestuint int64
    M_itestlong int64
    M_itestshort int16
    M_itestbyte byte
    M_itestfloat float32
    M_itestdouble float64
    M_iteststring string
}

func (_obj *FileListPlayInfo) resetDefault() {
    _obj.M_vid = ""
    _obj.M_title = ""
    _obj.M_type = 0
    _obj.M_status = 0
    _obj.M_duration = 0
    _obj.M_uin = 0
    _obj.M_vuid = 0
    _obj.M_deliverMethod = 0
    _obj.M_itestint = 21
    _obj.M_itestuint = 22
    _obj.M_itestlong = 23
    _obj.M_itestshort = 24
    _obj.M_itestbyte = 25
    _obj.M_itestfloat = 26.1
    _obj.M_itestdouble = 27.2
    _obj.M_iteststring = "28"
}

func (_obj *FileListPlayInfo) WriteTo(_os gojce.JceOutputStream) error {
    var _err error
    if _err = _os.Write(reflect.ValueOf(&_obj.M_vid), 1); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_title), 2); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_type), 3); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_status), 4); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_duration), 5); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_uin), 6); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_vuid), 7); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_deliverMethod), 8); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_testMap), 19); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_list), 20); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_itestint), 21); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_itestuint), 22); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_itestlong), 23); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_itestshort), 24); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_itestbyte), 25); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_itestfloat), 26); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_itestdouble), 27); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.M_iteststring), 28); _err != nil {
        return _err
    }
    return nil
}

func (_obj *FileListPlayInfo) ReadFrom(_is gojce.JceInputStream) error {
    var _err error
    var _i interface{}
    _obj.resetDefault()
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_vid), 1, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_vid = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_title), 2, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_title = _i.(string)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_type), 3, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_type = _i.(int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_status), 4, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_status = _i.(int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_duration), 5, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_duration = _i.(int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_uin), 6, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_uin = _i.(int64)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_vuid), 7, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_vuid = _i.(int64)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_deliverMethod), 8, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_deliverMethod = _i.(int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_testMap), 19, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_testMap = _i.(map[string]int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_list), 20, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_list = _i.([]FileListInfo)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_itestint), 21, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_itestint = _i.(int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_itestuint), 22, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_itestuint = _i.(int64)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_itestlong), 23, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_itestlong = _i.(int64)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_itestshort), 24, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_itestshort = _i.(int16)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_itestbyte), 25, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_itestbyte = _i.(byte)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_itestfloat), 26, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_itestfloat = _i.(float32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_itestdouble), 27, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_itestdouble = _i.(float64)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.M_iteststring), 28, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.M_iteststring = _i.(string)
    }
    return nil
}

func (_obj *FileListPlayInfo) Display(_ds gojce.JceDisplayer) {
    _ds.Display(reflect.ValueOf(&_obj.M_vid), "vid")
    _ds.Display(reflect.ValueOf(&_obj.M_title), "title")
    _ds.Display(reflect.ValueOf(&_obj.M_type), "type")
    _ds.Display(reflect.ValueOf(&_obj.M_status), "status")
    _ds.Display(reflect.ValueOf(&_obj.M_duration), "duration")
    _ds.Display(reflect.ValueOf(&_obj.M_uin), "uin")
    _ds.Display(reflect.ValueOf(&_obj.M_vuid), "vuid")
    _ds.Display(reflect.ValueOf(&_obj.M_deliverMethod), "deliverMethod")
    _ds.Display(reflect.ValueOf(&_obj.M_testMap), "testMap")
    _ds.Display(reflect.ValueOf(&_obj.M_list), "list")
    _ds.Display(reflect.ValueOf(&_obj.M_itestint), "itestint")
    _ds.Display(reflect.ValueOf(&_obj.M_itestuint), "itestuint")
    _ds.Display(reflect.ValueOf(&_obj.M_itestlong), "itestlong")
    _ds.Display(reflect.ValueOf(&_obj.M_itestshort), "itestshort")
    _ds.Display(reflect.ValueOf(&_obj.M_itestbyte), "itestbyte")
    _ds.Display(reflect.ValueOf(&_obj.M_itestfloat), "itestfloat")
    _ds.Display(reflect.ValueOf(&_obj.M_itestdouble), "itestdouble")
    _ds.Display(reflect.ValueOf(&_obj.M_iteststring), "iteststring")
}

