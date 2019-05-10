package web

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/mingz2013/lib-go/actor"
	"log"
	"net/http"
)

type App struct {
	redisMQActor *actor.RedisChannelActor
}

func NewApp(conf []byte) *App {
	//if defaultapp != nil {
	//	return DefaultApp
	//}
	app := &App{}
	app.Init(conf)
	return app
}

func (a *App) Init(conf []byte) {
	var confMap map[string]interface{}
	data := conf
	json.Unmarshal(data, &confMap)
	redisChannelConf := confMap["redisChannelConf"].(map[string]interface{})

	data, _ = json.Marshal(redisChannelConf)

	a.redisMQActor = actor.NewRedisChannelActor(string(data))

	connectorConf := confMap["api"].(map[string]interface{})
	data, _ = json.Marshal(connectorConf)

	a.redisMQActor.SetHandler(a)

}

//func (a *App) SendMail(mail actor.Mail) {
//	a.redisMQActor.SendMail(mail)
//}
//
//func (a *App) SendMsg(msg []byte) {
//	mail := actor.Mail{
//		Message: msg,
//		//From:a.redisMQActor.
//	}
//	a.SendMail(mail)
//}

func (a *App) RedisChannelActor() *actor.RedisChannelActor {
	return a.redisMQActor
}

func (a *App) OnRedisChannelMessage(message []byte) (retMsg []byte) {
	// 处理消息队列里面来的消息
	retMsg = message

	return
}

func (a *App) Start() {
	//

	a.redisMQActor.Start()
	a.StartHttp()

}

func (a *App) OnRobotStart(c echo.Context) error {
	log.Println("roboteStart...")
	msg := map[string]interface{}{
		"cmd": "manage",
		"param": map[string]interface{}{
			"action": "start",
		},
	}

	data, _ := json.Marshal(msg)
	backmsg := a.RedisChannelActor().Query("robot-server-go", data)

	c.JSON(http.StatusOK, "robot start..."+string(backmsg))
	return nil
}

func (a *App) StartHttp() {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "hello, world!")
	})

	//e.POST("/users", saveUser)
	//e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
	//
	//e.GET("/show", show)
	//e.POST("/save", save)
	//e.POST("/saveFile", saveFile)

	e.Static("/static", "./static")

	// ---

	e.GET("/robot/start/", a.OnRobotStart)

	e.Logger.Fatal(e.Start("0.0.0.0:8006"))
}
