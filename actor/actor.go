package actor

import (
	"github.com/mingz2013/lib-go/net_base"
	"github.com/mingz2013/lib-go/net_base/net_tcp"
	"github.com/mingz2013/lib-go/net_base/net_ws"
	"log"
)

type Actor struct {
	server  net_base.Server
	handler net_base.Handler
	Config  *Config
}

func NewActor(conf string) *Actor {
	a := &Actor{}
	a.Init(conf)
	return a
}

func (a *Actor) Init(conf string) {
	// 从数据库读取config，config init

	//a.Handler = NewConnectorApp()
	a.Config = NewConfig()
	a.Config.ParseFromStr(conf)

	log.Println("Actor.Init...PROTO_TYPE", a.Config.protocol)

	switch a.Config.protocol {
	case net_base.PROTO_TCP:
		a.server = net_tcp.NewServer(a.Config.address)
	case net_base.PROTO_WS:
		a.server = net_ws.NewServer(a.Config.address)
	default:
		log.Println("error...", a.Config)

	}

}

func (a *Actor) SetHandler(handler net_base.Handler) {
	a.handler = handler
	a.server.SetHandler(a.handler)
}

func (a *Actor) Start() {

	a.server.StartServer()

}

func (a *Actor) Close() {
	a.server.CloseServer()
}
