package servant

import (
	"net"
	"reflect"
	"strconv"
	"testing"

	"code.com/tars/goframework/jce/taf"
	"code.com/tars/goframework/jce_parser/gojce"
	"code.com/tars/goframework/tars/util/endpoint"
)

func TestEndpointManager(t *testing.T) {
	end := NewEndpointManager("Docker.DockerRegistry.QueryObj@tcp -h 10.125.228.4 -p 10240 -t 50000", &Communicator{})
	msg := &Message{}
	adapter := end.SelectAdapterProxy(msg)

	oe := gojce.NewOutputStream()
	x, y := 10, 11
	oe.Write(reflect.ValueOf(&x), 1)
	oe.Write(reflect.ValueOf(&y), 2)

	req := &taf.RequestPacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   11,
		SServantName: "Docker.DCNode.ServerObj",
		SFuncName:    "keepAlive",
		ITimeout:     3000,
		SBuffer:      oe.ToBytes(),
		Context:      map[string]string{"test": "test"},
		Status:       map[string]string{"test": "test"},
	}
	msg.Req = req
	err := adapter.Send(req)
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestEmptyMapGet(t *testing.T) {
	type Userinfo struct {
		Name string
	}
	var localmap map[string]*Userinfo = make(map[string]*Userinfo)
	localmap["test"] = &Userinfo{
		Name: "test",
	}
	get := func(name string) *Userinfo {
		return localmap[name]
	}

	if name := get("name"); name != nil {
		t.Error("should be empty", name)
	}

	if test := get("test"); test == nil || test.Name != "test" {
		t.Error("should be test", test)
	}
}

func BenchmarkMapWithObjKey(b *testing.B) {
	b.StopTimer()
	var objmap map[endpoint.Endpoint]int = make(map[endpoint.Endpoint]int)
	insertmap := func(num int) {
		for i := 0; i < num; i++ {
			end := endpoint.Endpoint{
				Host:   "127.0.0.1",
				Port:   int32(i),
				IPPort: net.JoinHostPort("127.0.0.1", strconv.FormatInt(int64(i), 10)),
			}
			objmap[end] = i
		}
	}
	insertmap(10)
	key := endpoint.Endpoint{
		Host:   "127.0.0.1",
		Port:   int32(5),
		IPPort: net.JoinHostPort("127.0.0.1", strconv.FormatInt(int64(5), 10)),
	}

	b.StartTimer()
	for k := 0; k < b.N; k++ {
		if _, ok := objmap[key]; !ok {
			b.Error("should be ok")
		}
	}
}

func BenchmarkMapWithStringKey(b *testing.B) {
	b.StopTimer()
	var objmap map[string]int = make(map[string]int)
	insertmap := func(num int) {
		for i := 0; i < num; i++ {
			end := endpoint.Endpoint{
				Host:   "127.0.0.1",
				Port:   int32(i),
				IPPort: net.JoinHostPort("127.0.0.1", strconv.FormatInt(int64(i), 10)),
			}
			objmap[end.IPPort] = i
		}
	}
	insertmap(10)
	key := endpoint.Endpoint{
		Host:   "127.0.0.1",
		Port:   int32(5),
		IPPort: net.JoinHostPort("127.0.0.1", strconv.FormatInt(int64(5), 10)),
	}

	b.StartTimer()
	for k := 0; k < b.N; k++ {
		if _, ok := objmap[key.IPPort]; !ok {
			b.Error("should be ok")
		}
	}
}
