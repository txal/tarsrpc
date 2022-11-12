package util_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"tarsrpc/kissgo/util"
)

func doingKey(t *testing.T, c *util.CacheInfo, key string, mustExist bool, mustSize int64) {
	r, err := c.GetCache(key)
	if mustExist {
		if err != nil {
			t.Fatalf("%s must exist", key)
		}
		if int64(len(r)) != mustSize {
			t.Fatalf("%s size error %d!=%d", key, len(r), mustSize)
		}
	} else {
		if err == nil {
			t.Fatalf("%s must not exist", key)
		}
	}
}

func doingStat(t *testing.T, c *util.CacheInfo, totalSize int64, totalCount int64, validCount int64) {
	r1, r2, r3 := c.Stat()
	if r1 != totalSize || r2 != totalCount || r3 != validCount {
		t.Fatalf("stat diff need:%d,%d,%d  but:%d,%d,%d", totalSize, totalCount, validCount, r1, r2, r3)
	}
}

func sleep(msec int64) {
	time.Sleep(time.Millisecond * time.Duration(msec))
}

func TestCacheQuick(t *testing.T) {
	c := util.NewCache(500, 0)

	//set cache
	doingStat(t, c, 0, 0, 0)
	c.SetCache("key1", make([]byte, 10))
	c.SetCache("key2", make([]byte, 20))
	doingStat(t, c, 30, 2, 2)

	//get cache
	doingKey(t, c, "key1", true, 10)
	doingKey(t, c, "key2", true, 20)
	doingKey(t, c, "key3", false, 0)

	//sleep
	sleep(501)

	//get cache
	doingStat(t, c, 30, 2, 0)
	doingKey(t, c, "key1", false, 10)

	return
}

func TestCacheWait(t *testing.T) {
	c := util.NewCache(150, 100)

	fn := func(id int64) {
		idstr := fmt.Sprintf("#%d", id)
		t.Log(idstr, "get start")
		r, err := c.GetCacheOrWait("key")
		if err == nil {
			t.Log(idstr, "get success size =", len(r))
		} else {
			t.Log(idstr, "get empty and doing")
			sleep(90)
			t.Log(idstr, "set success")
			c.SetCache("key", make([]byte, id))
		}
	}

	for k := 1; k <= 20; k++ {
		go fn(int64(k))
		sleep(50)
	}
	sleep(200)
	return
}

func testWg(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	for k := 1; k <= 3; k++ {
		kk := k
		go func() {
			wg.Wait()
			t.Log("gowait ", kk)
		}()
	}
	sleep(1)

	t.Log("sleep 1")
	sleep(100)

	wg.Done()
	sleep(1)

	t.Log("sleep 2")
	sleep(100)

	t.Log("sleep end")
	return
}
