package util_test

import (
	"testing"
	"time"

	"code.com/tars/goframework/kissgo/util"
	"gopkg.in/mgo.v2"
)

var g_lockKey = "testlock1"
var g_genkey = "testgen1"

func TestMgoLock(t *testing.T) {
	util.MgoUnlock(getMongo, g_lockKey)

	//上锁成功
	if util.MgoLock(getMongo, g_lockKey, 2) != true {
		t.Error("lock1 fail")
	}
	//重复上锁也会成功
	if util.MgoLock(getMongo, g_lockKey, 2) != true {
		t.Error("lock2 fail")
	}
	//解锁会成功
	if util.MgoUnlock(getMongo, g_lockKey) != true {
		t.Error("unlock1 fail")
	}
	return
}

func TestMgoLockTimeout(t *testing.T) {
	util.MgoUnlock(getMongo, g_lockKey)

	//第一次上锁成功
	if util.MgoLock(getMongo, g_lockKey, 1) != true {
		t.Error("lock1 fail")
	}

	//让锁超时
	time.Sleep(time.Second * 2)

	//超时后上锁成功
	if util.MgoLock(getMongo, g_lockKey, 1) != true {
		t.Error("lock2 fail")
	}
}

//func TestMgoLockBatch(t *testing.T) {
//	util.MgoUnlock(getMongo, g_lockKey)

//	succ := 0
//	for k := 0; k < 20; k++ {
//		go func() {
//			if util.MgoLock(getMongo, g_lockKey, 5) == true {
//				succ += 1
//				//t.Log("I am working")
//			} else {
//				//t.Log("I am skip")
//			}
//		}()
//	}
//	time.Sleep(time.Second * 3)
//	if succ != 1 {
//		t.Error("lock > 1 count=", succ)
//	}
//}

func TestMgoGen(t *testing.T) {
	util.MgoSetInc(getMongo, g_genkey, 0)

	//正常递增
	if id, err := util.MgoIncOne(getMongo, g_genkey); id != 1 {
		t.Error("error", id, err)
	}
	if id, err := util.MgoIncOne(getMongo, g_genkey); id != 2 {
		t.Error("error", id, err)
	}
	if id, err := util.MgoIncMore(getMongo, g_genkey, 1); id != 3 {
		t.Error("error", id, err)
	}
	if id, err := util.MgoIncMore(getMongo, g_genkey, 10); id != 4 {
		t.Error("error", id, err)
	}
	if id, err := util.MgoIncMore(getMongo, g_genkey, 20); id != 14 {
		t.Error("error", id, err)
	}
	if id, err := util.MgoIncMore(getMongo, g_genkey, 1); id != 34 {
		t.Error("error", id, err)
	}
	//重置后使用
	if err := util.MgoSetInc(getMongo, g_genkey, 1000); err != nil {
		t.Error("error", err)
	}
	if id, err := util.MgoIncOne(getMongo, g_genkey); id != 1001 {
		t.Error("error", id, err)
	}
}

func getMongo() (session *mgo.Session, err error) {
	session, err = mgo.Dial("101.226.20.73:32001")
	if err != nil {
		return
	}
	session.SetMode(mgo.Monotonic, true)
	return
}
