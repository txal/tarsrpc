package util

import (
	"strconv"
	"strings"
	"time"
)

const layouttime string = "20060102"            //go语言时间的格式化字符串
const TimeLayout string = "2006-01-02 15:04:05" //go语言时间的格式化字符串
const TimeLayoutOthers string = "20060102 15:04:05"
const TimeLayoutOther string = "20060102"
const TimeHtml5Layout string = "2006-01-02T15:04" //html5 time input format.

const TimeLayoutFormat2 string = "2006-01-02"

const TimeFormatYYYYMMDDHH string = "2006010215" //go语言时间的格式化字符串

func TimeStrToUnix(timeStr string) (tm int64) {
	loc, _ := time.LoadLocation("Local")
	tm2, _ := time.ParseInLocation(TimeHtml5Layout, timeStr, loc)
	return tm2.Unix()
}

func TimeStrToUnixOther(timeStr string) (tm int64) {
	temptime, _ := time.ParseInLocation(TimeLayoutOthers, timeStr, time.Local)
	tm = temptime.Unix()
	return tm
}

//根据今天的 时间戳 + interval（秒：负值就是减） 得到对应日期的字符串
func GetTodayStringByte(interval int64) string {
	timeStamp := time.Now().Unix() + interval
	return time.Unix(timeStamp, 0).Format(layouttime)
}

//获取离当前偏移interval秒的当天第一秒(0点0分0秒)
func GetTodayFirstSecond(interval int64) (tm int64) {
	timeStamp := time.Now().Unix() + interval
	timeStrs := strings.Fields(time.Unix(timeStamp, 0).Format(TimeLayout))
	timeStrs[1] = "00:00:00"
	loc, _ := time.LoadLocation("Local")
	tm2, _ := time.ParseInLocation(TimeLayout, strings.Join(timeStrs, " "), loc)
	return tm2.Unix()
}

//获取当天时间字符串(形如:2016-04-08)
func GetTodayDateInfo() string {
	timeStamp := time.Now().Unix()
	timeStrs := strings.Fields(time.Unix(timeStamp, 0).Format(TimeLayout))
	return timeStrs[0]
}

func GetGiveTimeStampDateInfo(secs int64) string {
	timeStrs := strings.Fields(time.Unix(secs, 0).Format(TimeLayout))
	return timeStrs[0]
}

func GetTodayDateInfoOther() string {
	timeStamp := time.Now().Unix()
	timeStrs := strings.Fields(time.Unix(timeStamp, 0).Format(TimeLayoutOther))
	return timeStrs[0]
}

func GetTheDayFirstSecond(timeStamp int64) (tm int64) {
	timeStrs := strings.Fields(time.Unix(timeStamp, 0).Format(TimeLayout))
	timeStrs[1] = "00:00:00"
	loc, _ := time.LoadLocation("Local")
	tm2, _ := time.ParseInLocation(TimeLayout, strings.Join(timeStrs, " "), loc)
	return tm2.Unix()
}

func GetCurUTCtimestamp() int64 {
	return time.Now().Unix()
}

func GetCurUTCNanotimestamp() int64 {
	return time.Now().UnixNano()
}

func GetCurCSTdate() string {
	return time.Now().Format("20060102")
}

func GetMidnightUTCtimestamp(curtimestamp int64, delay int64) int64 {
	ts := int64((curtimestamp+28800)/86400)*86400 + delay*86400 - 28800
	return ts
}

func Int64ToStringSlice(in []int64) (r []string) {
	for _, v := range in {
		r = append(r, strconv.FormatInt(v, 10))
	}
	return r
}

func StringToInt64Slice(in []string) (r []int64) {
	for _, v := range in {
		i, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			r = append(r, i)
		}
	}
	return r
}

//获取之前月份的日期 如 201608 201607 依次类推，参数 n表示之前几个月
func StringBeforDateByN(n int) (r []string) {
	for i := 0; i < n; i++ {
		now := time.Now()
		year := now.Year()
		mount := int(now.Month())
		if (mount - i) <= 0 {
			temp := (i-mount)/12 + 1
			year -= temp
			mount = 12 - (i-mount)%12
		} else {
			mount -= i
		}
		str := strconv.Itoa(year)
		if mount < 10 {
			str += ("0" + strconv.Itoa(mount))
		} else {
			str += strconv.Itoa(mount)
		}
		r = append(r, str)
	}
	return r
}

//return: yyyy-mm-dd, hh, todaySecs
func GetTodayTimeFormatInfo(secs int64) (dateInfo string, timeHour int, todayPassSecs int) {
	timeStrs := strings.Fields(time.Unix(secs, 0).Format(TimeLayout))
	timeVec := strings.Split(timeStrs[1], ":")

	if len(timeVec) >= 3 {
		timeHour, _ = strconv.Atoi(timeVec[0])
		timeMins, _ := strconv.Atoi(timeVec[1])
		todayPassSecs = timeHour*3600 + timeMins*60
	}

	return timeStrs[0], timeHour, todayPassSecs
}

func GetGiveTimesFirstSecs(secs int64) (firstSecs int64) {
	timeStrs := strings.Fields(time.Unix(secs, 0).Format(TimeLayout))
	timeStrs[1] = "00:00:00"

	loc, _ := time.LoadLocation("Local")
	tm2, _ := time.ParseInLocation(TimeLayout, strings.Join(timeStrs, " "), loc)
	firstSecs = tm2.Unix()

	return firstSecs
}

//yyyy-mm-dd
func GetTodayTimeYMD2() string {
	cur_tick_local, _ := time.ParseInLocation(TimeLayout, time.Unix(time.Now().Unix(), 0).Format(TimeLayout), time.Local)
	return time.Unix(cur_tick_local.Unix(), 0).Format(TimeLayoutFormat2)
}

// int64(utc) -> string(yyyy-mm-dd)
func GetUtcTimeYMD2(utc int64) string {
	return time.Unix(utc, 0).Format(TimeLayoutFormat2)
}

/**
  下面函数从分散在各个进程中的util来,先加进来,后续最好能把相同的函数合并
*/
func GetGiveTimeTimeStamp(str_time string) int64 {
	cur_tick_local, _ := time.ParseInLocation(TimeLayout, str_time, time.Local)
	return cur_tick_local.Unix()
}

//2006-01-02 15:04:05
func GetGiveTimeFormatTime(give_time int64) string {
	cur_tick_local, _ := time.ParseInLocation(TimeLayout, time.Unix(give_time, 0).Format(TimeLayout), time.Local)
	return time.Unix(cur_tick_local.Unix(), 0).Format(TimeLayout)
}

// time a bigger tan b YYYY-MM-DD HH:MI:SS
func IsTimeaGTETimeb(time_a string, time_b string) (is_gt bool) {
	tick_a, _ := time.ParseInLocation(TimeLayout, time_a, time.Local)
	tick_b, _ := time.ParseInLocation(TimeLayout, time_b, time.Local)
	//log.Info("time_a:%v,time_b:%v,tick_a:%v,tick_b:%v",time_a,time_b,tick_a.Unix(),tick_b.Unix())
	is_gt = tick_a.Unix() >= tick_b.Unix()
	return is_gt
}

//2006-01-02 15:04:05
func GetCurFormatTime() string {
	cur_tick_local, _ := time.ParseInLocation(TimeLayout, time.Unix(time.Now().Unix(), 0).Format(TimeLayout), time.Local)
	return time.Unix(cur_tick_local.Unix(), 0).Format(TimeLayout)
}

//20060102
func GetGiveTimeStampYMD(timestamp int64) string {
	cur_tick_local, _ := time.ParseInLocation(TimeLayout, time.Unix(timestamp, 0).Format(TimeLayout), time.Local)
	return time.Unix(cur_tick_local.Unix(), 0).Format(layouttime)
}

// yyyymmdd
func GetCurTimeYMD() string {
	cur_tick_local, _ := time.ParseInLocation(TimeLayout, time.Unix(time.Now().Unix(), 0).Format(TimeLayout), time.Local)
	return time.Unix(cur_tick_local.Unix(), 0).Format(layouttime)
}

// yyyymmddhh
func GetCurTimeYYYYMMDDHH() string {
	cur_tick_local, _ := time.ParseInLocation(TimeLayout, time.Unix(time.Now().Unix(), 0).Format(TimeLayout), time.Local)
	return time.Unix(cur_tick_local.Unix(), 0).Format(TimeFormatYYYYMMDDHH)
}

//yyyy-mm-dd hh:mm:ss -->yyyymmddhh
func GetGiveTimeYYYYMMDDHH(str_time string) string {
	cur_tick_local, _ := time.ParseInLocation(TimeLayout, str_time, time.Local)
	return time.Unix(cur_tick_local.Unix(), 0).Format(TimeFormatYYYYMMDDHH)
}

func TimestampToYYYYMMDDHH(timestamp int64) string {
	cur_tick_local, _ := time.ParseInLocation(TimeFormatYYYYMMDDHH, time.Unix(timestamp, 0).Format(TimeFormatYYYYMMDDHH), time.Local)
	return time.Unix(cur_tick_local.Unix(), 0).Format(TimeFormatYYYYMMDDHH)
}

func GetTomorrowFirstSecond(secs int64) (last_stamp int64) {
	timeStrs := strings.Fields(time.Unix(secs+86400, 0).Format(TimeLayout))
	timeStrs[1] = "00:00:00"

	loc, _ := time.LoadLocation("Local")
	tm2, _ := time.ParseInLocation(TimeLayout, strings.Join(timeStrs, " "), loc)
	last_stamp = tm2.Unix()
	return
}

func GetNextHourFirstSecond(secs int64) (next_hour_stamp int64) {
	timeStrs := strings.Fields(time.Unix(secs+3600, 0).Format(TimeLayout))
	//timeStrs[1] = "00:00:00"
	timeinfos := strings.Split(timeStrs[1], ":")
	if len(timeinfos) >= 3 {
		timeStrs[1] = timeinfos[0] + ":00:00"
	}

	loc, _ := time.LoadLocation("Local")
	tm2, _ := time.ParseInLocation(TimeLayout, strings.Join(timeStrs, " "), loc)
	next_hour_stamp = tm2.Unix()
	return
}
