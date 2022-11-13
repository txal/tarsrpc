package stat

import (
	"container/list"
	"strings"
	"sync"
	"time"

	"code.com/tars/goframework/kissgo/appzaplog/zap"

	"code.com/tars/goframework/kissgo/appzaplog"
)

type Tag struct {
	Namespace string // 命名空间
	Server    string // 服务名
	Object    string // 对象
	Function  string // 方法
	ClientIP  string // 客户端IP
	ServerIP  string // 服务端IP
}

type field struct {
	totalCount   int64 // 总数
	errorCount   int64 // 错误数  9999，返回
	timeoutCount int64 // 超时数
	totalCost    int64 // 单位时间内总耗时
}

func (f *field) increase(cost int64, isError, isTimeout bool) {
	f.totalCount++
	f.totalCost += cost

	if isError {
		f.errorCount++
	}
	if isTimeout {
		f.timeoutCount++
	}
}

type statNode struct {
	Timestamp int64
	Data      map[Tag]*field
}

type stat struct {
	list    list.List
	mutex   sync.Mutex
	storage *influxdb
}

var statInstance *stat

// Init 初始化stat
func Init(url, token, org string) error {
	token = strings.ReplaceAll(token, `\`, ``)
	appzaplog.Info("stat init", zap.String("url", url), zap.String("token", token), zap.String("org", org))
	influxdb, err := newInfluxdb(url, token, org)
	if err != nil {
		return err
	}
	statInstance = &stat{storage: influxdb}
	appzaplog.Info("stat init finish")
	go statInstance.timedPush()
	return nil
}

// Report 上报
func Report(tag Tag, cost int64, isError, isTimeout bool) {
	report(time.Now(), tag, cost, isError, isTimeout)
}

// report 提供时间设置，方便做单元测试
func report(t time.Time, tag Tag, cost int64, isError, isTimeout bool) {
	statInstance.mutex.Lock()
	defer statInstance.mutex.Unlock()

	// 时间精度改为分钟级别
	timestamp := minuteUnix(t)

	front := statInstance.list.Front()
	if front != nil {
		node := front.Value.(*statNode)
		if node.Timestamp == timestamp {
			if f, ok := node.Data[tag]; ok {
				f.increase(cost, isError, isTimeout)
			} else {
				f := new(field)
				f.increase(cost, isError, isTimeout)
				node.Data[tag] = f
			}
			return
		}
	}

	f := new(field)
	f.increase(cost, isError, isTimeout)
	node := statNode{
		Timestamp: timestamp,
		Data:      map[Tag]*field{tag: f},
	}
	statInstance.list.PushFront(&node)
}

// timedPush 定时推送
func (s *stat) timedPush() {
	ticker := time.NewTicker(10 * time.Second)

	for range ticker.C {
		s.push()
	}
}

func (s *stat) push() {
	var node = func() *statNode {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		back := s.list.Back()
		if back == nil {
			return nil
		}
		node := back.Value.(*statNode)
		if node.Timestamp >= minuteUnix(time.Now()) {
			return nil
		}
		s.list.Remove(back)
		return node
	}()

	if node == nil {
		return
	}
	s.storage.write(node)
}

func minuteUnix(t time.Time) int64 {
	return t.Unix() / 60 * 60
}
