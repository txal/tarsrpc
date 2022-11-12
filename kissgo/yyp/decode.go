// @author kordenlu
// @创建时间 2017/08/17 16:12
// 功能描述:

package yyp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

// decodeState represents the state while decoding a JSON value.
type DecodeState struct {
	data []byte
	off  int // read offset in data
}

func (d *DecodeState) init(data []byte) *DecodeState {
	d.data = data
	d.off = 0
	return d
}

var (
	CannotSetError = errors.New("can't set")
	illegalInput   = errors.New("illegal input")
)

func (e *DecodeState) readBool() bool {
	if len(e.data) >= (e.off + 1) {
		localbool := uint8(e.data[e.off])
		e.off++
		if localbool != 1 {
			return false
		}
		return true
	} else {
		return false
	}
}

func (e *DecodeState) readString() string {
	len := e.ReadU16()
	//fmt.Printf("readString len of str:%v", len)
	if len > 0 {
		start := e.off
		e.off += int(len)
		return string(e.data[start:e.off])
	} else {
		return ""
	}
}

func (e *DecodeState) ReadU8() uint8 {
	if len(e.data) >= (e.off + 1) {
		buf := e.data[e.off]
		e.off++
		return uint8(buf)
	} else {
		return 0
	}
}

func (e *DecodeState) ReadU16() uint16 {
	if len(e.data) >= (e.off + 2) {
		buf := e.data[e.off : e.off+2]
		e.off += 2
		return binary.LittleEndian.Uint16(buf)
	}
	return 0
}

func (e *DecodeState) ReadU32() uint32 {
	if len(e.data) >= (e.off + 4) {
		buf := e.data[e.off : e.off+4]
		e.off += 4
		return binary.LittleEndian.Uint32(buf)
	}
	return 0
}

func (e *DecodeState) ReadU64() uint64 {
	if len(e.data) >= (e.off + 8) {
		buf := e.data[e.off : e.off+8]
		e.off += 8
		return binary.LittleEndian.Uint64(buf)
	}
	return 0
}

func (e *DecodeState) ReadI8() int8 {
	return int8(e.ReadU8())
}

func (e *DecodeState) ReadI16() int16 {
	//there is a bug :Recursion
	//return int16(e.ReadI16())
	return int16(e.ReadU16())
}

func (e *DecodeState) ReadI32() int32 {
	return int32(e.ReadU32())
}

func (e *DecodeState) ReadI64() int64 {
	return int64(e.ReadU64())
}

//add by zc
func (e *DecodeState) ReadStruct(v interface{}) error {
	rv := reflect.ValueOf(v)
	elem := rv.Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if err := e.unmarshal(e.data, field.Addr().Interface()); err != nil {
			return err
		}
	}
	return nil
}

func (e *DecodeState) ReadSlice(v interface{}) error {
	rv := reflect.ValueOf(v)
	elem := rv.Elem()

	len := e.ReadU32()
	slivalue := reflect.MakeSlice(elem.Type(), int(len), int(len))

	for i := 0; i < int(len); i++ {
		data := slivalue.Index(i)
		if err := e.unmarshal(e.data, data.Addr().Interface()); err != nil {
			return err
		}
	}
	elem.Set(slivalue)
	return nil
}

func (e *DecodeState) ReadMap(v interface{}) error {
	rv := reflect.ValueOf(v)
	elem := rv.Elem()

	len := e.ReadU32()
	mapvalue := reflect.MakeMapWithSize(elem.Type(), int(len))

	for i := 0; i < int(len); i++ {
		key := reflect.New(elem.Type().Key())
		if err := e.unmarshal(e.data, key.Interface()); err != nil {
			return err
		}
		val := reflect.New(elem.Type().Elem())
		if err := e.unmarshal(e.data, val.Interface()); err != nil {
			return err
		}
		mapvalue.SetMapIndex(reflect.Indirect(key), reflect.Indirect(val))
	}
	elem.Set(mapvalue)
	return nil
}

type decoderFunc func(e *encodeState, v reflect.Value) error

// modify by zc
func Unmarshal(data []byte, v interface{}) (err error) {
	var d DecodeState
	d.init(data)
	return d.unmarshal(d.data, v)
}

func (d *DecodeState) unmarshal(data []byte, v interface{}) (err error) {
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
			elem.SetString(d.readString())
		}
	case reflect.Struct:
		if elem.CanSet() {
			err = d.ReadStruct(v)
		}
	case reflect.Slice:
		if elem.CanSet() {
			err = d.ReadSlice(v)
		}
	case reflect.Map:
		if elem.CanSet() {
			err = d.ReadMap(v)
		}
	default:
		err = NOtSupportType
	}
	return err
}

// end modify by zc

/*
type decoderFunc func(e *encodeState, v reflect.Value) error

func Unmarshal(data []byte, v interface{}) (err error) {
	var d DecodeState
	d.init(data)

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		//zc
		fmt.Println("illinput")
		return illegalInput
	}

	elem := rv.Elem()

	for i :=0;i<elem.NumField();i++{

		field := elem.Field(i)
		switch field.Kind() {
		case reflect.Bool:
			if field.CanSet(){
				field.SetBool(d.readBool())
			}
		case reflect.Int8:
			if field.CanSet(){
				field.SetInt(int64(d.ReadI8()))
			}
		case reflect.Int16:

			if field.CanSet(){
				field.SetInt(int64(d.ReadI16()))
			}
		case reflect.Int32:
			if field.CanSet(){
				field.SetInt(int64(d.ReadI32()))
			}
		case reflect.Int64:
			if field.CanSet(){
				field.SetInt(int64(d.ReadI64()))
			}
		case reflect.Uint8:
			if field.CanSet(){
				field.SetUint(uint64(d.ReadU8()))
			}
		case reflect.Uint16:
			if field.CanSet(){
				field.SetUint(uint64(d.ReadU16()))
			}
		case reflect.Uint32:
			if field.CanSet(){
				field.SetUint(uint64(d.ReadU32()))
			}
		case reflect.Uint64:
			if field.CanSet(){
				field.SetUint(d.ReadU64())
			}
		case reflect.String:
			if field.CanSet(){
				field.SetString(d.readString())
			}
		case reflect.Struct:



		default:
			return NOtSupportType
		}
	}
	return nil
}
*/
