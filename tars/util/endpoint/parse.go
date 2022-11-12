package endpoint

import (
	"flag"
	"strings"
	"strconv"
	"net"
)

func Parse(endpoint string) Endpoint {
	//tcp -h 10.219.139.142 -p 19386 -t 60000
	proto := endpoint[0:3]
	pFlag := flag.NewFlagSet(proto, flag.ContinueOnError)
	var host, bind string
	var port, timeout int
	pFlag.StringVar(&host, "h", "", "host")
	pFlag.IntVar(&port, "p", 0, "port")
	pFlag.IntVar(&timeout, "t", 3000, "timeout")
	pFlag.StringVar(&bind, "b", "", "bind")
	pFlag.Parse(strings.Fields(endpoint)[1:])

	return Endpoint{
		Host:    host,
		Port:    int32(port),
		IPPort:  net.JoinHostPort(host,strconv.FormatInt(int64(port),10)),
		Timeout: int32(timeout),
		Proto:   proto,
		Bind:    bind,
	}
}
