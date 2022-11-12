/***************************************************
	时间转换 by huangzhibin
	------------------------------------------------
	1、只提供常用转换，特殊转换需要自定义format，不过要先转换为utc或fullstr
	2、format格式： go/time/format.go:stdLongMonth
	3、函数名字前缀为Time，为了方便编辑器自动补全
***************************************************/
package util

import (
	"strings"
	"time"
)

const (
	formatUtc     = 1486465081            //时间戳，1970年1月1日到现在的秒数
	formatYmd     = "20060102"            //年月日
	formatYmdh    = "2006010215"          //年月日时
	formatYmdhms  = "20060102150405"      //年月日时分秒，只实现utc转换
	formatFullstr = "2006-01-02 15:04:05" //时间日期字符串
	formatDatestr = "2006-01-02"          //日期字符串
	formatTimestr = "15:04:05"            //时间字符串
)

//---------------------- utc -------------------------
//utc --> ymd
func TimeUtcToYmd(utc int64) string {
	return time.Unix(utc, 0).Format(formatYmd)
}

//utc --> ymdh
func TimeUtcToYmdh(utc int64) string {
	return time.Unix(utc, 0).Format(formatYmdh)
}

//utc -> ymdhms
func TimeUtcToYmdhms(utc int64) string {
	return time.Unix(utc, 0).Format(formatYmdhms)
}

//utc --> fullstr
func TimeUtcToFullstr(utc int64) string {
	return time.Unix(utc, 0).Format(formatFullstr)
}

//utc -> datestr
func TimeUtcToDatestr(utc int64) string {
	return time.Unix(utc, 0).Format(formatDatestr)
}

//utc -> timestr
func TimeUtcToTimestr(utc int64) string {
	return time.Unix(utc, 0).Format(formatTimestr)
}

//utc -> format
func TimeUtcToFormat(format string, utc int64) string {
	return time.Unix(utc, 0).Format(format)
}

//--------------------- now/offset --------------------------
//now --> utc
func TimeNowUtc(offset ...int64) int64 {
	utc := time.Now().Unix()
	if len(offset) > 0 {
		utc += offset[0]
	}
	return utc
}

//now --> ymd
func TimeNowYmd(offset ...int64) string {
	utc := time.Now().Unix()
	if len(offset) > 0 {
		utc += offset[0]
	}
	return TimeUtcToYmd(utc)
}

//now --> ymdh
func TimeNowYmdh(offset ...int64) string {
	utc := time.Now().Unix()
	if len(offset) > 0 {
		utc += offset[0]
	}
	return TimeUtcToYmdh(utc)
}

//now --> ymdhms
func TimeNowYmdhms(offset ...int64) string {
	utc := time.Now().Unix()
	if len(offset) > 0 {
		utc += offset[0]
	}
	return TimeUtcToYmdhms(utc)
}

//now --> fullstr
func TimeNowFullstr(offset ...int64) string {
	utc := time.Now().Unix()
	if len(offset) > 0 {
		utc += offset[0]
	}
	return TimeUtcToFullstr(utc)
}

//now -> datestr
func TimeNowDatestr(offset ...int64) string {
	utc := time.Now().Unix()
	if len(offset) > 0 {
		utc += offset[0]
	}
	return TimeUtcToDatestr(utc)
}

//now -> timestr
func TimeNowTimestr(offset ...int64) string {
	utc := time.Now().Unix()
	if len(offset) > 0 {
		utc += offset[0]
	}
	return TimeUtcToTimestr(utc)
}

//now -> format
func TimeNowFormat(format string, offset ...int64) string {
	utc := time.Now().Unix()
	if len(offset) > 0 {
		utc += offset[0]
	}
	return TimeUtcToFormat(format, utc)
}

//---------------------- fullstr -------------------------
//fullstr -> utc
func TimeFullstrToUtc(str string) int64 {
	tm, _ := time.ParseInLocation(formatFullstr, str, time.Local)
	return tm.Unix()
}

//fullstr -> ymd
func TimeFullstrToYmd(str string) string {
	tm, _ := time.ParseInLocation(formatFullstr, str, time.Local)
	return time.Unix(tm.Unix(), 0).Format(formatYmd)
}

//fullstr --> ymdh
func TimeFullstrToYmdh(str string) string {
	tm, _ := time.ParseInLocation(formatFullstr, str, time.Local)
	return time.Unix(tm.Unix(), 0).Format(formatYmdh)
}

//fullstr --> ymdhms
func TimeFullstrToYmdhms(str string) string {
	tm, _ := time.ParseInLocation(formatFullstr, str, time.Local)
	return time.Unix(tm.Unix(), 0).Format(formatYmdhms)
}

//fullstr --> datestr
func TimeFullstrToDatestr(str string) string {
	d := strings.Split(str, " ")
	if len(d) >= 1 {
		return d[0]
	}
	return ""
}

//fullstr --> timestr
func TimeFullstrToTimestr(str string) string {
	d := strings.Split(str, " ")
	if len(d) >= 2 {
		return d[1]
	}
	return ""
}

//fullstr --> format
func TimeFullstrToFormat(format string, str string) string {
	tm, _ := time.ParseInLocation(formatFullstr, str, time.Local)
	return time.Unix(tm.Unix(), 0).Format(format)
}

//---------------------- datestr -------------------------
//datestr -> utc
func TimeDatestrToUtc(str string) int64 {
	tm, _ := time.ParseInLocation(formatDatestr, str, time.Local)
	return tm.Unix()
}

//datestr -> ymd
func TimeDatestrToYmd(str string) string {
	tm, _ := time.ParseInLocation(formatDatestr, str, time.Local)
	return time.Unix(tm.Unix(), 0).Format(formatYmd)
}

//datestr --> ymdh
func TimeDatestrToYmdh(str string) string {
	tm, _ := time.ParseInLocation(formatDatestr, str, time.Local)
	return time.Unix(tm.Unix(), 0).Format(formatYmdh)
}

//datestr --> fullstr
func TimeDatestrToFullstr(str_time string) string {
	return str_time + " 00:00:00"
}

//datestr --> timestr
//不存在转换

//---------------------- timestr -------------------------
//不存在转换

//---------------------- ymd -------------------------
//ymd -> utc
func TimeYmdToUtc(str string) int64 {
	tm, _ := time.ParseInLocation(formatYmd, str, time.Local)
	return tm.Unix()
}

//ymd -> ymdh
func TimeYmdToYmdh(str string) string {
	return str + "00"
}

//ymd -> fullstr
func TimeYmdToFullstr(str string) string {
	return TimeUtcToFullstr(TimeYmdToUtc(str)) //todo待优化
}

//ymd -> datestr
func TimeYmdToDatestr(str string) string {
	return TimeUtcToDatestr(TimeYmdToUtc(str)) //todo待优化
}

//ymd -> timestr
//不存在转换

//---------------------- ymdh -------------------------
//ymdh -> utc
func TimeYmdhToUtc(str string) int64 {
	tm, _ := time.ParseInLocation(formatYmdh, str, time.Local)
	return tm.Unix()
}

//ymdh -> ymd
func TimeYmdhToYmd(str string) string {
	if len(str) >= 10 {
		return str[:len(str)-2]
	}
	return str
}

//ymdh -> fullstr
func TimeYmdhToFullstr(str string) string {
	return TimeUtcToFullstr(TimeYmdhToUtc(str)) //todo待优化
}

//ymdh -> datestr
func TimeYmdhToDatestr(str string) string {
	return TimeUtcToDatestr(TimeYmdhToUtc(str)) //todo待优化
}

//ymdh -> timestr
//不存在转换

//---------------------- utc relate -------------------------
//utc的当前分钟的0秒
func TimeFirstMinute(utc int64) int64 {
	return utc - utc%60
}

//utc的当前小时的0分0秒
func TimeFirstHour(utc int64) int64 {
	return utc - utc%3600
}

func timeUtcToZeroTm(utc int64) time.Time {
	str_time := time.Unix(utc, 0).Format(formatDatestr) + " 00:00:00"
	//	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation(formatFullstr, str_time, time.Local)
	return tm
}

//utc的当前日期的0点
func TimeFirstDate(utc int64) int64 {
	tm := timeUtcToZeroTm(utc)
	return tm.Unix()
}

//utc的当前星期的周日0点（周日是第一天）
func TimeFirstWeek(utc int64) int64 {
	tm := timeUtcToZeroTm(utc)
	return tm.Unix() - int64(tm.Weekday())*86400
}

//utc的当前月份的1号0点
func TimeFirstMonth(utc int64) int64 {
	tm := timeUtcToZeroTm(utc)
	return tm.Unix() - int64(tm.Day()-1)*86400
}

//utc的当前年份的1月1日0点
func TimeFirstYear(utc int64) int64 {
	tm := timeUtcToZeroTm(utc)
	return tm.Unix() - int64(tm.YearDay()-1)*86400
}

//utc是否同一天
func TimeCmpSameDate(utc1 int64, utc2 int64) bool {
	return TimeFirstDate(utc1) == TimeFirstDate(utc2)
}

//utc是否同一周
func TimeCmpSameWeek(utc1 int64, utc2 int64) bool {
	return TimeFirstWeek(utc1) == TimeFirstWeek(utc2)
}

//utc是否同一月
func TimeCmpSameMonth(utc1 int64, utc2 int64) bool {
	return TimeFirstMonth(utc1) == TimeFirstMonth(utc2)
}

//utc是否同一年
func TimeCmpSameYear(utc1 int64, utc2 int64) bool {
	return TimeFirstYear(utc1) == TimeFirstYear(utc2)
}
