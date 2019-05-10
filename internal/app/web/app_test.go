package server

import "testing"
import "encoding/json"

func TestNewApp(t *testing.T) {

	confMap := map[string]map[string]interface{}{
		"redisChannelConf": {
			"host":    "localhost",
			"port":    "6379",
			"db":      1,
			"channel": "connector-server",
		},
		"api": {
			"host": "localhost",
			"port": "8000",
		},
	}
	data, _ := json.Marshal(confMap)
	a := NewApp(data)
	a.Start()
}
