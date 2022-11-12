// @author kordenlu
// @创建时间 2017/01/24 15:05
// 功能描述:

package util

import "testing"

func TestContain(t *testing.T) {
	// string
	testkey := "test"
	strslice := []string{testkey, "ok", "no"}
	exist := Contain(strslice, testkey)
	if !exist {
		t.Error("Contain failed")
	}

	exist = Contain(strslice, "tes")
	if exist {
		t.Error("Contain failed")
	}
}
