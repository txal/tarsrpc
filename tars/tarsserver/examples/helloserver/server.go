package main

import (
	"encoding/binary"
	"time"

	"code.com/tars/goframework/tars/tarsserver"
)

type MyServer struct{}

func (s *MyServer) Invoke(req []byte) (rsp []byte, err error) {
	rsp = make([]byte, 4)
	rsp = append(rsp, []byte("Hello ")...)
	rsp = append(rsp, req...)
	binary.BigEndian.PutUint32(rsp[:4], uint32(len(rsp)))
	return
}

func (s *MyServer) ParsePackage(buff *[]byte) (pkg []byte, ret int) {
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
	pkg = (*buff)[4:length]
	*buff = (*buff)[length:]
	return pkg, tarsserver.PACKAGE_FULL
}

func (s *MyServer) InvokeTimeout(pkg []byte) ([]byte, error) {
	payload := []byte("timeout")
	ret := make([]byte, 4+len(payload))
	binary.BigEndian.PutUint32(pkg[:4], uint32(len(ret)))
	copy(pkg[4:], payload)
	return ret, nil
}

func main() {
	conf := &tarsserver.TarsServerConf{
		Proto:         "tcp",
		Address:       "localhost:3333",
		MaxAccept:     500,
		MaxInvoke:     20,
		AcceptTimeout: time.Millisecond * 500,
		ReadTimeout:   time.Millisecond * 100,
		WriteTimeout:  time.Millisecond * 100,
		HandleTimeout: time.Millisecond * 6000,
		IdleTimeout:   time.Millisecond * 600000,
	}
	s := MyServer{}
	svr := tarsserver.NewTarsServer(&s, conf)
	svr.Serve()
}
