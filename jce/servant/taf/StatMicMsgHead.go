// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 1.6.0 by WSRD Tencent.
// Generated from `StatF.jce'
// **********************************************************************

package taf

import (
	"reflect"
	"tarsrpc/jce_parser/gojce"
)

type StatMicMsgHead struct {
	MasterName       string
	SlaveName        string
	InterfaceName    string
	MasterIp         string
	SlaveIp          string
	SlavePort        int32
	ReturnValue      int32
	SlaveSetName     string
	SlaveSetArea     string
	SlaveSetID       string
	TafVersion       string
	SMasterSetInfo   string
	SSlaveContainer  string
	SMasterContainer string
	IStatVer         int32
}

func (_obj *StatMicMsgHead) resetDefault() {
	_obj.MasterName = ""
	_obj.SlaveName = ""
	_obj.InterfaceName = ""
	_obj.MasterIp = ""
	_obj.SlaveIp = ""
	_obj.SlavePort = 0
	_obj.ReturnValue = 0
	_obj.SlaveSetName = ""
	_obj.SlaveSetArea = ""
	_obj.SlaveSetID = ""
	_obj.TafVersion = ""
	_obj.SMasterSetInfo = ""
	_obj.SSlaveContainer = ""
	_obj.SMasterContainer = ""
	_obj.IStatVer = 1
}

func (_obj *StatMicMsgHead) WriteTo(_os gojce.JceOutputStream) error {
	var _err error
	if _err = _os.Write(reflect.ValueOf(&_obj.MasterName), 0); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SlaveName), 1); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.InterfaceName), 2); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.MasterIp), 3); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SlaveIp), 4); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SlavePort), 5); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.ReturnValue), 6); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SlaveSetName), 7); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SlaveSetArea), 8); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SlaveSetID), 9); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.TafVersion), 10); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SMasterSetInfo), 11); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SSlaveContainer), 12); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.SMasterContainer), 13); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.IStatVer), 14); _err != nil {
		return _err
	}
	return nil
}

func (_obj *StatMicMsgHead) ReadFrom(_is gojce.JceInputStream) error {
	var _err error
	var _i interface{}
	_obj.resetDefault()
	_i, _err = _is.Read(reflect.TypeOf(_obj.MasterName), 0, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.MasterName = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SlaveName), 1, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SlaveName = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.InterfaceName), 2, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.InterfaceName = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.MasterIp), 3, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.MasterIp = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SlaveIp), 4, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SlaveIp = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SlavePort), 5, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SlavePort = _i.(int32)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.ReturnValue), 6, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.ReturnValue = _i.(int32)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SlaveSetName), 7, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SlaveSetName = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SlaveSetArea), 8, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SlaveSetArea = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SlaveSetID), 9, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SlaveSetID = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.TafVersion), 10, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.TafVersion = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SMasterSetInfo), 11, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SMasterSetInfo = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SSlaveContainer), 12, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SSlaveContainer = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.SMasterContainer), 13, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SMasterContainer = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.IStatVer), 14, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.IStatVer = _i.(int32)
	}
	return nil
}

func (_obj *StatMicMsgHead) Display(_ds gojce.JceDisplayer) {
	_ds.Display(reflect.ValueOf(&_obj.MasterName), "masterName")
	_ds.Display(reflect.ValueOf(&_obj.SlaveName), "slaveName")
	_ds.Display(reflect.ValueOf(&_obj.InterfaceName), "interfaceName")
	_ds.Display(reflect.ValueOf(&_obj.MasterIp), "masterIp")
	_ds.Display(reflect.ValueOf(&_obj.SlaveIp), "slaveIp")
	_ds.Display(reflect.ValueOf(&_obj.SlavePort), "slavePort")
	_ds.Display(reflect.ValueOf(&_obj.ReturnValue), "returnValue")
	_ds.Display(reflect.ValueOf(&_obj.SlaveSetName), "slaveSetName")
	_ds.Display(reflect.ValueOf(&_obj.SlaveSetArea), "slaveSetArea")
	_ds.Display(reflect.ValueOf(&_obj.SlaveSetID), "slaveSetID")
	_ds.Display(reflect.ValueOf(&_obj.TafVersion), "tafVersion")
	_ds.Display(reflect.ValueOf(&_obj.SMasterSetInfo), "sMasterSetInfo")
	_ds.Display(reflect.ValueOf(&_obj.SSlaveContainer), "sSlaveContainer")
	_ds.Display(reflect.ValueOf(&_obj.SMasterContainer), "sMasterContainer")
	_ds.Display(reflect.ValueOf(&_obj.IStatVer), "iStatVer")
}

func (_obj *StatMicMsgHead) WriteJson(_en gojce.JceJsonEncoder) ([]byte, error) {
	var _err error
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.MasterName), "masterName")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SlaveName), "slaveName")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.InterfaceName), "interfaceName")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.MasterIp), "masterIp")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SlaveIp), "slaveIp")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SlavePort), "slavePort")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.ReturnValue), "returnValue")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SlaveSetName), "slaveSetName")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SlaveSetArea), "slaveSetArea")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SlaveSetID), "slaveSetID")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.TafVersion), "tafVersion")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SMasterSetInfo), "sMasterSetInfo")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SSlaveContainer), "sSlaveContainer")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SMasterContainer), "sMasterContainer")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.IStatVer), "iStatVer")
	if _err != nil {
		return nil, _err
	}
	return _en.ToBytes(), nil
}

func (_obj *StatMicMsgHead) ReadJson(_de gojce.JceJsonDecoder) error {
	return _de.DecodeJSON(reflect.ValueOf(_obj))
}