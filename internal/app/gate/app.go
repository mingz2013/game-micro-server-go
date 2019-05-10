package app

import (
	"encoding/json"
	"github.com/mingz2013/lib-go/actor"
	"github.com/mingz2013/lib-go/net_base"
	"log"
)

type App struct {
	//net_base.Handler
	//config *Config

	redisMQActor   *actor.RedisChannelActor
	connectorActor *actor.ConnectorActor

	//userMap map[*tcp.Conn]*ConnectorUser
	//connMap map[int]*net_base.Conn

	//msgQueuePool int // 消息队列连接池
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

	connectorConf := confMap["connectorConf"].(map[string]interface{})
	data, _ = json.Marshal(connectorConf)

	a.connectorActor = actor.NewConnectorActor(string(data))
	a.redisMQActor.SetHandler(a)
	a.connectorActor.SetHandler(a)

}

func (a *App) OnRedisChannelMessage(message []byte) (retMsg []byte) {
	// 处理消息队列里面来的消息
	retMsg = message

	// 大部分消息是转发给客户端
	// 一少部分，比如robot的消息，或者其他，需要转发给other

	return
}

func (a *App) OnConn(c net_base.Conn) (err error) {

	return
}

func (a *App) OnClose(c net_base.Conn) (err error) {
	return
}

func (a *App) Serve(c net_base.Conn, buf []byte) {
	// 处理客户端来的消息
	//userId := c.GetExtra()
	//var err error
	//
	//if userId == 0 {
	//	userId, err = a.createUser(r)
	//	if err != nil {
	//		//c.WriteResponse()
	//		return
	//	}
	//
	//	c.SetExtra(userId)
	//	a.connMap[userId] = &c
	//
	//}
	//
	//a.doRequest(c, r, userId)

	// 解析成json，ServeJson
	var js map[string]interface{}
	err := json.Unmarshal(buf, js)
	if err == nil {
		a.ServeJson(c, js)
	} else {
		log.Println(err, buf)
	}

}

func (a *App) ServeJson(c net_base.Conn, js map[string]interface{}) {
	// 前端发第一个协议，bind_user, 绑定用户连接，前端数据中应该有userId和token
	cmd := js["cmd"].(string)
	userId := js["userId"].(int)
	token := js["token"].(string)
	// 验证token和userId

	if cmd == "bind_user" {
		err := a.updateSession(c, cmd, userId, token, js)
		if err == nil {

		} else {

		}
	} else {
		err := a.checkToken(c, userId, token)
		if err != nil {
			return
		} else {
			// token验证通过，正常协议
			a.serveNormalMsg(c, cmd, userId, js)
		}
	}

}

func (a *App) updateSession(c net_base.Conn, cmd string, userId int, token string, msg map[string]interface{}) (err error) {
	// 更新session
	// 更新bind信息
	// 开启心跳
	return
}

func (a *App) checkToken(c net_base.Conn, userId int, token string) (err error) {
	// 验证token
	return
}

func (a *App) serveNormalMsg(c net_base.Conn, cmd string, userId int, msg map[string]interface{}) {
	if cmd == "keep_alive" {
		a.keepAlive(c, cmd, userId, msg)
	} else {
		// 转发消息到其他actor
		a.routeToOther(c, cmd, userId, msg)
	}
}

func (a *App) keepAlive(c net_base.Conn, cmd string, userId int, msg map[string]interface{}) {

}

func (a *App) routeToOther(c net_base.Conn, cmd string, userId int, msg map[string]interface{}) {

}

func (a *App) Start() {
	//
	a.redisMQActor.Start()

	a.connectorActor.Start()

}
