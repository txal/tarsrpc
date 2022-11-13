package main

import (
	"log"
	"time"

	"code.com/tars/goframework/kissgo/yyp"
	"code.com/tars/goframework/kissgo/yypclient"
)

type PQueryUserAttentionListReq struct {
	Uid     uint32
	AppData string
	Extend  map[string]string
}

type PQueryUserAttentionListRsp struct {
	Result        uint32
	Uid           uint32
	AppData       string
	Attentionlist map[uint32]string
	Extend        map[string]string
}

func createRequest() []byte {
	var req PQueryUserAttentionListReq
	req.Uid = 1304934619
	req.AppData = "mobAttentionLite"
	req.Extend = make(map[string]string)
	reqbuf, err := yyp.Marshal(req)
	if err != nil {
		log.Panicf("yyp.Marshal(): %v", err)
	}

	//log.Printf("reqbuf len=%d", len(reqbuf))
	//fmt.Printf("%s", hex.Dump(reqbuf))
	return reqbuf
}

func parseResponse(rspbuf []byte) {
	//log.Printf("rspbuf len=%d", len(rspbuf))
	//fmt.Printf("%s", hex.Dump(rspbuf))

	var rsp PQueryUserAttentionListRsp
	err := yyp.Unmarshal(rspbuf, &rsp)
	if err != nil {
		log.Panicf("yyp.Unmarshal(): %v", err)
	}
	log.Printf("rsp %+v", rsp)
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	s := yypclient.Client{
		S2SName: "attentionlist",
		Groups:  []int{17610035},
		Timeout: 5000,
		//Address: &addr{network: "tcp", str: "10.0.2.4:8000"},
	}
	s.Setup()

	// Let s2s run
	time.Sleep(1 * time.Second)

	for {
		time.Sleep(1 * time.Second)
		rspbuf, err := s.DoRoundtrip(5<<8|225, 6<<8|225, createRequest())
		if err != nil {
			log.Printf("s.DoRoundtrip(): %v", err)
			continue
		}
		parseResponse(rspbuf)
	}
}
