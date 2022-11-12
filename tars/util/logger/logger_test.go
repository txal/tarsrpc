// @author kordenlu
// @创建时间 2018/01/26 14:43
// 功能描述:

package logger

import "testing"

func TestString2Level(t *testing.T) {
	level := "DEBUG"
	if intlevel,err := String2Level(level);err != nil{
		t.Error("String2Level err",err)
	}else if intlevel != DEBUG{
		t.Error("should be debug")
	}
}
