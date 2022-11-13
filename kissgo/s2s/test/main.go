package main

import (
	"flag"
	"fmt"
	"sync"

	"code.com/tars/goframework/kissgo/s2s"
)

var host = flag.String("h", "127.0.0.1", "host")

func main() {
	flag.Parse()

	err := s2s.InitS2s("ramcloud_test", "d10e28e567224404e7d93f35067fa54b6c4e8def35ed93cd33e4df6db5d3c5a0")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 2; i++ {
		ptr, err := s2s.NewS2sClient(111)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := ptr.Register([]string{*host}, 2222+i); err != nil {
			fmt.Println(err)
			return
		}
		/*
			if err := ptr.UnRegister(); err != nil {
				fmt.Println(err)
				return
			}
		*/

	}
	var w sync.WaitGroup
	w.Add(1)
	w.Wait()
}
