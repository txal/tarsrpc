package util_test

import (
	"testing"

	"code.com/tars/goframework/kissgo/util"
)

var utc int64 = 1486465081                 //时间戳
var fullstr string = "2017-02-07 18:58:01" //全量字符串
var datestr string = "2017-02-07"          //日期字符串
var timestr string = "18:58:01"            //时间字符串
var ymd string = "20170207"                //年月日
var ymdh string = "2017020718"             //年月日时
var ymdhms string = "20170207185801"       //年月日时分秒

var pp bool = true

func TestTimeUtc(t *testing.T) {
	if util.TimeUtcToFullstr(utc) != fullstr {
		t.Fatal("fail utc -> fullstr")
	}
	if util.TimeUtcToDatestr(utc) != datestr {
		t.Fatal("fail utc -> datestr")
	}
	if util.TimeUtcToTimestr(utc) != timestr {
		t.Fatal("fail utc -> timestr")
	}
	if util.TimeUtcToYmd(utc) != ymd {
		t.Fatal("fail utc -> ymd")
	}
	if util.TimeUtcToYmdh(utc) != ymdh {
		t.Fatal("fail utc -> ymdh")
	}
	if util.TimeUtcToYmdhms(utc) != ymdhms {
		t.Fatal("fail utc -> ymdhms")
	}
	if pp {
		t.Log(util.TimeUtcToFormat("2006-2006--*&^%&%--01-02", utc))
	}
}

func TestTimeNow(t *testing.T) {
	if pp {
		t.Log(util.TimeNowUtc(), util.TimeNowUtc(-100))
		t.Log(util.TimeNowYmd(), util.TimeNowYmd(-100))
		t.Log(util.TimeNowYmdh(), util.TimeNowYmdh(-100))
		t.Log(util.TimeNowYmdhms(), util.TimeNowYmdhms(-100))
		t.Log(util.TimeNowFullstr(), util.TimeNowFullstr(-100))
		t.Log(util.TimeNowDatestr(), util.TimeNowDatestr(-100))
		t.Log(util.TimeNowTimestr(), util.TimeNowTimestr(-100))
	}
	if pp {
		t.Log(util.TimeNowFormat("2006-2006 01-02"), util.TimeNowFormat("2006-2006 01-02", -86400))
	}
}

func TestTimeFullstr(t *testing.T) {
	if util.TimeFullstrToUtc(fullstr) != utc {
		t.Fatal("fail fullstr -> utc")
	}
	if util.TimeFullstrToYmd(fullstr) != ymd {
		t.Fatal("fail fullstr -> ymd")
	}
	if util.TimeFullstrToYmdh(fullstr) != ymdh {
		t.Fatal("fail fullstr -> ymdh")
	}
	if util.TimeFullstrToYmdhms(fullstr) != ymdhms {
		t.Fatal("fail fullstr -> ymdhms")
	}
	if util.TimeFullstrToDatestr(fullstr) != datestr {
		t.Fatal("fail fullstr -> datestr")
	}
	if util.TimeFullstrToTimestr(fullstr) != timestr {
		t.Fatal("fail fullstr -> timestr")
	}
	if pp {
		t.Log(util.TimeFullstrToFormat("2006-2006 01-02", fullstr))
	}
}

func TestTimeDatestr(t *testing.T) {
	targetUtc := util.TimeFullstrToUtc(datestr + " 00:00:00")

	if util.TimeDatestrToUtc(datestr) != targetUtc {
		t.Fatal("fail datestr -> utc")
	}
	if util.TimeDatestrToYmd(datestr) != ymd {
		t.Fatal("fail datestr -> ymd")
	}
	if util.TimeDatestrToYmdh(datestr) != util.TimeUtcToYmdh(targetUtc) {
		t.Fatal("fail datestr -> ymdh")
	}
	if util.TimeDatestrToFullstr(datestr) != util.TimeUtcToFullstr(targetUtc) {
		t.Fatal("fail datestr -> fullstr")
	}
}

func TestTimeYmd(t *testing.T) {
	targetUtc := util.TimeFullstrToUtc(datestr + " 00:00:00")

	if util.TimeYmdToUtc(ymd) != targetUtc {
		t.Fatal("fail ymd -> utc")
	}
	if util.TimeYmdToYmdh(ymd) != util.TimeUtcToYmdh(targetUtc) {
		t.Fatal("fail ymd -> ymdh")
	}
	if util.TimeYmdToFullstr(ymd) != util.TimeUtcToFullstr(targetUtc) {
		t.Fatal("fail ymd -> fullstr")
	}
	if util.TimeYmdToDatestr(ymd) != util.TimeUtcToDatestr(targetUtc) {
		t.Fatal("fail ymd -> datestr")
	}
}

func TestTimeYmdh(t *testing.T) {
	targetUtc := util.TimeFullstrToUtc(datestr + " " + ymdh[8:] + ":00:00") //hh:00:00

	if util.TimeYmdhToUtc(ymdh) != targetUtc {
		t.Fatal("fail ymdh -> utc")
	}
	if util.TimeYmdhToYmd(ymdh) != util.TimeUtcToYmd(targetUtc) {
		t.Fatal("fail ymdh -> ymd")
	}
	if util.TimeYmdhToFullstr(ymdh) != util.TimeUtcToFullstr(targetUtc) {
		t.Fatal("fail ymdh -> fullstr")
	}
	if util.TimeYmdhToDatestr(ymdh) != util.TimeUtcToDatestr(targetUtc) {
		t.Fatal("fail ymdh -> datestr")
	}
}

func TestTimeFirst(t *testing.T) {
	if pp {
		t.Log("TimeFirstMinute", util.TimeFirstMinute(utc), util.TimeUtcToFullstr(util.TimeFirstMinute(utc)))
		t.Log("TimeFirstHour", util.TimeFirstHour(utc), util.TimeUtcToFullstr(util.TimeFirstHour(utc)))
		t.Log("TimeFirstDate", util.TimeFirstDate(utc), util.TimeUtcToFullstr(util.TimeFirstDate(utc)))
		t.Log("TimeFirstWeek", util.TimeFirstWeek(utc), util.TimeUtcToFullstr(util.TimeFirstWeek(utc)))
		t.Log("TimeFirstMonth", util.TimeFirstMonth(utc), util.TimeUtcToFullstr(util.TimeFirstMonth(utc)))
		t.Log("TimeFirstYear", util.TimeFirstYear(utc), util.TimeUtcToFullstr(util.TimeFirstYear(utc)))
	}
	if pp {
		if util.TimeCmpSameDate(utc, utc-86400) != false {
			t.Fatal("fail TimeCmpSameDate")
		}
		if util.TimeCmpSameDate(utc, utc+1) != true {
			t.Fatal("fail TimeCmpSameDate")
		}
		//util.TimeCmpSameWeek
		//util.TimeCmpSameMonth
		//util.TimeCmpSameYear
	}
}
