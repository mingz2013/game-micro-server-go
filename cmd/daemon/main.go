package main

import "encoding/json"
import "github.com/mingz2013/daemon-server-go/server"

func main() {
	confMap := map[string]map[string]interface{}{
		"redisChannelConf": {
			"host":    "redis-mq",
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
	a := server.NewApp(data)
	a.Start()
}
