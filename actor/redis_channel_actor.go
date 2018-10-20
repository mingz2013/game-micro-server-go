package actor

import (
	"encoding/json"
	"github.com/mingz2013/lib-go/mq/redismq"
	"log"
)

type RedisChannelActor struct {
	//server  net_base.Server
	//handler net_base.Handler
	Config *Config

	//mailbox chan interface{}

	callbacks map[string]map[int64]chan Mail // {"channel": {"mark_id": chan}}

	channel string // 频道id

	redisMQClient *redismq.RedisMQClient
}

func NewRedisChannelActor(conf string) *RedisChannelActor {
	a := &RedisChannelActor{}
	a.Init(conf)
	return a
}

func (a *RedisChannelActor) Init(conf string) {
	// 从数据库读取config，config init

	//a.Handler = NewConnectorApp()
	a.Config = NewConfig()
	a.Config.ParseFromStr(conf)

	a.channel = "123"

	a.callbacks = make(map[string]map[int64]chan Mail)

	a.redisMQClient = redismq.NewRedisMQClient(conf)

	a.redisMQClient.Subscribe(a.channel, a.OnMessage, a.onSubscription)

	//a.mailbox = make(chan interface{}, 1024)

}

//func (a *Actor) SetHandler(handler net_base.Handler) {
//	a.handler = handler
//	//a.server.SetHandler(a.handler)
//}

func (a *RedisChannelActor) Start() {

	//a.server.StartServer()

}

func (a *RedisChannelActor) Close() {
	//a.server.CloseServer()
}

func (a *RedisChannelActor) SendMail(mail Mail) {
	// 异步发送，不需要返回

	data, _ := json.Marshal(mail)

	a.redisMQClient.Pubscribe(mail.to, data)

}

func (a *RedisChannelActor) SendMailNeedBack(mail Mail) Mail {
	// 同步请求，需要回调

	data, _ := json.Marshal(mail)

	a.redisMQClient.Pubscribe(mail.to, data)

	channelmails, ok := a.callbacks[mail.to]
	if !ok {
		channelmails = make(map[int64]chan Mail)
		a.callbacks[mail.to] = channelmails

	}

	_, ok = channelmails[mail.mark]
	if ok {
		log.Fatalln("err....., mark_id has already exits")

	}
	channelmails[mail.mark] = make(chan Mail)

	return <-channelmails[mail.mark]

}

func (a *RedisChannelActor) OnMessage(channel string, data []byte) {
	if a.channel != channel {
		log.Fatalln("err...channel not equal")
		return
	}
	var mail Mail
	json.Unmarshal(data, mail)

	a.ReceiveMail(mail)

}

func (a *RedisChannelActor) onSubscription(channel string, kind string, count int) {

}

func (a *RedisChannelActor) ReceiveMail(mail Mail) {
	if mail.isBack {
		a.callbacks[mail.from][mail.mark] <- mail
		return
	}

	message := a.ReadMail(mail)

	if mail.needBack {
		backmail := NewMail(a.channel, mail.from, message, false, true)
		data, _ := json.Marshal(backmail)
		a.redisMQClient.Pubscribe(backmail.to, data)

	}

}

func (a *RedisChannelActor) ReadMail(mail Mail) (message []byte) {
	// 读消息
	return
}
