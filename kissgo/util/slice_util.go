// @author kordenlu
// @创建时间 2017/01/24 15:00
// 功能描述:

package util

import "reflect"

func Contain(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

func ContainIndex(s interface{}, elem interface{}) (bool, int64) {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == elem {
				return true, int64(i)
			}
		}
	}
	return false, 0
}
