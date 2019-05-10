package server

import "github.com/mingz2013/lib-go/actor"
import (
	"encoding/json"
	"github.com/mingz2013/game-manager-server-go/manager"
	"github.com/mingz2013/lib-go/msg"
)

type App struct {
	redisChannelActor *actor.RedisChannelActor
	gameManager       *manager.GameManager
}

func NewApp(conf []byte) (a *App) {
	a = &App{}
	a.Init(conf)
	return a
}

func (a *App) Init(conf []byte) {
	var confMap map[string]interface{}
	data := conf
	json.Unmarshal(data, &confMap)
	redisChannelConf := confMap["redisChannelConf"].(map[string]interface{})

	data, _ = json.Marshal(redisChannelConf)

	a.redisChannelActor = actor.NewRedisChannelActor(string(data))

	a.gameManager = manager.NewGameManager()
	//a.gameManager.SetDelegate(a.redisChannelActor)

}

func (a *App) Start() {
	a.redisChannelActor.Start()
	a.redisChannelActor.Wait()
}

func (a *App) OnRedisChannelMessage(message []byte) (retMsg []byte) {
	retMsg = message

	var msgObj msg.Msg

	json.Unmarshal(message, msgObj)

	var retMsgObj msg.Msg

	json.Unmarshal(retMsg, retMsgObj)

	switch msgObj.GetCmd() {
	case "create":
		tableId, err := a.gameManager.CreateTable(msgObj.GetParam("userId").(int), msgObj.GetParam("gameId").(int))
		if err != nil {
			return
		} else {
			retMsgObj.SetResults(map[string]interface{}{
				"tableId": tableId,
			})
		}

	case "del":
		a.gameManager.DelTable()

	}

	return
}
