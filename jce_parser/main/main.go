package main

import (
	_ "bytes"
	"encoding/hex"
	"fmt"
	"yytars/jce_parser/gojce"
	"yytars/jce_parser/jce/include/Test"
)

// type B struct {
// 	a int
// 	f float32
// }

// type A struct {
// 	a int
// 	b B
// }

// type TestInfo struct {
// 	iBegin int
// 	b      bool
// 	si     int16
// 	by     byte
// 	ii     int
// 	li     int64
// 	f      float32
// 	d      float64
// 	s      string
// 	vi     []int
// 	mi     map[int]string
// 	aa     A
// 	iend   int
// 	vb     []byte
// 	vi2    []A
// 	mi2    map[int]A
// 	uii    uint32
// 	msv    map[string][]A
// 	vf     []float32
// }

// func (b *B) WriteTo(os gojce.JceOutputStream) error {
// 	var err error
// 	if err = os.Write(reflect.ValueOf(&b.a), 1); err != nil {
// 		fmt.Printf("err: %s", err.Error())
// 		return nil
// 	}

// 	if err = os.Write(reflect.ValueOf(&b.f), 2); err != nil {
// 		fmt.Printf("err: %s", err.Error())
// 		return nil
// 	}
// 	return nil
// }

// func (b *B) ReadFrom(is gojce.JceInputStream) error {
// 	var err error
// 	var i interface{}
// 	i, err = is.Read(reflect.TypeOf(b.a), 1, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	b.a = i.(int)

// 	i, err = is.Read(reflect.TypeOf(b.f), 2, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	b.f = i.(float32)
// 	return nil
// }

// func (a *A) WriteTo(os gojce.JceOutputStream) error {
// 	var err error
// 	if err = os.Write(reflect.ValueOf(&a.a), 1); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&a.b), 2); err != nil {
// 		panic(err)
// 	}
// 	return nil
// }

// func (a *A) ReadFrom(is gojce.JceInputStream) error {
// 	var err error
// 	var i interface{}
// 	i, err = is.Read(reflect.TypeOf(a.a), 1, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	a.a = i.(int)

// 	i, err = is.Read(reflect.TypeOf(a.b), 2, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	a.b = i.(B)

// 	return nil
// }

// func (t *TestInfo) WriteTo(os gojce.JceOutputStream) error {
// 	var err error
// 	if err = os.Write(reflect.ValueOf(&t.iBegin), 1); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.b), 2); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.si), 3); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.by), 4); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.ii), 5); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.li), 6); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.f), 7); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.d), 8); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.s), 9); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.vi), 10); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.mi), 11); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.aa), 12); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.iend), 13); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.vb), 14); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.vi2), 15); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.mi2), 16); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.uii), 17); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.msv), 18); err != nil {
// 		panic(err)
// 	}

// 	if err = os.Write(reflect.ValueOf(&t.vf), 19); err != nil {
// 		panic(err)
// 	}
// 	return nil
// }

// func (t *TestInfo) ReadFrom(is gojce.JceInputStream) error {
// 	var err error
// 	var i interface{}
// 	i, err = is.Read(reflect.TypeOf(t.iBegin), 1, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.iBegin = i.(int)

// 	i, err = is.Read(reflect.TypeOf(t.b), 2, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.b = i.(bool)

// 	i, err = is.Read(reflect.TypeOf(t.si), 3, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.si = i.(int16)

// 	i, err = is.Read(reflect.TypeOf(t.by), 4, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.by = i.(byte)

// 	i, err = is.Read(reflect.TypeOf(t.ii), 5, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.ii = i.(int)

// 	i, err = is.Read(reflect.TypeOf(t.li), 6, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.li = i.(int64)

// 	i, err = is.Read(reflect.TypeOf(t.f), 7, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.f = i.(float32)

// 	i, err = is.Read(reflect.TypeOf(t.d), 8, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.d = i.(float64)

// 	i, err = is.Read(reflect.TypeOf(t.s), 9, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.s = i.(string)

// 	i, err = is.Read(reflect.TypeOf(t.vi), 10, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.vi = i.([]int)

// 	i, err = is.Read(reflect.TypeOf(t.mi), 11, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.mi = i.(map[int]string)

// 	i, err = is.Read(reflect.TypeOf(t.aa), 12, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.aa = i.(A)

// 	i, err = is.Read(reflect.TypeOf(t.iend), 13, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.iend = i.(int)

// 	i, err = is.Read(reflect.TypeOf(t.vb), 14, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.vb = i.([]byte)

// 	i, err = is.Read(reflect.TypeOf(t.vi2), 15, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.vi2 = i.([]A)

// 	i, err = is.Read(reflect.TypeOf(t.mi2), 16, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.mi2 = i.(map[int]A)

// 	i, err = is.Read(reflect.TypeOf(t.uii), 17, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.uii = i.(uint32)

// 	i, err = is.Read(reflect.TypeOf(t.msv), 18, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.msv = i.(map[string][]A)

// 	i, err = is.Read(reflect.TypeOf(t.vf), 19, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.vf = i.([]float32)

// 	return nil
// }

func main() {
	//test1 := Test.TestInfo{}
	test1 := Test.TestInfo{
		M_ibegin: 1,
		M_b:      false,
		M_si:     255,
		M_by:     5,
		M_ii:     100000,
		M_li:     999999999,
		M_f:      1.3234,
		M_d:      2.94387691734,
		M_s:      "hello, world!",
		M_vi:     []int32{4, 56, 23},
		M_mi:     make(map[int32]string),
		M_aa:     Test.A{M_a: 123, M_b: Test.B{M_a: 456, M_f: 2.223}},
		M_iend:   4325,
		M_vb:     []byte{1, 5, 6, 2},
		M_vi2: []Test.A{
			Test.A{M_a: 1, M_b: Test.B{M_a: 222, M_f: 55.5}},
			Test.A{M_a: 2, M_b: Test.B{M_a: 333, M_f: 66.5}},
			Test.A{M_a: 3, M_b: Test.B{M_a: 444, M_f: 77.55}},
		},
		M_mi2: make(map[int32]Test.A),
		M_uii: 435346895,
		M_msv: make(map[string][]Test.A),
		M_vf:  []float32{123.53452, 4345.6546},
		M_msb: make(map[string]bool),
	}
	test1.M_mi[1] = "is"
	test1.M_mi[3] = "this"
	test1.M_mi[5] = "a"
	test1.M_mi[4] = "test"
	test1.M_mi2[1] = Test.A{M_a: 4, M_b: Test.B{M_a: 555, M_f: 234.66}}
	test1.M_mi2[2] = Test.A{M_a: 5, M_b: Test.B{M_a: 5899, M_f: 435.2}}
	test1.M_msv["kgjfdng"] = []Test.A{
		Test.A{M_a: 6, M_b: Test.B{M_a: 31, M_f: 77.7}},
		Test.A{M_a: 7, M_b: Test.B{M_a: 44, M_f: 88.8}},
	}
	test1.M_msv["ggg"] = []Test.A{
		Test.A{M_a: 8, M_b: Test.B{M_a: 1111, M_f: 89}},
		Test.A{M_a: 9, M_b: Test.B{M_a: 2222, M_f: 90}},
	}
	test1.M_msb["sss"] = false
	test1.M_msb["sss2"] = true

	os := gojce.NewOutputStream()
	test1.M_vi2 = nil
	test1.M_vi = nil
	//fmt.Printf("%+v\n", test1.M_vi2)
	if err := test1.WriteTo(os); err != nil {
		fmt.Printf("%+v", err)
	}
	bs := os.ToBytes()
	fmt.Printf("bs=\n%s\n", hex.Dump(bs))

	var test2 Test.TestInfo
	is := gojce.NewInputStream(bs)
	test2.ReadFrom(is)
	fmt.Printf("test2=%+v\n", test2)

	// buf := bytes.NewBuffer(nil)
	// ds := gojce.NewDisplayer(buf, 0)
	// test1.Display(ds)
	// fmt.Println(buf.String())

	en := gojce.NewJceJsonEncoder()
	js, err := test2.WriteJson(en)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))
}
