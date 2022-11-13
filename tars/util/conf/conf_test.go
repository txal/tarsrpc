package conf

import (
	"fmt"
	"strconv"
	"testing"

	"code.com/tars/goframework/kissgo/appzaplog"
)

func TestConf(t *testing.T) {
	f1 := NewConf("./MMGR.TestServer.conf")
	//f1 := NewConf("Sumeru.HostApi.config.conf")
	d := f1.GetDomain("/tars/application/server")
	if len(d) != 2 {
		t.Error("/tars/application/server failed")
	}
	t.Logf("server map:%v", d)

	domaintaf := f1.GetDomain("/taf")
	if len(domaintaf) != 5 {
		t.Error("GetDomain /tars failed", domaintaf)
	}

	d2 := f1.GetString("/tars/application/server<node>")
	if d2 != "Docker.DCNode.ServerObj@tcp -h 100.97.11.84 -p 9931 -t 180000" {
		t.Error("server<node> failed", d2)
	}

	d3 := f1.GetString("/tars/application/client<locator>")
	if d3 != "Docker.DockerRegistry.QueryObj@tcp -h 10.173.5.41 -p 9903  -t 50000:tcp -h 100.107.177.31 -p 9903 -t 50000" {
		t.Error("client<locator> failed")
	}

	d4 := f1.GetInt("/tars/application/client<sample-rate>")
	if d4 != 1000 {
		t.Error("client<sample-rate> failed", d4)
	}

	d5 := f1.GetMap("/tars/application/server")
	if d5["node"] != "Docker.DCNode.ServerObj@tcp -h 100.97.11.84 -p 9931 -t 180000" {
		t.Error("node failed")
	}

	if netthread := d5["netthread"]; netthread == "" {
		t.Error("no netthread found")
	}

	domainserver := f1.GetDomain("/taf/application/server/Sumeru.HostApi.iprunObjAdapter")
	t.Logf("domainserver:%v", domainserver)

	d6 := f1.GetString("/tars/application/server/MMGR.NetClip.clipServerObjAdapter<servant>")
	if d6 != "MMGR.NetClip.clipServerObj" {
		t.Error("should be empty", d6)
	}

	// loglevel
	servermap := f1.GetMap("/tars/application/server")
	if servermap["logLevel"] != "DEBUG" {
		t.Error("loglevel:", servermap["loglevel"])
	}

	if netthread := f1.GetString("/tars/application/server<netthread>"); netthread == "" {
		t.Error("empty netthread:", netthread)
		intNetThread, err := strconv.Atoi(netthread)
		if err != nil {
			t.Errorf("intNetThread:%v failed", intNetThread)
		}
	}
}

func TestNewConf(t *testing.T) {
	appzaplog.InitAppLog()
	conf := NewConf("MMGR.TestServer.conf")
	m := conf.GetMap("/tars/application/client")
	for k, v := range m {
		fmt.Println(k, v)
	}

	m = conf.GetMap("/tars/application/server")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
