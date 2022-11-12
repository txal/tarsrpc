package stat

import (
	"fmt"
	"testing"
	"time"

	"tarsrpc/kissgo/appzaplog/zap"

	"tarsrpc/kissgo/appzaplog"
)

const (
	url   = "http://106.55.188.90:8086"
	token = "yP_pgRWy72pQxtlRfwvu6CMs8BzTMScL_pZINmOT3Hc0UvxFTOOsoGYXevpjvmIR75DPgA8jC23OFZfFpqWnlw=="
	org   = "dianyun"
)

func TestNewStat(t *testing.T) {
	appzaplog.InitAppLog()
	appzaplog.SetLogLevel("debug")

	appzaplog.Info("info", zap.String("a", "a=="))

	Init(url, token, org)

	report(time.Now().Add(-6*time.Minute), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, false, false)
	report(time.Now().Add(-6*time.Minute), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, false, false)
	report(time.Now().Add(-6*time.Minute), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, false, false)

	report(time.Now().Add(-6*time.Minute), Tag{"kaiheiyun", "user", "UserExtObj", "CreateUser", "127.0.0.1", "127.0.0.2"}, 10, false, false)

	fmt.Println(statInstance)

	// 1分钟后
	report(time.Now().Add(-5*time.Minute), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, false, false)
	report(time.Now().Add(-5*time.Minute), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, false, false)

	// 5分钟后
	report(time.Now(), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, false, false)
	select {}
}

func display(s *stat) {
	front := s.list.Front()

	for front != nil {
		node := front.Value.(*statNode)
		fmt.Print(node.Timestamp)
		for k, v := range node.Data {
			fmt.Print(k, *v)
		}
		fmt.Println()
		front = front.Next()
	}
}

func Test_stat_Report(t *testing.T) {
	statInstance = &stat{}

	report(time.Now().Add(-time.Minute), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, false, false)
	report(time.Now().Add(-time.Minute), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, false, false)
	report(time.Now().Add(-time.Minute), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, false, true)

	report(time.Now().Add(-time.Minute), Tag{"kaiheiyun", "user", "UserExtObj", "CreateUser", "127.0.0.1", "127.0.0.2"}, 10, false, false)

	// 一分钟后
	report(time.Now(), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, true, false)
	report(time.Now(), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, false, true)

	display(statInstance)
	if statInstance.list.Len() != 2 {
		t.Fatal()
	}
	if len(statInstance.list.Front().Value.(*statNode).Data) != 1 {
		t.Fatal()
	}
	if len(statInstance.list.Front().Next().Value.(*statNode).Data) != 2 {
		t.Fatal()
	}
}

/*
*
goos: darwin
goarch: amd64
pkg: tarsrpc/tars/servant/stat
cpu: Intel(R) Core(TM) i5-8500B CPU @ 3.00GHz
BenchmarkName
BenchmarkName-6   	 8449570	       135.5 ns/op
PASS

7407407 ops/s
*/
func BenchmarkName(b *testing.B) {
	Init(url, token, org)
	for i := 0; i < b.N; i++ {
		report(time.Now(), Tag{"kaiheiyun", "user", "UserExtObj", "GetUser", "127.0.0.1", "127.0.0.2"}, 10, false, false)
	}
}
