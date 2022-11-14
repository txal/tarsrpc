package util_test

import (
	"fmt"
	"testing"
	"time"
	"code.com/tars/goframework/kissgo/util"
)

func sleepMsec(key string, n int) int {
	fmt.Println("sleepMsec", key, n, "begin")
	time.Sleep(time.Millisecond * time.Duration(n))
	fmt.Println("sleepMsec", key, n, "end")
	return n * 10
}

func TestBatchWork(t *testing.T) {
	n1 := 0
	f1 := func() error {
		n1 = sleepMsec("f1", 100)
		return nil
	}

	n2 := 0
	f2 := func() error {
		n2 = sleepMsec("f2", 200)
		//return fmt.Errorf("f2 error")
		return nil
	}

	n3 := 0
	f3 := func() error {
		n3 = sleepMsec("f3", 300)
		return nil
	}

	err := util.BatchWork(f1, f2, f3)
	if err != nil {
		t.Fatal(err)
	}

	if n1 != 1000 || n2 != 2000 || n3 != 3000 {
		t.Fatal(n1, n2, n3)
	}
	fmt.Println(n1, n2, n3)
}

func TestBatchWorkErr(t *testing.T) {
	n1 := 0
	f1 := func() error {
		n1 = sleepMsec("f1", 100)
		return fmt.Errorf("f1 error")
	}

	n2 := 0
	f2 := func() error {
		n2 = sleepMsec("f2", 200)
		return nil
	}

	err := util.BatchWork(f1, f2)
	if err.Error() != "f1 error" {
		t.Fatal(err)
	}
}
