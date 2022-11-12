// @author kordenlu
// @创建时间 2017/08/17 17:29
// 功能描述:

package yyp

import (
	"fmt"
	"reflect"
	"testing"

	"math/rand"
	"strconv"
	"time"
)

/*
func TestReadbool(t *testing.T)  {
	data := []byte{1}
	stat := &DecodeState{}
	stat.init(data)
	num := stat.readBool()
	t.Logf("num:%v",num)
	if num{
		t.Error("should be false")
	}
}
*/
//add by zc

func Test_Unmarshal_uint16(t *testing.T) {
	testuint16 := uint16(123)
	marshaled, err := Marshal(testuint16)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res uint16
	Unmarshal(marshaled, &res)
	fmt.Println("raw data :", testuint16, "res :", res)
	if res != testuint16 {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_uint8(t *testing.T) {

	testuint8 := uint8(64)
	marshaled, err := Marshal(testuint8)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res uint8
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", testuint8, "res :", res)
	if res != testuint8 {
		t.Error("cannot unmarshal")
	}
}
func Test_Unmarshal_uint32(t *testing.T) {

	testuint32 := uint32(100000)
	marshaled, err := Marshal(testuint32)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res uint32
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", testuint32, "res :", res)
	if res != testuint32 {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_uint64(t *testing.T) {

	testuint64 := uint64(10002121200)
	marshaled, err := Marshal(testuint64)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res uint64
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", testuint64, "res :", res)
	if res != testuint64 {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_int8(t *testing.T) {

	testint8 := int8(127)
	marshaled, err := Marshal(testint8)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res int8
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", testint8, "res :", res)
	if res != testint8 {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_int16(t *testing.T) {

	testint16 := int16(1227)
	marshaled, err := Marshal(testint16)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res int16
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", testint16, "res :", res)
	if res != testint16 {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_int32(t *testing.T) {

	testint32 := int32(12227)
	marshaled, err := Marshal(testint32)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res int32
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", testint32, "res :", res)
	if res != testint32 {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_int64(t *testing.T) {

	testint64 := int64(1122227)
	marshaled, err := Marshal(testint64)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res int64
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", testint64, "res :", res)
	if res != testint64 {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_bool(t *testing.T) {

	testbool := bool(true)
	marshaled, err := Marshal(testbool)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res bool
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", testbool, "res :", res)
	if res != testbool {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_string(t *testing.T) {

	teststring := string("test unmarshal")
	marshaled, err := Marshal(teststring)

	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res string
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", teststring, "res :", res)
	if res != teststring {
		t.Error("cannot unmarshal")
	}
}
func Test_Unmarshal_struct_1(t *testing.T) {

	type Test_truct struct {
		Ui8 uint8
	}
	teststruct := Test_truct{Ui8: uint8(123)}

	marshaled, err := Marshal(teststruct)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res Test_truct
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", teststruct, "res :", res)
	if res != teststruct {
		t.Error("cannot unmarshal")
	}
}
func Test_Unmarshal_struct_2(t *testing.T) {

	type Test_truct struct {
		Ui8  uint8
		Ui16 uint16
	}
	teststruct := Test_truct{Ui8: uint8(123), Ui16: uint16(123)}

	marshaled, err := Marshal(teststruct)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res Test_truct
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", teststruct, "res :", res)
	if res != teststruct {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_struct_3(t *testing.T) {

	type Test_truct struct {
		Ui8  uint8
		Ui16 uint16
		Ui32 uint32
	}
	teststruct := Test_truct{Ui8: uint8(123), Ui16: uint16(1232), Ui32: uint32(238912)}

	marshaled, err := Marshal(teststruct)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res Test_truct
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", teststruct, "res :", res)
	if res != teststruct {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_struct_4(t *testing.T) {

	type Test_truct struct {
		Ui8  uint8
		Ui16 uint16
		Ui32 uint32
		Ui64 uint64
	}
	teststruct := Test_truct{Ui8: uint8(123), Ui16: uint16(1232), Ui32: uint32(238912), Ui64: uint64(1789789789)}

	marshaled, err := Marshal(teststruct)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res Test_truct
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", teststruct, "res :", res)
	if res != teststruct {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_struct_5(t *testing.T) {

	type Test_truct struct {
		Ui8   uint8
		Ui16  uint16
		Ui32  uint32
		Ui64  uint64
		Blean bool
	}
	teststruct := Test_truct{Ui8: uint8(123), Ui16: uint16(1232), Ui32: uint32(238912), Ui64: uint64(1789789789),
		Blean: true}

	marshaled, err := Marshal(teststruct)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res Test_truct
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", teststruct, "res :", res)
	if res != teststruct {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_struct_6(t *testing.T) {

	type Test_truct struct {
		Ui8   uint8
		Ui16  uint16
		Ui32  uint32
		Ui64  uint64
		Blean bool
		Str   string
	}
	teststruct := Test_truct{Ui8: uint8(123), Ui16: uint16(1232), Ui32: uint32(238912), Ui64: uint64(1789789789),
		Blean: true, Str: string("hello world")}

	marshaled, err := Marshal(teststruct)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res Test_truct
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", teststruct, "res :", res)
	if res != teststruct {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_struct_substruct(t *testing.T) {

	type Test_Sub struct {
		Ui32  uint32
		Ui64  uint64
		Blean bool
		Str   string
	}
	type Test_truct struct {
		Ui8   uint8
		Ui16  uint16
		Ui32  uint32
		Ui64  uint64
		Blean bool
		Str   string
		Sub   Test_Sub
	}
	teststruct := Test_truct{Ui8: uint8(123), Ui16: uint16(1232), Ui32: uint32(238912), Ui64: uint64(1789789789),
		Blean: true, Str: string("hello world"),
		Sub: Test_Sub{Ui32: uint32(323123), Ui64: uint64(23423566), Blean: true, Str: string("hello world")}}

	marshaled, err := Marshal(teststruct)
	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res Test_truct
	Unmarshal(marshaled, &res)
	//fmt.Printf("%v",marshaled);
	fmt.Println("raw data :", teststruct, "res :", res)
	if res != teststruct {
		t.Error("cannot unmarshal")
	}
}

func Test_Unmarshal_slice_1(t *testing.T) {

	testslice := []uint32{1, 2, 3, 4, 5, 6, 7}
	marshaled, err := Marshal(testslice)

	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res []uint32
	Unmarshal(marshaled, &res)

	//fmt.Println("marshaled len:",len(marshaled),"\n")
	fmt.Println("raw data :", testslice, "res :", res)

	if !reflect.DeepEqual(testslice, res) {
		t.Error(" not equal! or one is nil")
	}

}

func Test_Unmarshal_slice_2(t *testing.T) {

	testslice := []string{"this", "is", "a", "go", "program"}
	marshaled, err := Marshal(testslice)

	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res []string
	Unmarshal(marshaled, &res)

	//fmt.Println("marshaled len:",len(marshaled),"\n")
	fmt.Println("raw data :", testslice, "res :", res)

	if !reflect.DeepEqual(testslice, res) {
		t.Error(" not equal! or one is nil")
	}

}

func Test_Unmarshal_slice_3(t *testing.T) {
	type Test struct {
		U32 uint32
		U16 uint16
		U8  uint8
		Str string
		//Sli []uint64
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	test_slice_struct := make([]Test, 10)
	for i := 0; i < 10; i++ {
		test_slice_struct[i].U32 = uint32(r.Uint32())
		test_slice_struct[i].U16 = uint16(r.Uint32())
		test_slice_struct[i].U8 = uint8(r.Uint32())
		test_slice_struct[i].Str = "test data " + strconv.Itoa(i)
		//test_slice_struct[i].Sli = Append(test_slice_struct[i].Sli,uint64(i))
	}

	marshaled, err := Marshal(test_slice_struct)

	if err != nil {
		t.Error(" marshal failed:", err)
	}

	var res []Test

	Unmarshal(marshaled, &res)

	fmt.Println("raw data :", test_slice_struct, "res :", res)
	if !reflect.DeepEqual(test_slice_struct, res) {
		t.Error(" not equal! or one is nil")
	}

}

func Test_Unmarshal_slice_4(t *testing.T) {
	type Test struct {
		U32 uint32
		U16 uint16
		U8  uint8
		Str string
		Sli []uint64
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	test_slice_struct := make([]Test, 10)
	//插入数据
	for i := 0; i < 10; i++ {
		test_slice_struct[i].U32 = r.Uint32()
		test_slice_struct[i].U16 = uint16(r.Uint32())
		test_slice_struct[i].U8 = uint8(r.Uint32())
		test_slice_struct[i].Str = "test data " + strconv.Itoa(i)
		for j := 0; j < 10; j++ {
			test_slice_struct[i].Sli = append(test_slice_struct[i].Sli, uint64(r.Uint32()))
		}

	}

	marshaled, err := Marshal(test_slice_struct)

	if err != nil {
		t.Error(" marshal failed:", err)
	}
	//fmt.Printf("%v",marshaled);

	var res []Test

	Unmarshal(marshaled, &res)

	//fmt.Println("marshaled len:",len(marshaled),"\n")
	fmt.Println("raw data :", test_slice_struct, "res :", res)
	//比较数据
	if !reflect.DeepEqual(test_slice_struct, res) {
		t.Error(" not equal! or one is nil")
	}

}

/*

func Test_Unmarshal_fromcpp(t *testing.T)  {

	type Test_truct struct{
		Ui8 uint8
		Ui16 uint16
		Ui32 uint32
		Ui64 uint64
		Blean bool
		Str string
	}

	const dir = "/home/zhangcong/workspace/Serialization_test"

		teststruct := Test_truct{Ui8:uint8(123), Ui16:uint16(1232),Ui32 : uint32(238912), Ui64:uint64(1789789789),
								Blean:true,Str:string("hello world")}


		marshaled, err :=  Marshal(teststruct)
		if  err!=nil{
			t.Error(" marshal failed:",err)
		}
		//fmt.Printf("%v",marshaled);

		var res Test_truct
		Unmarshal(marshaled,&res)
		//fmt.Printf("%v",marshaled);
		fmt.Println("raw data :",teststruct,"res :",res);
		if res != teststruct {
			t.Error("cannot unmarshal")
		}
}



/*
func Test_Unmarshal_struct(t *testing.T)  {

	type Test_truct struct{
		Ui8 uint8
		Ui16 uint16
		Ui32 uint32
		Ui64 uint64
		Blean bool
		Str string
	}
		teststruct := Test_truct{Ui8:uint8(123), Ui16:uint16(1212),
								Ui32 : uint32(238912), Ui64:uint64(1789789789), Blean:true, Str:string("hello world")}


		marshaled, err :=  Marshal(teststruct)
		if  err!=nil{
			t.Error(" marshal failed:",err)
		}
		//fmt.Printf("%v",marshaled);

		var res Test_truct
		Unmarshal(marshaled,&res)
		//fmt.Printf("%v",marshaled);
		fmt.Println("raw data :",teststruct,"res :",res);
		if res != teststruct {
			t.Error("cannot unmarshal")
		}
}

*/
//end add by zc
