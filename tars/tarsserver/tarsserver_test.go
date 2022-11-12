// @author kordenlu
// @创建时间 2018/03/07 14:18
// 功能描述:

package tarsserver

import (
	"testing"
	"time"
)

func TestIsZombie(t *testing.T)  {
	server := NewTarsServer(nil,&TarsServerConf{
		MaxAccept:100,
		MaxInvoke:200,
	})
	time.Sleep(5*time.Second)
	if server.IsZombie(2*time.Second){
		t.Error("should't zombie")
	}
}
