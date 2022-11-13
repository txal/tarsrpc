package util_test

import (
	"code.com/tars/goframework/kissgo/util"
	"net/http"
	"testing"
)

func TestHttpApiName(t *testing.T) {
	apiname := "testapiname"
	req, err := http.NewRequest("GET", "http://127.0.0.1:100/v1/"+apiname, nil)
	if err != nil {
		t.Errorf("NewRequest err:%v", err)
	}
	if httpapiname := util.HttpApiName(req); httpapiname != apiname {
		t.Errorf("HttpApiName not equal,httpapiname:%v,apiname:%v", httpapiname, apiname)
	}
}
