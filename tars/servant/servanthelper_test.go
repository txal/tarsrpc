// @author kordenlu
// @创建时间 2018/02/08 20:31
// 功能描述:

package servant

import "testing"

func TestFullobjnameFull(t *testing.T)  {
	shortobjname := "App.Server.Hello"
	fullobjname,err := fullObjName(shortobjname)
	if fullobjname != shortobjname || err != nil{
		t.Errorf("shortobjname:%v,fullname:%v,err:%v",shortobjname, fullobjname,err)
	}
}

func TestFullobjnameShort(t *testing.T)  {
	shortobjname := "Hello"
	objname,err := fullObjName(shortobjname)
	if err != NilServerConfig{
		t.Errorf("shortobjname:%v,fullname:%v,err:%v",shortobjname, objname,err)
	}

	// set serverconfig
	svrCfg = &serverConfig{
	}

	objname,err = fullObjName(shortobjname)
	if err != EmptyAppOrServerName{
		t.Errorf("shortobjname:%v,fullname:%v,err:%v",shortobjname, objname,err)
	}

	svrCfg.App, svrCfg.Server = "app","server"
	objname,err = fullObjName(shortobjname)
	if err != nil || objname != "app.server."+shortobjname{
		t.Errorf("shortobjname:%v,fullname:%v,err:%v",shortobjname, objname,err)
	}
}