/***************************************************
	mongo辅助工具 by huangzhibin
	------------------------------------------------
	1、数据库分布式锁：主要用于多个进程部署，但某些操作只允许单个进程同时工作
		key需要统一风格以免冲突：程序名:业务描述[:其他]， 例如： fts_http_server:luckytreasures:clean
		一般用于定时器+协程的方式，起码要超过1小时才有必要用吧。
	2、数据库唯一id生成器：int64自增，1为第一个
	3、函数名字前缀为Mgo，为了方便编辑器自动补全

	todo 考虑把多程序同时使用的mongodb/redis的名字或者操作都移到 yy.com/common/def
***************************************************/
package util

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//命名规则 MGOKEY为前缀，DB或TABLE或FIELD为中缀，后缀为具体名
const (
	MGOKEY_DB_GLDB    = "gldb" //全局db
	MGOKEY_TABLE_LOCK = "lock" //全局锁
	MGOKEY_TABLE_GEN  = "gen"  //全局id生成器
)

//------------------------------- 锁 -------------------------------------
//锁结构
type MgoLockInfo struct {
	Id       string `bson:"_id"`      //key
	LockTime int64  `bson:"lockTime"` //加锁失效时间，避免进程占用锁之后挂掉之类的
	Desc     string `bson:"desc"`     //描述：name|ip|pid|hostname
}

type getMongoFn func() (*mgo.Session, error)

//数据库分布式锁-加锁
func MgoLock(fn getMongoFn, key string, second int64) (ret bool) {
	session, err := fn()
	if err != nil {
		return
	}
	defer session.Close()

	var info MgoLockInfo
	now := time.Now().Unix()
	desc := getDesc()

	client := session.DB(MGOKEY_DB_GLDB).C(MGOKEY_TABLE_LOCK)
	_, err = client.Find(bson.M{"_id": key, "$or": []bson.M{bson.M{"desc": desc}, bson.M{"lockTime": bson.M{"$lt": now}}}}).
		Apply(mgo.Change{Update: bson.M{"$set": bson.M{"lock": true, "lockTime": now + second, "desc": desc}}, ReturnNew: true, Upsert: true}, &info)
	if err == nil {
		ret = true
	}
	return
}

//数据库分布式锁-加锁（防止电脑时间被大幅度修改）
func MgoLockEx(fn getMongoFn, key string, second int64) (ret bool, err error) {
	info, err := MgoGetLock(fn, key)
	if err != nil {
		return
	}
	if info.LockTime > 0 {
		now := time.Now().Unix()
		if now >= (info.LockTime+second*10) || now <= (info.LockTime-second*10) {
			err = fmt.Errorf("LockTime exception lock=%d now=%d diff=%d", info.LockTime, now, now-info.LockTime)
			return
		}
	}
	if MgoLock(fn, key, second) != true {
		return
	}
	ret = true
	return
}

//数据库分布式锁-解锁
func MgoUnlock(fn getMongoFn, key string) (ret bool) {
	//	session, err := fn()
	//	if err != nil {
	//		return
	//	}
	//	defer session.Close()

	//	var info MgoLockInfo

	//	client := session.DB(MGOKEY_DB_GLDB).C(MGOKEY_TABLE_LOCK)
	//	_, err = client.Find(bson.M{"_id": key, "lock": true}).
	//		Apply(mgo.Change{Update: bson.M{"$set": bson.M{"lock": false}}}, &info)
	//	if err == nil {
	//		ret = true
	//	}
	return MgoLock(fn, key, -1)
}

//获取某个锁信息
func MgoGetLock(fn getMongoFn, key string) (info MgoLockInfo, err error) {
	session, err := fn()
	if err != nil {
		return
	}
	defer session.Close()

	client := session.DB(MGOKEY_DB_GLDB).C(MGOKEY_TABLE_LOCK)
	err = client.Find(bson.M{"_id": key}).One(&info)
	if err == mgo.ErrNotFound {
		err = nil
	}
	return
}

var g_lockdesc string = ""

func getDesc() string {
	if g_lockdesc == "" {
		filename := os.Args[0]
		pos := strings.LastIndex(filename, "/")
		if pos > 0 {
			filename = filename[pos+1:]
		} else {
			pos := strings.LastIndex(filename, "\\")
			if pos > 0 {
				filename = filename[pos+1:]
			}
		}
		hostname, _ := os.Hostname()
		g_lockdesc = fmt.Sprintf("%s|%s|%d|%s", filename, IpGetOne(), os.Getpid(), hostname)
	}
	return g_lockdesc
}

//------------------------------- id分配 -------------------------------------
//自增序列结构
type MgoGenInfo struct {
	Id  string `bson:"_id"` //key
	Use int64  `bson:"use"` //当前已经使用的id
}

//生成一个自增序列
func MgoIncOne(fn getMongoFn, key string) (id int64, err error) {
	return MgoIncMore(fn, key, 1)
}

//生成多个自增序列
func MgoIncMore(fn getMongoFn, key string, count int64) (id int64, err error) {
	session, err := fn()
	if err != nil {
		return
	}
	defer session.Close()

	var info MgoGenInfo
	client := session.DB(MGOKEY_DB_GLDB).C(MGOKEY_TABLE_GEN)
	_, err = client.Find(bson.M{"_id": key}).Apply(mgo.Change{Update: bson.M{"$inc": bson.M{"use": count}}, ReturnNew: true, Upsert: true}, &info)
	if err != nil {
		return
	}
	id = info.Use - count + 1
	return
}

//重置自增序列，不能乱用！！！
func MgoSetInc(fn getMongoFn, key string, value int64) (err error) {
	session, err := fn()
	if err != nil {
		return
	}
	defer session.Close()

	var info MgoGenInfo

	client := session.DB(MGOKEY_DB_GLDB).C(MGOKEY_TABLE_GEN)
	_, err = client.Find(bson.M{"_id": key}).Apply(mgo.Change{Update: bson.M{"$set": bson.M{"use": value}}, ReturnNew: true, Upsert: true}, &info)
	if err != nil {
		return
	}
	return
}

//获取自增序列
func MgoGetInc(fn getMongoFn, key string, count int64) (id int64, err error) {
	session, err := fn()
	if err != nil {
		return
	}
	defer session.Close()

	var info MgoGenInfo

	client := session.DB(MGOKEY_DB_GLDB).C(MGOKEY_TABLE_GEN)
	err = client.Find(bson.M{"_id": key}).One(&info)
	if err != nil {
		return
	}
	id = info.Use
	return
}
