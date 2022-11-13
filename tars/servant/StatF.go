package servant

import (
	"strings"

	"code.com/tars/goframework/kissgo/appzaplog"
	"code.com/tars/goframework/kissgo/appzaplog/zap"

	"code.com/tars/goframework/kissgo/gobreaker"

	"code.com/tars/goframework/jce/servant/taf"

	"code.com/tars/goframework/tars/servant/stat"
)

func parseObjName(obj string) (string, string, string) {
	strs := strings.Split(obj, ".")
	if len(strs) != 3 {
		return "", "", ""
	}
	return strs[0], strs[1], strs[2]
}

// ReportStat 接口监控上报入口,这里就当做适配器吧
func ReportStat(msg IMessage) {
	var (
		isError   = false
		isTimeout = false
	)

	if msg.getRespRet() == taf.JCEINVOKETIMEOUT {
		isTimeout = true
	} else if msg.getRespRet() != 0 && msg.getRespRet() < gobreaker.Failure {
		isError = true
	}

	namespace, server, object := parseObjName(msg.getSServantName())
	cfg := GetServerConfig()
	serverIP := ""
	if msg.getAdapterProxy() != nil && msg.getAdapterProxy().GetPoint() != nil {
		serverIP = msg.getAdapterProxy().GetPoint().Host
	}
	tag := stat.Tag{
		Namespace: namespace,
		Server:    server,
		Object:    object,
		Function:  msg.getFuncName(),
		ClientIP:  cfg.LocalIP,
		ServerIP:  serverIP,
	}
	stat.Report(tag, msg.Cost(), isError, isTimeout)
	appzaplog.Debug("ReportStat", zap.Any("tag", tag), zap.Int64("cost", msg.Cost()),
		zap.Bool("is_error", isError), zap.Bool("is_timeout", isTimeout))
}

type httpStatInfo struct {
	reqAddr    string
	pattern    string
	statusCode int
	costTime   int64
}

func reportHttpStat(st *httpStatInfo) {

}
