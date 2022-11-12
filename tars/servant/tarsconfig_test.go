// @author kordenlu
// @创建时间 2018/02/09 19:30
// 功能描述:

package servant

import (
	"os"
	"testing"
)

//	var testconfig *serverConfig = &serverConfig{
//		Node:"tars.tarsnode.ServerObj@tcp -h 183.36.111.89 -p 19386 -t 60000",
//
// }
func TestMain(m *testing.M) {
	comm := startFrameWorkComm()
	srvcfg := &serverConfig{
		config:   "tars.tarsconfig.ConfigObj@tcp -h 58.215.138.213 -t 60000 -p 10001",
		App:      "HelloProtoTest",
		Server:   "HelloProto",
		BasePath: "/Users/o_o/develop/goprojects/src/tarsrpc/tars/testdata",
	}
	initTarConfig(comm, srvcfg, 5)
	os.Exit(m.Run())
}

func TestGetBaseConfPath(t *testing.T) {
	basepath, err := GetConfBasePath()
	if err != nil {
		t.Error("should be nil ", basepath)
	}

	basepath, err = GetConfBasePath()
	if err != nil || basepath != "/Users/o_o/develop/goprojects/src/tarsrpc/tars/testdata/" {
		t.Errorf("err:%v,basepath:%v", err, basepath)
	}
}

func TestSamecontent(t *testing.T) {
	// test same
	same, err := samecontent("../testdata/file1.json", "../testdata/file2.json")
	if err != nil {
		t.Error("samecontent failed", err)
	}
	if !same {
		t.Error("should be same")
	}

	// test diff, there is space in filediff :)
	same, err = samecontent("../testdata/file1.json", "../testdata/filediff.json")
	if err != nil || same {
		t.Errorf("err or should not be same")
	}

	// test empty oldfile
	same, err = samecontent("../testdata/file1.json", "../testdata/notexist.json")
	if err != nil || same {
		t.Errorf("err or should not be same")
	}
}

func TestIndex2file(t *testing.T) {
	basename := "whatsapp/1/2/abc"
	if indexname := index2file(basename, 10); indexname != basename+".10.bak" {
		t.Error("should be eq")
	}
}

func TestGetRemoteFile(t *testing.T) {
	if defaultTarConfig == nil {
		t.Error("initTarConfig failed")
	}

	// what's fuck,, not exist file still return 0,nil
	newfile, err := defaultTarConfig.getRemoteFile("noexist", false)
	if err == nil {
		t.Error("getRemoteFile failed", err, newfile)
	}

	newfile, err = defaultTarConfig.getRemoteFile("whatapp", false)
	if err != nil {
		t.Error("getRemoteFile failed", err, newfile)
	}
	t.Logf("newfile:%v", newfile)
}

func TestAddConfig(t *testing.T) {
	if defaultTarConfig == nil {
		t.Error("initTarConfig failed")
	}
	result, err := defaultTarConfig.addConfig("whatapp", false)
	if err != nil {
		t.Error("addConfig failed", err)
	}
	t.Logf("result:%v", result)
}
