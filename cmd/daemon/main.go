package main

import "encoding/json"
import "github.com/mingz2013/game-micro-server-go/internal/app/daemon"

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
	a := daemon.NewApp(data)
	a.Start()
}
