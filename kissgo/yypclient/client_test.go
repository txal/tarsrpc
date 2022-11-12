package yypclient

import (
	"encoding/hex"
	"fmt"
	"runtime"
	"tarsrpc/kissgo/yyp"
	"testing"
	"time"
)

func _func() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	function := f.Name()
	file, line := f.FileLine(pc[0])
	return fmt.Sprintf("%s:%d [%s]", file, line, function)
}

func createRequest() []byte {
	type PGreeting struct {
		Callid   string
		Greeting string
		StrList  []string
		MList    map[uint32]string
	}
	var req PGreeting
	req.Callid = "golang"
	req.Greeting = "Hello LiLei"
	req.StrList = []string{"item0", "item1", "item2"}
	req.MList = map[uint32]string{
		3: "item2",
		1: "item0",
		2: "item1",
	}
	reqbuf, err := yyp.Marshal(req)
	if err != nil {
		fmt.Println(_func(), err)
		return nil
	}
	fmt.Println(_func(), "req buf len=", len(reqbuf))
	fmt.Printf("%s", hex.Dump(reqbuf))
	return reqbuf
}

func parseResponse(rspbuf []byte) {
	fmt.Println(_func(), "rsp buf len=", len(rspbuf))
	fmt.Printf("%s", hex.Dump(rspbuf))
	type PGreetingRes struct {
		Callid   string
		Greeting string
		StrList  []string
		MList    map[uint32]string
	}
	var rsp PGreetingRes
	err := yyp.Unmarshal(rspbuf, &rsp)
	if err != nil {
		fmt.Println(_func(), err)
		return
	}
	fmt.Println(_func(), "rsp =", rsp)
}

func TestDoRoundtrip(t *testing.T) {
	s := Client{
		S2SName: "lilei",
		Groups:  []int{17610035},
		Timeout: 5000,
		Address: []addr{
			{network: "tcp", str: "10.0.2.4:8000"},
		},
	}
	s.Setup()

	time.Sleep(1 * time.Second)

	for {
		time.Sleep(1 * time.Second)
		rspbuf, err := s.DoRoundtrip(1<<8|255, 2<<8|255, createRequest())
		if err != nil {
			fmt.Println(_func(), err)
			continue
		}
		parseResponse(rspbuf)
	}
}

func TestDoRoundtripRetry(t *testing.T) {
	s := Client{
		S2SName: "lilei",
		Groups:  []int{17610035},
		Timeout: 5000,
		Address: []addr{
			{network: "tcp", str: "10.0.2.4:8000"},
		},
	}
	s.Setup()

	time.Sleep(1 * time.Second)

	for {
		time.Sleep(1 * time.Second)
		rspbuf, err := s.DoRoundtripRetry(1<<8|255, 2<<8|255, createRequest())
		if err != nil {
			fmt.Println(_func(), err)
			continue
		}
		parseResponse(rspbuf)
	}
}
