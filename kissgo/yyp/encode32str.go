// @author kordenlu
// @创建时间 2017/09/14 12:09
// 功能描述:

package yyp

import (
	"reflect"
	"runtime"
)

func newTypeEncoder32Str(t reflect.Type) encoderFunc {
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
		return write32String
	case reflect.Struct:
		return StructEncoder32Str

	//case reflect.Map:
	//	return newMapEncoder(t)
	case reflect.Slice:
		return SliceEncoder32Str
	//case reflect.Array:
	//	return newArrayEncoder(t)
	default:
		return unsupportedTypeEncoder
	}
}

func write32String(e *encodeState, value reflect.Value) error {
	if err := WriteU32(e, reflect.ValueOf(uint32(len(value.String())))); err != nil {
		return err
	}
	_, err := e.WriteString(value.String())
	return err
}

func Marshal32Str(v interface{}) ([]byte, error) {
	e := &encodeState{}
	err := e.marshal32str(v)
	if err != nil {
		return nil, err
	}
	return e.Bytes(), nil
}

func (e *encodeState) marshal32str(v interface{}) (err error) {
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
	e.reflectValue32Str(reflect.ValueOf(v))
	return nil
}

func (e *encodeState) reflectValue32Str(v reflect.Value) {
	valueEncoder32Str(v)(e, v)
}

func valueEncoder32Str(v reflect.Value) encoderFunc {
	if !v.IsValid() {
		return invalidValueEncoder
	}
	return newTypeEncoder32Str(v.Type())
}

func StructEncoder32Str(e *encodeState, v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		e.reflectValue32Str(v.Field(i))
	}
	return nil
}

func SliceEncoder32Str(e *encodeState, v reflect.Value) error {
	//序列化长度uint32_t
	err := e.marshal32str(uint32(v.Len()))
	if err != nil {
		return err
	}
	//从头到尾依次序列数据
	for i := 0; i < v.Len(); i++ {
		e.reflectValue32Str(v.Index(i))
	}
	return nil
}
