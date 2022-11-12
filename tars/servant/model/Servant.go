package model

import (
	"context"
	"tarsrpc/jce/taf"
	pbtaf "tarsrpc/tars/servant/protocol"
)

type Servant interface {
	Taf_invoke(
		ctx context.Context,
		ctype byte,
		sFuncName string,
		buf []byte) (*taf.ResponsePacket, error)
	Proxy_invoke(ctx context.Context, ctype byte, sFuncName string,
		buf []byte, ipPort string) (*taf.ResponsePacket, error)
	GetProxyEndPoints() []string
}

type PbServant interface {
	Pb_invoke(
		ctx context.Context,
		ctype byte,
		sFuncName string,
		buf []byte,
		status map[string]string,
		context map[string]string) (*pbtaf.ResponsePacket, error)
}
