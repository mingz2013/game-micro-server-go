package actor

import (
	"github.com/mingz2013/lib-go/net_base"
	"github.com/mingz2013/lib-go/net_base/net_tcp"
	"github.com/mingz2013/lib-go/net_base/net_ws"
	"log"
)

type ConnectorActor struct {
	server  net_base.Server
	handler net_base.Handler
	Config  *Config

	mailbox chan interface{}
}

func NewConnectorActor(conf string) *ConnectorActor {
	a := &ConnectorActor{}
	a.Init(conf)
	return a
}

func (a *ConnectorActor) Init(conf string) {
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

func (a *ConnectorActor) SetHandler(handler net_base.Handler) {
	a.handler = handler
	a.server.SetHandler(a.handler)
}

func (a *ConnectorActor) Start() {

	a.server.StartServer()

}

func (a *ConnectorActor) Close() {
	a.server.CloseServer()
}

func (a *ConnectorActor) Send() {
	// 异步发送，不需要返回
}

func (a *ConnectorActor) Query() {
	// 同步请求，需要回调
}
