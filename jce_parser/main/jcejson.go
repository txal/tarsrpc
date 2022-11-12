package main

import (
	_ "bytes"
	"fmt"
	"yytars/jce_parser/gojce"
	"yytars/jce_parser/jce/include/Test"
	"reflect"
)

func main() {
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

	en := gojce.NewJceJsonEncoder()
	js, err := test1.WriteJson(en)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))

	var test2 Test.TestInfo
	de := gojce.NewJceJsonDecoder(js)
	err = de.DecodeJSON(reflect.ValueOf(&test2))
	if err != nil {
		panic(err)
	}

	en = gojce.NewJceJsonEncoder()
	js, err = test2.WriteJson(en)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))
}
