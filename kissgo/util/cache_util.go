/***************************************************
	缓存功能 by huangzhibin
	-------------------------------------------------
	1、key为string，value为[]byte。
	2、用于系统数据的缓存，不能用于用户数据
	3、http相应封装 http_util.go：HttpCacheReturn/HttpCacheRespJson
	优点：
	1、对于大量公共性的并发请求，能降低延时、降低redis/mongodb/cpu消耗
	2、如果多个请求同时获取同一个缓存，只会处理第一个请求，其他请求会阻塞等待直到超时
***************************************************/
package util

import (
	"errors"
	"sync"
	"time"
)

var ErrCacheNil = errors.New("cache nil")

//缓存结构
type CacheInfo struct {
	mu           sync.RWMutex          //锁
	cache        map[string]*cacheData //所有缓存数据
	cacheTimeout time.Duration         //缓存超时时间
	wait         map[string]*waitData  //所有等待数据
	waitTimeout  time.Duration         //等待超时时间（防止资源失效后一小段时间大量请求同时申请数据）
}

//缓存数据
type cacheData struct {
	buf         []byte    //缓存
	createTime  time.Time //创建时间
	updateTime  time.Time //更新时间
	expiredTime time.Time //有效时间
	wait        int64     //等待数
}

//等待数据
type waitData struct {
	createTime time.Time      //创建时间
	waitTime   time.Time      //等待结束时间
	wg         sync.WaitGroup //等待组
	r          []byte         //ret
	err        error          //ret
}

//申请缓存（cacheTimeout为缓存超时毫秒，最少1秒；waitTimeout为等待资源的超时毫秒）
func NewCache(cacheTimeout int64, waitTimeout int64) (c *CacheInfo) {
	c = &CacheInfo{}
	c.cacheTimeout = time.Millisecond * time.Duration(cacheTimeout)
	c.waitTimeout = time.Millisecond * time.Duration(waitTimeout)
	c.reset()
	go c.loop()
	return
}

//设置缓存（内部会复制一份buf）
func (c *CacheInfo) SetCache(key string, buf []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()

	var p *cacheData
	if v, ok := c.cache[key]; ok {
		p = v
	} else {
		p = &cacheData{
			createTime: now,
			wait:       0,
		}
		c.cache[key] = p
	}

	p.updateTime = now
	p.expiredTime = now.Add(c.cacheTimeout)

	p.buf = make([]byte, len(buf))
	copy(p.buf, buf)

	c.finishWait(key, p.buf, nil)
}

//获取缓存（不管缓存是否存在，都立即返回）（外部不能修改r）
func (c *CacheInfo) GetCache(key string) (r []byte, err error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	now := time.Now()

	if v, ok := c.cache[key]; ok {
		if now.Before(v.expiredTime) {
			r = v.buf
			return
		}
	}

	err = ErrCacheNil
	return
}

//获取缓存（缓存存在或者自己是第一个申请的，就立即返回；否则就等待缓存直到超时）（外部不能修改r）
func (c *CacheInfo) GetCacheOrWait(key string) (r []byte, err error) {
	r, err = c.GetCache(key)
	if err == nil {
		return
	}
	if c.waitTimeout == 0 {
		return
	}

	c.mu.Lock()
	if p, ok := c.wait[key]; ok {
		//等待别人取完数据
		c.mu.Unlock()
		p.wg.Wait()

		if p.err == nil {
			//等待成功
			return p.r, p.err
		} else {
			//等待失败，超时后会导致大量请求一起处理，考虑优化。。。 todo
			return p.r, p.err
		}
	}

	//必须自己去拿数据
	now := time.Now()
	p := &waitData{
		createTime: now,
		waitTime:   now.Add(c.waitTimeout),
	}
	p.wg.Add(1)
	c.wait[key] = p
	c.mu.Unlock()

	timer := time.NewTimer(c.waitTimeout)
	go func() {
		<-timer.C
		c.mu.Lock()
		c.finishWait(key, nil, ErrCacheNil)
		c.mu.Unlock()
	}()

	err = ErrCacheNil
	return
}

//获取统计数据（内存总大小、节点总数、有效节点数）
func (c *CacheInfo) Stat() (totalSize int64, totalCount int64, validCount int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()

	totalCount = int64(len(c.cache))
	for _, v := range c.cache {
		totalSize += int64(len(v.buf))
		if now.Before(v.expiredTime) {
			validCount += 1
		}
	}
	return
}

//清理过期缓存
func (c *CacheInfo) loop() {
	for {
		//每1分钟清理一次过期数据
		time.Sleep(time.Second * 60)
		now := time.Now()
		del := 0

		c.mu.Lock()
		for k, v := range c.cache {
			if now.After(v.expiredTime) {
				delete(c.cache, k)
				del += 1
			}
		}
		c.mu.Unlock()
	}
}

//重置数据
func (c *CacheInfo) reset() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[string]*cacheData)
	c.wait = make(map[string]*waitData)
	return
}

//处理结果
func (c *CacheInfo) finishWait(key string, r []byte, err error) {
	if p, ok := c.wait[key]; ok {
		delete(c.wait, key)
		p.r = r
		p.err = err
		p.wg.Done()
	}
}
