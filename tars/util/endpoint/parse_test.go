package endpoint

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	e := Parse("tcp -h 10.219.139.142 -p 19386 -t 60000")
	if e.Host != "10.219.139.142" || e.Proto != "tcp" || e.Port != 19386 || e.Timeout != 60000 || e.Bind != "" {
		t.Error("parse failed")
	}

	e2 := Parse("udp -h 10.219.139.142 -p 19386 -t 60000")
	if e2.Host != "10.219.139.142" || e2.Proto != "udp" || e2.Port != 19386 || e2.Timeout != 60000 {
		t.Error("udb parse failed")
	}

	taf := Endpoint2taf(e2)
	if taf.Istcp != 0 || taf.Host != e2.Host || taf.Port != e2.Port {
		t.Error("Endpoint2taf failed", taf)
	}

	fmt.Println(Taf2endpoint(taf))
}
