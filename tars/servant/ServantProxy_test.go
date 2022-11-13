package servant

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"code.com/tars/goframework/jce_parser/gojce"
)

func TestServantProxy(t *testing.T) {

	oe := gojce.NewOutputStream()
	x, y := 10, 22
	oe.Write(reflect.ValueOf(&x), 1)
	oe.Write(reflect.ValueOf(&y), 2)

	s := NewServantProxy(NewPbCommunicator(), "Docker.DockerRegistry.QueryObj@tcp -h 10.125.228.4 -p 10240 -t 50000")

	resp, err := s.Taf_invoke(context.TODO(), 0, "taf_ping", oe.ToBytes())

	if err != nil {
		t.Error("Taf_invoke failed", err)
	}

	is := gojce.NewInputStream(resp.SBuffer)

	var r, j int32

	a, err := is.Read(reflect.TypeOf(r), 0, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
	b, err := is.Read(reflect.TypeOf(j), 3, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
	var e int32
	e = a.(int32)
	fmt.Println(e, b)
}
