package appzaplog

import (
	"tarsrpc/kissgo/appzaplog/zap"
	"tarsrpc/kissgo/appzaplog/zap/zapcore"
)

func UID(uid uint64) zapcore.Field {
	return zap.Uint64("uid", uid)
}

func ROOMID(roomid uint64) zapcore.Field {
	return zap.Uint64("roomid", roomid)
}
