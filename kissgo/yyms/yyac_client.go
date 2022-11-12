package yyms

import (
	"net"
	"time"
	"syscall"
	"bytes"
	"encoding/json"
	"encoding/binary"
	"fmt"
	"path"
	"os"
)

const (
	Fid_Sun_Path = "/tmp/yymp.agent.sock"
	//Log_Sun_Path = "/tmp/yyac_worker_ant.api.sock"
	Max_Data_Len = 65477
	//LOGID_URI    = uint32(3000 << 8 | 11)
	FID_URI      = uint32(3001 << 8 | 11)
)

type YYMSClient struct{
	Pname      string
	Pid        int
	//FidSunAddr *net.UnixAddr
	//FidConn    *net.UnixConn
}

type AgentPkgHead struct {
	Len    uint32
	Uri    uint32
	Code   uint16
}

type AgentPkg struct {
	Log_id uint64
	Type   uint16
	Ts     uint64
	Data   string
}

type AgentMsg struct {
	FID   uint64 `json:"id"`
	SID   uint64 `json:"sid"`
	Alarm int32  `json:"alarm"`
	Pre   int32  `json:"pre"`
	Value int32  `json:"value"`
	Pname string `json:"pname"`
	Pid   int32  `json:"pid"`
	Msg   string `json:"msg"`
}

var (
	yymsclient *YYMSClient
)
func init()  {
	yymsclient = &YYMSClient{
		Pname:path.Base(os.Args[0]),
			Pid:syscall.Getpid(),
	}
}

func SendAlarm(fid,sid uint64, msg string) bool {
	if yymsclient == nil{
		return false
	}
	unixaddr,err := net.ResolveUnixAddr("unixgram",Fid_Sun_Path)
	if err != nil {
		return false
	}
	c,err := net.DialUnix("unixgram",nil,unixaddr)
	if err != nil {
		return false
	}
	defer c.Close()
	return yymsclient.sendAlarm(c,fid,sid,msg)
}

func (m *YYMSClient) toJsonAlarmData(feature_id uint64, strategy_id uint64, early_warning int32, msg string) ([]byte, error){
	agentmsg := &AgentMsg{
		FID:   feature_id,
		SID:   strategy_id,
		Alarm: 1,
		Pre:   early_warning,
		Pname: m.Pname,
		Pid:   int32(m.Pid),
		Msg:   msg,
	}
	return json.Marshal(agentmsg)
}

func (m *YYMSClient) FidUnixDomainSocketSend(c *net.UnixConn,buf *bytes.Buffer) (bool) {
	b, err := c.Write(buf.Bytes())
	if b != buf.Len() || err != nil {
		fmt.Printf("Failed to FidSend, expect %d > %d, err:%v", buf.Len(), b, err.(*net.OpError).Error())
		return false;
	}
	return true
}

func (m *YYMSClient) AgentClientSend(c *net.UnixConn,str string, uri uint32, log_id uint64) (bool){
	if len(str) == 0 || len(str) >= Max_Data_Len {
		return false;
	}
	head := &AgentPkgHead{0, uri, 200}
	pkg  := &AgentPkg{log_id, 0 , uint64(time.Now().Unix()), str}
	pkgbuf:= new (bytes.Buffer)
	if err := Marshal(&pkg, pkgbuf, binary.LittleEndian, BlobLength16); err != nil {
		fmt.Printf("AgentClientSend.marshal pkg fail:%v", err)
		return false
	}
	head.Len = 10 + uint32(pkgbuf.Len())
	buf  := new (bytes.Buffer)
	if err := Marshal(&head, buf, binary.LittleEndian, BlobLength16); err != nil {
		fmt.Printf("AgentClientSend.marshal head fail:%v", err)
		return false
	}
	buf.Write(pkgbuf.Bytes())
	return m.FidUnixDomainSocketSend(c,buf)
}


func (m *YYMSClient) sendAlarm(c *net.UnixConn,feature_id uint64, strategy_id uint64, context string) (bool){
	bytes, err := m.toJsonAlarmData(feature_id, strategy_id, 0, context)
	if err != nil {
		fmt.Printf("sendAlarm.toJsonAlarmData fail:%v", err)
		return false
	}
	return m.AgentClientSend(c,string(bytes), FID_URI, 1)
}

