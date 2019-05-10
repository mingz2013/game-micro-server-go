package main

import (
	"encoding/json"
	"github.com/mingz2013/connector-server-go/app"
	"log"
)

func parseArgs() {
	// 根据参数，确定是什么协议的服务，还有service id

}

func main() {
	log.Println("main...")

	confMap := map[string]map[string]interface{}{
		"redisChannelConf": {
			"host":    "redis-mq",
			"port":    "6379",
			"db":      1,
			"channel": "connector-server",
		},
		"connectorConf": {
			"host": "localhost",
			"port": "8000",
		},
	}
	data, _ := json.Marshal(confMap)
	a := app.NewApp(data)
	a.Start()

}
