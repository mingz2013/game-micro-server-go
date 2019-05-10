package table_manager

import "testing"
import "encoding/json"

func TestNewTableApp(t *testing.T) {
	confMap := map[string]interface{}{

		"host":    "redis-mq",
		"port":    "6379",
		"db":      1,
		"channel": "table-manager-server-go",
	}
	data, _ := json.Marshal(confMap)
	a := NewApp(data)
	a.Start()
}
