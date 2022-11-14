package servant

import (
	"fmt"
	"testing"
	"time"
	"code.com/tars/goframework/jce/servant/taf"
)

func Test_StatFHelper(t *testing.T) {
	fmt.Println("test")
	end := new(StatFHelper)
	fmt.Println("-1")
	req := StatInfo{}
	fmt.Println("0")
	//comm := newCommunicator()
	fmt.Println("init")
	end.Init(startFrameWorkComm(), "taf.tafstat.StatObj@tcp -h 100.65.11.78 -p 10032 -t 60000")
	go end.Run()
	fmt.Println("1")
	var _statInfo taf.StatMicMsgHead
	_statInfo.MasterName = "Test.grayserver"
	_statInfo.SlaveName = "Docker.DockerRegistry"
	_statInfo.InterfaceName = "getModuleRule"
	_statInfo.MasterIp = "100.97.13.160"
	_statInfo.SlaveIp = "100.115.10.164"
	_statInfo.SlavePort = 9910
	_statInfo.ReturnValue = 0
	_statInfo.TafVersion = "1.0"
	_statInfo.SSlaveContainer = ""
	_statInfo.SMasterContainer = ""
	req.Head = _statInfo
	fmt.Println("2")
	var _statBody taf.StatMicMsgBody
	_statBody.Count = 2
	_statBody.TimeoutCount = 1
	_statBody.ExecCount = 1
	_statBody.TotalRspTime = 621
	_statBody.MaxRspTime = 21
	_statBody.MinRspTime = 3
	req.Body = _statBody
	fmt.Println("going to, master name: ", req.Head.MasterName)
	end.ReportMicMsg(req)
	time.Sleep(1 * time.Second)
	fmt.Println("1 Over")

	end.ReportMicMsg(req)
	time.Sleep(1 * time.Second)
	fmt.Println("2 Over")

	end.ReportMicMsg(req)
	time.Sleep(1 * time.Second)
	fmt.Println("3 Over")

	end.ReportMicMsg(req)
	time.Sleep(1 * time.Second)
	fmt.Println("4 Over")

	end.ReportMicMsg(req)
	time.Sleep(1 * time.Second)
	fmt.Println("5 Over")

	end.ReportMicMsg(req)
	time.Sleep(1 * time.Second)
	fmt.Println("6 Over")

	end.ReportMicMsg(req)
	time.Sleep(1 * time.Second)
	fmt.Println("7 Over")

	end.ReportMicMsg(req)
	time.Sleep(1 * time.Second)
	fmt.Println("8 Over")

	time.Sleep(5 * time.Second)
}

func TestMapUsage(t *testing.T)  {
	var testmap map[taf.StatMicMsgHead]taf.StatMicMsgBody = make(map[taf.StatMicMsgHead]taf.StatMicMsgBody)
	head := taf.StatMicMsgHead{
		MasterName:"test1",
	}
	testmap[head] = taf.StatMicMsgBody{
		Count:100,
	}
	if v,ok := testmap[head];ok{
		v.Count = v.Count +1
		testmap[head] = v
	}
	if testmap[head].Count != 101 {
		t.Errorf("should be 101,actual:%v",testmap[head].Count)
	}
}