package servant

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"testing"

	"code.com/tars/goframework/jce/taf"
	"code.com/tars/goframework/jce_parser/gojce"
)

func TestObjectProxy(t *testing.T) {
	comm := NewPbCommunicator()
	obj := NewObjectProxy(comm, "Docker.DockerRegistry.QueryObj@tcp -h 10.125.228.4 -p 10240 -t 50000")
	ObjTest(obj)
}

func ObjTest(obj *ObjectProxy) {
	oe := gojce.NewOutputStream()
	x, y := 10, 123
	oe.Write(reflect.ValueOf(&x), 1)
	oe.Write(reflect.ValueOf(&y), 2)

	req := taf.RequestPacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   int32(12),
		SServantName: "Docker.DCNode.ServerObj",
		SFuncName:    "keepAlive",
		ITimeout:     3000,
		SBuffer:      oe.ToBytes(),
		Context:      map[string]string{"test": "test"},
		Status:       map[string]string{"test": "test"},
	}

	var msg Message
	msg.Req = &req
	obj.Invoke(context.TODO(), &msg, 3000)

	buf := bytes.NewBuffer(nil)
	ds := gojce.NewDisplayer(buf, 0)
	resp := msg.Resp
	resp.Display(ds)
	fmt.Println(resp.SBuffer)
}
