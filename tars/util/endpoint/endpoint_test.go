// @author kordenlu
// @创建时间 2018/01/24 14:11
// 功能描述:

package endpoint

import "testing"

func TestEndpoint2taf(t *testing.T) {
	ed := Endpoint{
		Proto: "tcp",
	}
	if ed.istcp() != int32(1) {
		t.Error("should be tcp:", ed.istcp())
	}

	ed.Proto = "udp"
	if ed.istcp() != 0 {
		t.Error("should't be tcp", ed.istcp())
	}
}
