package server

import (
	"encoding/json"
	"github.com/mingz2013/lib-go/actor"
)

type App struct {
	redisMQActor *actor.RedisChannelActor
}

func NewApp(conf []byte) *App {
	a := &App{}
	a.Init(conf)
	return a
}

func (a *App) Init(conf []byte) {
	var confMap map[string]interface{}
	data := conf
	json.Unmarshal(data, &confMap)
	redisChannelConf := confMap["redisChannelConf"].(map[string]interface{})

	data, _ = json.Marshal(redisChannelConf)

	a.redisMQActor = actor.NewRedisChannelActor(string(data))

	//connectorConf := confMap["api"].(map[string]interface{})
	//data, _ = json.Marshal(connectorConf)

	a.redisMQActor.SetHandler(a)

}

func (a *App) OnRedisChannelMessage(message []byte) (retMsg []byte) {
	// 处理消息队列里面来的消息
	retMsg = message

	return
}

func (a *App) Start() {
	//
	a.redisMQActor.Start()
	a.redisMQActor.Wait()

}
