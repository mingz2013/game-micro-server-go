package server

import "testing"
import "encoding/json"

func TestNewApp(t *testing.T) {

	confMap := map[string]interface{}{

		"host":    "localhost",
		"port":    "6379",
		"db":      1,
		"channel": "connector-server",
	}

	data, _ := json.Marshal(confMap)

	a := NewApp(data)
	a.Start()
}
