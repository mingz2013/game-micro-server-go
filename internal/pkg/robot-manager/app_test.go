package robot_manager

import (
	"encoding/json"
	"testing"
)

func TestNewApp(t *testing.T) {
	confMap := map[string]interface{}{

		"host":    "redis-mq",
		"port":    "6379",
		"db":      1,
		"channel": "robot-manager-server-go",
	}

	data, _ := json.Marshal(confMap)

	a := NewApp(data)
	a.Start()
}
