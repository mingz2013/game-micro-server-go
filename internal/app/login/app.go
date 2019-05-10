package login

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/mingz2013/game-micro-server-go/internal/app/web/controllers"
	"github.com/mingz2013/game-micro-server-go/internal/pkg/actor"
	"github.com/mingz2013/game-micro-server-go/internal/pkg/net_base"
	"gopkg.in/go-playground/validator.v8"
	"net/http"

	"log"
)

// 桌子进程，很有可能，只是个客户端，连接到中心服务器，不用server监听

type App struct {
	redisChannelActor *actor.RedisChannelActor
}

func (a *App) Init(conf []byte) {
	a.redisChannelActor = actor.NewRedisChannelActor(string(conf))
	a.redisChannelActor.SetHandler(a)
}

func NewApp(conf []byte) *App {
	a := &App{}
	a.Init(conf)
	return a
}

func (a *App) Start() {
	a.redisChannelActor.Start()
	//a.manager.Start()
	a.StartHTTP()
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
	//
	//a.manager.MsgIn <- js

}

func notFound(c echo.Context) error {
	return c.String(http.StatusNotFound, "not found")
}

func (a *App) StartHTTP() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New(nil)}

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "hello, world!")
	})

	controllers.RegisterRouters(e)

	e.Any("/", notFound)

	e.Static("/static", "./static")

	e.Logger.Fatal(e.Start(":8001"))
}
