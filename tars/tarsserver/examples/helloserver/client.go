package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"code.com/tars/goframework/tars/tarsserver"
)

type MyClient struct {
}

func (c *MyClient) Recv(pkg []byte) {
	fmt.Println("recv", string(pkg))
}
func (c *MyClient) ParsePackage(buff *[]byte) (pkg []byte, ret int) {
	if len(*buff) < 4 {
		return nil, tarsserver.PACKAGE_LESS
	}
	length := binary.BigEndian.Uint32((*buff)[:4])

	if length > 10485760 || len((*buff)) > 10485760 { // 10MB
		return nil, tarsserver.PACKAGE_ERROR
	}
	if len((*buff)) < int(length) {
		return nil, tarsserver.PACKAGE_LESS
	}
	pkg = append(pkg, (*buff)[4:length]...) // must deep copy
	*buff = (*buff)[length:]
	return pkg, tarsserver.PACKAGE_FULL
}

func getMsg(name string) []byte {
	payload := []byte(name)
	pkg := make([]byte, 4+len(payload))
	binary.BigEndian.PutUint32(pkg[:4], uint32(len(pkg)))
	copy(pkg[4:], payload)
	return pkg
}

func main() {
	cp := &MyClient{}
	conf := &tarsserver.TarsClientConf{
		Proto:        "tcp",
		NumConnect:   1,
		QueueLen:     10000,
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Millisecond * 100,
		WriteTimeout: time.Millisecond * 1000,
	}
	client := tarsserver.NewTarsClient("localhost:3333", cp, conf)

	name := "Bob"
	for i := 0; i < 5; i++ {
		msg := getMsg(name + strconv.Itoa(i))
		client.Send(msg)
	}

	time.Sleep(time.Second * 1)
	client.Close()
}
