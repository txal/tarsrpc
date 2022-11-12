package util_test

import (
	"fmt"
	"tarsrpc/kissgo/util"
	"testing"
)

func TestIps(t *testing.T) {
	fmt.Printf("TestMisc\n")

	//	ips2222 := util.GetSelfIps222()
	//	fmt.Printf("GetSelfIps ips2222=%v\n", ips2222)
	//	return

	ips, err := util.GetSelfIps()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("GetSelfIps ips=%v\n", ips)

	if len(ips) == 0 {
		t.Fatal("ips is null!!")
	}

	if len(ips) > 0 {
		ip := ips[0]
		b, err := util.IsSelfIp(ip)
		if err != nil {
			t.Fatal(err)
		} else if b != true {
			t.Fatalf("IsSelfIp error ip=%v!!!", ip)
		}
		fmt.Printf("check ok ip=%v ret=%v\n", ip, b)
	}
	if len(ips) > 0 {
		ip := "127.0.0.99"
		b, err := util.IsSelfIp(ip)
		if err != nil {
			t.Fatal(err)
		} else if b != false {
			t.Fatalf("IsSelfIp error ip=%v!!!", ip)
		}
		fmt.Printf("check ok ip=%v ret=%v\n", ip, b)
	}
}

func TestIpGetOneByEthName(t *testing.T) {
	ip := util.GetDWCNCIp()
	if len(ip) == 0 {
		t.Log("ip not exist for cnc ip")
		return
	}
	t.Logf("ip:%v", ip)
}
