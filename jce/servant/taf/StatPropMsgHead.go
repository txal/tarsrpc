// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 1.6.0 by WSRD Tencent.
// Generated from `PropertyF.jce'
// **********************************************************************

package taf

import (
	"reflect"
	"tarsrpc/jce_parser/gojce"
)

type StatPropMsgHead struct {
	ModuleName   string
	Ip           string
	PropertyName string
	SetName      string
	SetArea      string
	SetID        string
	SContainer   string
	IPropertyVer int32
}

func (_obj *StatPropMsgHead) resetDefault() {
	_obj.ModuleName = ""
	_obj.Ip = ""
	_obj.PropertyName = ""
	_obj.SetName = ""
	_obj.SetArea = ""
	_obj.SetID = ""
	_obj.SContainer = ""
	_obj.IPropertyVer = 1
}

func (_obj *StatPropMsgHead) WriteTo(_os gojce.JceOutputStream) error {
	var _err error
	if _err = _os.Write(reflect.ValueOf(&_obj.ModuleName), 0); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.Ip), 1); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.PropertyName), 2); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SetName), 3); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SetArea), 4); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SetID), 5); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SContainer), 6); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.IPropertyVer), 7); _err != nil {
		return _err
	}
	return nil
}

func (_obj *StatPropMsgHead) ReadFrom(_is gojce.JceInputStream) error {
	var _err error
	var _i interface{}
	_obj.resetDefault()
	_i, _err = _is.Read(reflect.TypeOf(_obj.ModuleName), 0, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.ModuleName = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.Ip), 1, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.Ip = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.PropertyName), 2, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.PropertyName = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SetName), 3, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SetName = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SetArea), 4, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SetArea = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SetID), 5, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SetID = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SContainer), 6, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SContainer = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.IPropertyVer), 7, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.IPropertyVer = _i.(int32)
	}
	return nil
}

func (_obj *StatPropMsgHead) Display(_ds gojce.JceDisplayer) {
	_ds.Display(reflect.ValueOf(&_obj.ModuleName), "moduleName")
	_ds.Display(reflect.ValueOf(&_obj.Ip), "ip")
	_ds.Display(reflect.ValueOf(&_obj.PropertyName), "propertyName")
	_ds.Display(reflect.ValueOf(&_obj.SetName), "setName")
	_ds.Display(reflect.ValueOf(&_obj.SetArea), "setArea")
	_ds.Display(reflect.ValueOf(&_obj.SetID), "setID")
	_ds.Display(reflect.ValueOf(&_obj.SContainer), "sContainer")
	_ds.Display(reflect.ValueOf(&_obj.IPropertyVer), "iPropertyVer")
}

func (_obj *StatPropMsgHead) WriteJson(_en gojce.JceJsonEncoder) ([]byte, error) {
	var _err error
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.ModuleName), "moduleName")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.Ip), "ip")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.PropertyName), "propertyName")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SetName), "setName")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SetArea), "setArea")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SetID), "setID")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SContainer), "sContainer")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.IPropertyVer), "iPropertyVer")
	if _err != nil {
		return nil, _err
	}
	return _en.ToBytes(), nil
}

func (_obj *StatPropMsgHead) ReadJson(_de gojce.JceJsonDecoder) error {
	return _de.DecodeJSON(reflect.ValueOf(_obj))
}