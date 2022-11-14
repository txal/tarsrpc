package util_test

import (
	"fmt"
	"testing"
	"time"
	"code.com/tars/goframework/kissgo/util"
)

func TestRecordQuick(t *testing.T) {
	f := func() {
		all := 0
		for i := 0; i < 10000*500; i++ {
			all += i*10 + 1
		}
	}
	util.RecordQuick("test1", f)
	util.RecordQuick("test2", f)

	u3 := util.RecordQuick("", f)
	fmt.Printf("RecordQuick use %d ms\n", u3)

	return
}

func f100ms(n int64) {
	time.Sleep(time.Millisecond * 100 * time.Duration(n))
}

func TestRecordStart(t *testing.T) {

	rt := util.RecordStart()
	f100ms(1)
	rt.Mark("name_100")
	f100ms(2)
	rt.Mark("name_200")
	f100ms(3)
	rt.Mark("name_300")
	rt.Print("mytest")

	if true {
		if rt.TotalMsec() != 600 {
			t.Fatal("rt.TotalMsec()!=600", rt.TotalMsec())
		}
		if rt.TotalCount() != 3 {
			t.Fatal("rt.TotalMsec()!=3", rt.TotalCount())
		}
		rt.Reset()
		if rt.TotalCount() != 0 {
			t.Fatal("Reset() after, rt.TotalMsec()!=0", rt.TotalCount())
		}
	}

	return
}
