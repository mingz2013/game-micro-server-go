package table_manager

import (
	"encoding/json"
	"github.com/mingz2013/lib-go/internal/pkg/actor"
	"github.com/mingz2013/lib-go/internal/pkg/msg"
	"github.com/mingz2013/lib-go/internal/pkg/net_base"
	"log"
	"time"
)

// 桌子进程，很有可能，只是个客户端，连接到中心服务器，不用server监听

type App struct {
	redisChannelActor *actor.RedisChannelActor

	manager Manager
	msgOut  chan msg.Msg
	msgIn   chan msg.Msg
}

func (a *App) Init(conf []byte) {
	a.redisChannelActor = actor.NewRedisChannelActor(string(conf))
	a.redisChannelActor.SetHandler(a)

	a.msgOut = make(chan msg.Msg)
	a.msgIn = make(chan msg.Msg)

	a.manager = NewTableManager("111", a.msgIn, a.msgOut)
}

func NewApp(conf []byte) *App {
	a := &App{}
	a.Init(conf)
	return a
}

func (a *App) Start() {
	a.redisChannelActor.Start()
	a.manager.Run()
	a.redisChannelActor.Wait()

	for {
		select {
		case m, ok := <-a.msgOut:
			{
				if !ok {
					continue
				}
				data, _ := json.Marshal(m)
				a.redisChannelActor.Send("", data)
			}
		case <-time.After(1 * time.Second):
			{
				continue
			}

		}
	}

}

func (a *App) OnRedisChannelMessage(message []byte) (retMsg []byte) {
	retMsg = message
	return
}

func (a *App) Serve(c net_base.Conn, buf []byte) {
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
	//cmd := js["cmd"].(string)
	//userId := js["userId"].(int)
	//token := js["token"].(string)
	//// 验证token和userId

	m := msg.Msg(js)
	a.msgIn <- m

}
