package main

import (
	"log"
	"os"
	"time"

	"code.com/tars/goframework/kissgo/s2s"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("run as: %s SERVICE_NAME", os.Args[0])
	}

	err := s2s.InitS2s("ramcloud_test", "d10e28e567224404e7d93f35067fa54b6c4e8def35ed93cd33e4df6db5d3c5a0")
	if err != nil {
		log.Fatalf("init s2s fail: %v", err)
	}

	serviceName := os.Args[1]
	res, err := s2s.Subscribe(serviceName, 0)
	if err != nil {
		log.Fatalf("sub s2s fail: %v", err)
	}

	timeout := time.After(5 * time.Hour)
	for {
		select {
		case m := <-res:
			output(m)
			continue
		case <-timeout:
			os.Exit(0)
		}
	}
}
