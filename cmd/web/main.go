package main

import "github.com/mingz2013/web-proxy-server-go/server"
import (
	"encoding/json"
	//"github.com/mingz2013/web-proxy-server-go/global"
)

func main() {

	confMap := map[string]map[string]interface{}{
		"redisChannelConf": {
			"host":    "redis-mq",
			"port":    "6379",
			"db":      1,
			"channel": "web-proxy-server-go",
		},
		"api": {
			"host": "localhost",
			"port": "8006",
		},
	}
	data, _ := json.Marshal(confMap)
	a := server.NewApp(data)
	//global.DefaultApp = a
	a.Start()
}
