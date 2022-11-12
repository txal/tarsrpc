// @author kordenlu
// @创建时间 2017/08/16 17:42
// 功能描述:

package yyp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"reflect"
	"runtime"
)

type encodeState struct {
	bytes.Buffer // accumulated output
	scratch      [64]byte
}

func (e *encodeState) reflectValue(v reflect.Value) {
	err := valueEncoder(v)(e, v)
	if err != nil {
		log.Panicf("encode %v: %v", v, err)
	}
}

// newTypeEncoder constructs an encoderFunc for a type.
// The returned encoder only checks CanAddr when allowAddr is true.
func newTypeEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	case reflect.Bool:
		return boolEncoder
	case reflect.Int8:
		return WriteI8
	case reflect.Int16:
		return WriteI16
	case reflect.Int32:
		return WriteI32
	case reflect.Int64:
		return WriteI64
	case reflect.Uint8:
		return WriteU8
	case reflect.Uint16:
		return WriteU16
	case reflect.Uint32:
		return WriteU32
	case reflect.Uint64:
		return WriteU64
	case reflect.String:
		return writeString
	case reflect.Struct:
		return StructEncoder
	case reflect.Map:
		return MapEncoder
	case reflect.Slice:
		return SliceEncoder
	//case reflect.Array:
	//	return newArrayEncoder(t)
	default:
		return unsupportedTypeEncoder
	}
}

var (
	NOtSupportType = errors.New("not encode support type")
	InvalidValue   = errors.New("invalidvalue")
)

type encoderFunc func(e *encodeState, v reflect.Value) error

func invalidValueEncoder(e *encodeState, v reflect.Value) error {
	return InvalidValue
}

func unsupportedTypeEncoder(e *encodeState, v reflect.Value) error {
	return NOtSupportType
}

func boolEncoder(e *encodeState, v reflect.Value) error {
	if v.Bool() {
		return e.WriteByte(1)
	} else {
		return e.WriteByte(0)
	}
}

func writeString(e *encodeState, value reflect.Value) error {
	if err := WriteI16(e, reflect.ValueOf(len(value.String()))); err != nil {
		return err
	}
	_, err := e.WriteString(value.String())
	return err
}

func WriteU8(e *encodeState, value reflect.Value) error {
	loaclvalue := uint8(value.Uint())
	return e.WriteByte(loaclvalue)
}

func WriteU16(e *encodeState, value reflect.Value) error {
	loaclvalue := uint16(value.Uint())
	v := e.scratch[0:2]
	binary.LittleEndian.PutUint16(v, loaclvalue)
	_, err := e.Write(v)
	return err
}

func WriteU32(e *encodeState, value reflect.Value) error {
	loaclvalue := uint32(value.Uint())
	v := e.scratch[0:4]
	binary.LittleEndian.PutUint32(v, loaclvalue)
	_, err := e.Write(v)
	return err
}

func WriteU64(e *encodeState, value reflect.Value) error {
	loaclvalue := uint64(value.Uint())
	v := e.scratch[0:8]
	binary.LittleEndian.PutUint64(v, loaclvalue)
	_, err := e.Write(v)
	return err
}

func WriteI8(e *encodeState, value reflect.Value) error {
	return e.WriteByte(uint8(value.Int()))
}

func WriteI16(e *encodeState, value reflect.Value) error {
	v := e.scratch[0:2]
	binary.LittleEndian.PutUint16(v, uint16(value.Int()))
	_, err := e.Write(v)
	return err
}

func WriteI32(e *encodeState, value reflect.Value) error {
	v := e.scratch[0:4]
	binary.LittleEndian.PutUint32(v, uint32(value.Int()))
	_, err := e.Write(v)
	return err
}

func WriteI64(e *encodeState, value reflect.Value) error {
	v := e.scratch[0:8]
	binary.LittleEndian.PutUint64(v, uint64(value.Int()))
	_, err := e.Write(v)
	return err
}

func valueEncoder(v reflect.Value) encoderFunc {
	if !v.IsValid() {
		return invalidValueEncoder
	}
	return newTypeEncoder(v.Type())
}

func StructEncoder(e *encodeState, v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		e.reflectValue(v.Field(i))
	}
	return nil
}

// add by zc
func SliceEncoder(e *encodeState, v reflect.Value) error {
	//序列化长度uint32_t
	err := e.marshal(uint32(v.Len()))
	if err != nil {
		return err
	}
	//从头到尾依次序列数据
	for i := 0; i < v.Len(); i++ {
		e.reflectValue(v.Index(i))
	}
	return nil
}

func MapEncoder(e *encodeState, v reflect.Value) error {
	//序列化长度uint32_t
	err := e.marshal(uint32(v.Len()))
	if err != nil {
		return err
	}
	//从头到尾依次序列数据
	for _, key := range v.MapKeys() {
		e.reflectValue(key)
		e.reflectValue(v.MapIndex(key))
	}
	return nil
}

//end by zc
func (e *encodeState) marshal(v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			if s, ok := r.(string); ok {
				panic(s)
			}
			err = r.(error)
		}
	}()
	e.reflectValue(reflect.ValueOf(v))
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	e := &encodeState{}
	err := e.marshal(v)
	if err != nil {
		return nil, err
	}
	return e.Bytes(), nil
}
