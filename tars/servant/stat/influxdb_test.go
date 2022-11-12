package stat

import (
	"fmt"
	"testing"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/http"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestInitTag(t *testing.T) {
	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPI(org, "api")
	writeAPI.SetWriteFailedCallback(func(batch string, error http.Error, retryAttempts uint) bool {
		fmt.Println(batch, error, retryAttempts)
		return false
	})

	p := influxdb2.NewPoint(
		"api",
		map[string]string{
			"namespace": "all",
			"server":    "all",
			"object":    "all",
			"function":  "all",
			"client_ip": "all",
			"server_ip": "all",
		},
		map[string]interface{}{
			"total_count":   0,
			"error_count":   0,
			"timeout_count": 0,
			"cost":          0,
		},
		time.Unix(time.Now().Unix(), 0))
	writeAPI.WritePoint(p)
	writeAPI.Flush()
}
