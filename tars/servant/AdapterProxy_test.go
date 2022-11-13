package servant

import (
	"reflect"
	"testing"

	end "code.com/tars/goframework/jce/servant/taf"
	"code.com/tars/goframework/jce/taf"
	"code.com/tars/goframework/jce_parser/gojce"
)

func TestAdapterProxy(t *testing.T) {
	comm := NewPbCommunicator()
	node := end.EndpointF{
		Host:        "10.125.228.4",
		Port:        10240,
		Timeout:     0,
		Istcp:       0,
		Grid:        0,
		Groupworkid: 0,
		Grouprealid: 0,
		Qos:         0,
		BakFlag:     0,
		GridFlag:    0,
		Weight:      0,
		WeightType:  0,
		Cpuload:     0,
		Sampletime:  0,
	}
	adapter := NewAdapterProxy(&node, comm)

	for z := 1; z < 10; z++ {
		oe := gojce.NewOutputStream()
		x, y := 10, z
		oe.Write(reflect.ValueOf(&x), 1)
		oe.Write(reflect.ValueOf(&y), 2)

		req := taf.RequestPacket{
			IVersion:     1,
			CPacketType:  0,
			IRequestId:   int32(z),
			SServantName: "Docker.DCNode.ServerObj",
			SFuncName:    "keepAlive",
			ITimeout:     3000,
			SBuffer:      oe.ToBytes(),
			Context:      map[string]string{"test": "test"},
			Status:       map[string]string{"test": "test"},
		}
		adapter.Send(&req)
		//i, err := adapter.OutputHandle(&req)
		//if err != nil {
		//	t.Fatalf(err.Error())
		//}
		//if i > 0 {
		//	t.Log(i)
		//}
		//
		//buf := bytes.NewBuffer(nil)
		//ds := gojce.NewDisplayer(buf, 0)
		//resp := <-adapter.Resp
		//resp.Display(ds)
		//t.Log(resp.IRequestId)
		//fmt.Println(resp.SBuffer)
		//
		//is := gojce.NewInputStream(resp.SBuffer)
		//
		////netclip.jce
		//var r, j int32
		//
		////r, err := is.Read(reflect.TypeOf(r), 0 ,true)
		//a, err := is.Read(reflect.TypeOf(r), 0, true)
		//if err != nil {
		//	t.Fatalf(err.Error())
		//}
		//b, err := is.Read(reflect.TypeOf(j), 3, true)
		//if err != nil {
		//	t.Fatalf(err.Error())
		//}
		//var e int32
		//e = a.(int32)
		//fmt.Println(e, b)
	}

}
