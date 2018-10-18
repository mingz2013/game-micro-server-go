package actor

import (
	"github.com/mingz2013/lib-go/net/tcp"
	"github.com/mingz2013/lib-go/net/ws"
	"log"
)

type Actor struct {
	Server  net_frame.Server
	Handler net_frame.Handler
	Config  *Config
}

func NewActor(conf string) *Actor {
	a := &Actor{}
	a.Init(conf)
	return a
}

func (a *Actor) Init(conf string) {
	// 从数据库读取config，config init

	a.Handler = NewConnectorApp()
	a.Config = NewConfig()
	a.Config.ParseFromStr(conf)

	log.Println("Actor.Init...PROTO_TYPE", a.Config.protocol)

	switch a.Config.protocol {
	case net_frame.PROTO_TCP:
		a.Server = tcp.NewServer(a.Config.address)
	case net_frame.PROTO_WS:
		a.Server = ws.NewServer(a.Config.address)
	default:
		log.Println("error...", a.Config)

	}

	a.Server.SetHandler(a.Handler)

}

func (a *Actor) Start() {

	a.Server.Start()

}

func (a *Actor) Close() {
	a.Server.Close()
}
