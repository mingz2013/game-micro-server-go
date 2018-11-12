package robot_manager

import (
	"encoding/json"
	"github.com/mingz2013/lib-go/actor"
	"github.com/mingz2013/lib-go/msg"
	"log"
)

type App struct {
	redisChannelActor *actor.RedisChannelActor
	rbManager         *RobotManager
	MsgIn             chan msg.Msg
	MsgOut            chan msg.Msg
}

func NewApp(conf []byte) (a *App) {
	a = &App{}
	a.Init(conf)
	return a
}

func (a *App) Init(conf []byte) {
	a.MsgIn = make(chan msg.Msg)
	a.MsgOut = make(chan msg.Msg)
	a.rbManager = NewRobotManager(a.MsgIn, a.MsgOut)

	a.redisChannelActor = actor.NewRedisChannelActor(string(conf))
	a.redisChannelActor.SetHandler(a)

}

func (a *App) OnRedisChannelMessage(message []byte) (retMsg []byte) {
	//retMsg = message
	// 消息分两部分，一部分是用户消息，一部分是管理消息

	var mJs msg.Msg

	json.Unmarshal(message, &mJs)

	if mJs.GetCmd() == "robot" && mJs.GetParams()["action"] == "start" {
		// TODO 开启一个机器人去连接co，去创建房间
		log.Println("on robot start....")
	}

	return
}

func (a *App) Start() {
	a.redisChannelActor.Start()
	a.redisChannelActor.Wait()
}
