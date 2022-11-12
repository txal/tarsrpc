package endpoint

import (
	"net"
	"strconv"
	"tarsrpc/jce/servant/taf"
)

func Taf2endpoint(end taf.EndpointF) Endpoint {
	proto := "tcp"
	if end.Istcp == 0 {
		proto = "udp"
	}

	return Endpoint{
		Host:      end.Host,
		Port:      int32(end.Port),
		IPPort:    net.JoinHostPort(end.Host, strconv.FormatInt(int64(end.Port), 10)),
		Timeout:   int32(end.Timeout),
		Proto:     proto,
		Container: end.ContainerName,
	}

}

func Endpoint2taf(end Endpoint) taf.EndpointF {
	return taf.EndpointF{
		Host:          end.Host,
		Port:          int32(end.Port),
		Timeout:       int32(end.Timeout),
		Istcp:         end.istcp(),
		ContainerName: end.Container,
	}
}
