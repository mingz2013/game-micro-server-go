package actor

type Actor struct {
	//server  net_base.Server
	//handler net_base.Handler
	Config *Config

	mailbox chan interface{}

	callbacks []chan interface{}
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

	a.mailbox = make(chan interface{}, 1024)

	//log.Println("Actor.Init...PROTO_TYPE", a.Config.protocol)
	//switch a.Config.protocol {
	//case net_base.PROTO_TCP:
	//	a.server = net_tcp.NewServer(a.Config.address)
	//case net_base.PROTO_WS:
	//	a.server = net_ws.NewServer(a.Config.address)
	//default:
	//	log.Println("error...", a.Config)
	//
	//}

}

//func (a *Actor) SetHandler(handler net_base.Handler) {
//	a.handler = handler
//	//a.server.SetHandler(a.handler)
//}

func (a *Actor) Start() {

	//a.server.StartServer()

}

func (a *Actor) Close() {
	//a.server.CloseServer()
}

func (a *Actor) Send() {
	// 异步发送，不需要返回
}

func (a *Actor) Query() {
	// 同步请求，需要回调
}

func (a *Actor) ReceiveMail() {
	// 往mailbox塞消息
}

func (a *Actor) ReadMail() {
	// 读消息
}
