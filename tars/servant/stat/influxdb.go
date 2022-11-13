package stat

import (
	"context"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/http"

	"code.com/tars/goframework/kissgo/appzaplog"
	"code.com/tars/goframework/kissgo/appzaplog/zap"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

const bucketName = "api"

type influxdb struct {
	writeAPI api.WriteAPI
}

func newInfluxdb(url, token, org string) (*influxdb, error) {
	client := influxdb2.NewClientWithOptions(url, token, influxdb2.DefaultOptions().SetFlushInterval(5000))

	// 这里有两个作用
	// 1.检查bucket是否存在
	// 2.检查token有效，曾经因为tars解析配置出错，解析出错误的token,数据写不进去，花费数个小时，fuck
	_, err := client.BucketsAPI().FindBucketByName(context.TODO(), bucketName)
	if err != nil {
		return nil, err
	}

	writeAPI := client.WriteAPI(org, bucketName)
	writeAPI.SetWriteFailedCallback(func(batch string, error http.Error, retryAttempts uint) bool {
		appzaplog.Error("influxdb failed callback", zap.String("batch", batch), zap.Any("error", error), zap.Uint("retry", retryAttempts))
		if retryAttempts > 3 {
			return false
		}
		return true
	})
	return &influxdb{writeAPI: writeAPI}, nil
}

func (f *influxdb) write(node *statNode) {
	appzaplog.Debug("influxdb write", zap.Any("node", node))

	for k, v := range node.Data {
		tags := map[string]string{
			"namespace": k.Namespace,
			"server":    k.Server,
			"object":    k.Object,
			"function":  k.Function,
			"client_ip": k.ClientIP,
			"server_ip": k.ServerIP,
		}
		fields := map[string]interface{}{
			"total_count":   v.totalCount,
			"error_count":   v.errorCount,
			"timeout_count": v.timeoutCount,
			"cost":          v.totalCost / v.totalCount,
		}

		f.writeAPI.WritePoint(influxdb2.NewPoint("api", tags, fields, time.Unix(node.Timestamp, 0)))
		appzaplog.Debug("influxdb write point", zap.Any("tags", tags), zap.Any("fields", fields), zap.Int64("t", node.Timestamp))
	}
}
