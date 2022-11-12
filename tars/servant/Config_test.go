// @author kordenlu
// @创建时间 2018/02/09 11:00
// 功能描述:

package servant

import (
	"testing"
	"time"
)

func TestTakeclientsynctimeinvoketimeout(t *testing.T) {
	duration := takeclientsynctimeinvoketimeout()
	if duration != 3*time.Second{
		t.Error("default should be 3s")
	}

	cltCfg = &clientConfig{
		syncInvokeTimeout:4000,
	}

	if duration = takeclientsynctimeinvoketimeout();duration != 3*time.Second{
		t.Error("should be:",cltCfg.syncInvokeTimeout)
	}
}