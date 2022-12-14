// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 1.6.0 by WSRD Tencent.
// Generated from `StatF.jce'
// **********************************************************************

package taf
import "reflect"
import "code.com/tars/goframework/jce_parser/gojce"

type StatMicMsgBody struct {
    Count int32
    TimeoutCount int32
    ExecCount int32
    IntervalCount map[int32]int32
    TotalRspTime int64
    MaxRspTime int32
    MinRspTime int32
    RetValue map[int64]int32
    WeightValue int32
    WeightCount int32
}

func (_obj *StatMicMsgBody) resetDefault() {
    _obj.Count = 0
    _obj.TimeoutCount = 0
    _obj.ExecCount = 0
    _obj.TotalRspTime = 0
    _obj.MaxRspTime = 0
    _obj.MinRspTime = 0
    _obj.WeightValue = 0
    _obj.WeightCount = 0
}

func (_obj *StatMicMsgBody) WriteTo(_os gojce.JceOutputStream) error {
    var _err error
    if _err = _os.Write(reflect.ValueOf(&_obj.Count), 0); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.TimeoutCount), 1); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.ExecCount), 2); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.IntervalCount), 3); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.TotalRspTime), 4); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.MaxRspTime), 5); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.MinRspTime), 6); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.RetValue), 7); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.WeightValue), 8); _err != nil {
        return _err
    }
    if _err = _os.Write(reflect.ValueOf(&_obj.WeightCount), 9); _err != nil {
        return _err
    }
    return nil
}

func (_obj *StatMicMsgBody) ReadFrom(_is gojce.JceInputStream) error {
    var _err error
    var _i interface{}
    _obj.resetDefault()
    _i, _err = _is.Read(reflect.TypeOf(_obj.Count), 0, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.Count = _i.(int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.TimeoutCount), 1, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.TimeoutCount = _i.(int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.ExecCount), 2, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.ExecCount = _i.(int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.IntervalCount), 3, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.IntervalCount = _i.(map[int32]int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.TotalRspTime), 4, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.TotalRspTime = _i.(int64)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.MaxRspTime), 5, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.MaxRspTime = _i.(int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.MinRspTime), 6, true)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.MinRspTime = _i.(int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.RetValue), 7, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.RetValue = _i.(map[int64]int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.WeightValue), 8, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.WeightValue = _i.(int32)
    }
    _i, _err = _is.Read(reflect.TypeOf(_obj.WeightCount), 9, false)
    if _err != nil {
        return _err
    }
    if _i != nil {
        _obj.WeightCount = _i.(int32)
    }
    return nil
}

func (_obj *StatMicMsgBody) Display(_ds gojce.JceDisplayer) {
    _ds.Display(reflect.ValueOf(&_obj.Count), "count")
    _ds.Display(reflect.ValueOf(&_obj.TimeoutCount), "timeoutCount")
    _ds.Display(reflect.ValueOf(&_obj.ExecCount), "execCount")
    _ds.Display(reflect.ValueOf(&_obj.IntervalCount), "intervalCount")
    _ds.Display(reflect.ValueOf(&_obj.TotalRspTime), "totalRspTime")
    _ds.Display(reflect.ValueOf(&_obj.MaxRspTime), "maxRspTime")
    _ds.Display(reflect.ValueOf(&_obj.MinRspTime), "minRspTime")
    _ds.Display(reflect.ValueOf(&_obj.RetValue), "retValue")
    _ds.Display(reflect.ValueOf(&_obj.WeightValue), "weightValue")
    _ds.Display(reflect.ValueOf(&_obj.WeightCount), "weightCount")
}

func (_obj *StatMicMsgBody) WriteJson(_en gojce.JceJsonEncoder) ([]byte, error) {
    var _err error
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.Count), "count")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.TimeoutCount), "timeoutCount")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.ExecCount), "execCount")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.IntervalCount), "intervalCount")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.TotalRspTime), "totalRspTime")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.MaxRspTime), "maxRspTime")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.MinRspTime), "minRspTime")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.RetValue), "retValue")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.WeightValue), "weightValue")
    if _err != nil {
        return nil, _err
    }
    _err = _en.EncodeJSON(reflect.ValueOf(&_obj.WeightCount), "weightCount")
    if _err != nil {
        return nil, _err
    }
    return _en.ToBytes(), nil
}

func (_obj *StatMicMsgBody) ReadJson(_de gojce.JceJsonDecoder) error {
    return _de.DecodeJSON(reflect.ValueOf(_obj))
}

