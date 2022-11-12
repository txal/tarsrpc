// @author kordenlu
// @创建时间 2017/09/14 12:12
// 功能描述:

package yyp

import (
	"fmt"
	"reflect"
	"runtime"
)

func Unmarshal32Str(data []byte, v interface{}) (err error) {
	var d DecodeState
	d.init(data)
	return d.unmarshal32Str(d.data, v)
}

func (d *DecodeState) unmarshal32Str(data []byte, v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic!!!")
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {

		return illegalInput
	}
	elem := rv.Elem()

	switch elem.Kind() {
	case reflect.Bool:
		if elem.CanSet() {
			elem.SetBool(d.readBool())
		}
	case reflect.Int8:
		if elem.CanSet() {
			elem.SetInt(int64(d.ReadI8()))
		}
	case reflect.Int16:
		if elem.CanSet() {
			elem.SetInt(int64(d.ReadI16()))
		}
	case reflect.Int32:
		if elem.CanSet() {
			elem.SetInt(int64(d.ReadI32()))
		}
	case reflect.Int64:
		if elem.CanSet() {
			elem.SetInt(int64(d.ReadI64()))
		}
	case reflect.Uint8:
		if elem.CanSet() {
			elem.SetUint(uint64(d.ReadU8()))
		}
	case reflect.Uint16:
		if elem.CanSet() {
			elem.SetUint(uint64(d.ReadU16()))
		}
	case reflect.Uint32:
		if elem.CanSet() {
			elem.SetUint(uint64(d.ReadU32()))
		}
	case reflect.Uint64:
		if elem.CanSet() {
			elem.SetUint(d.ReadU64())
		}
	case reflect.String:
		if elem.CanSet() {
			elem.SetString(d.readStupidString())
		}
	case reflect.Struct:
		if elem.CanSet() {
			d.Read32Struct(v)
		}
	case reflect.Slice:
		if elem.CanSet() {

			d.Read32Slice(v)
		}
	default:
		return NOtSupportType
	}
	return nil
}

func (e *DecodeState) readStupidString() string {
	len := e.ReadU32()
	//fmt.Printf("read32String len of str:%v", len)
	if len > 0 {
		start := e.off
		e.off += int(len)
		return string(e.data[start:e.off])
	} else {
		return ""
	}
}

func (e *DecodeState) Read32Struct(v interface{}) error {
	rv := reflect.ValueOf(v)
	elem := rv.Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if err := e.unmarshal32Str(e.data, field.Addr().Interface()); err != nil {
			return err
		}
	}
	return nil
}

func (e *DecodeState) Read32Slice(v interface{}) error {
	rv := reflect.ValueOf(v)
	elem := rv.Elem()

	len := e.ReadU32()
	slivalue := reflect.MakeSlice(elem.Type(), int(len), int(len))

	for i := 0; i < int(len); i++ {
		data := slivalue.Index(i)

		if err := e.unmarshal32Str(e.data, data.Addr().Interface()); err != nil {
			return err
		}
	}
	elem.Set(slivalue)
	return nil
}
