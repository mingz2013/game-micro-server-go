package actor

import (
	"encoding/json"
	"github.com/mingz2013/lib-go/mq/redismq"
	"log"
	"sync"
)

type RedisChannelActor struct {
	//server  net_base.Server
	//handler net_base.Handler
	Config *Config

	//mailbox chan interface{}

	callbacks      map[string]map[int64]chan Mail // {"channel": {"mark_id": chan}}
	callbacksMutex sync.Mutex

	channel string // 频道id, 同时也是actorid

	redisMQClient *redismq.RedisMQClient

	handler RedisChannelActorHandler
}

func NewRedisChannelActor(conf string) *RedisChannelActor {
	a := &RedisChannelActor{}
	a.Init(conf)
	return a
}

func (a *RedisChannelActor) SetHandler(handler RedisChannelActorHandler) {
	a.handler = handler
}

func (a *RedisChannelActor) Init(conf string) {
	// 从数据库读取config，config init

	//a.Handler = NewConnectorApp()
	a.Config = NewConfig()
	//a.Config.ParseFromStr(conf)

	var confJs map[string]interface{}
	err := json.Unmarshal([]byte(conf), &confJs)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(confJs)

	a.channel = confJs["channel"].(string)

	a.callbacks = make(map[string]map[int64]chan Mail)

	a.redisMQClient = redismq.NewRedisMQClient(conf)

	//a.mailbox = make(chan interface{}, 1024)

}

//func (a *Actor) SetHandler(handler net_base.Handler) {
//	a.handler = handler
//	//a.server.SetHandler(a.handler)
//}

func (a *RedisChannelActor) Start() {

	a.redisMQClient.Subscribe(a.channel, a.OnMessage, a.onSubscription)

	//a.Wait()

}

func (a *RedisChannelActor) Close() {
	//a.server.CloseServer()
	a.redisMQClient.Unsubscribe(a.channel)
}

func (a *RedisChannelActor) SendMail(mail Mail) {
	// 异步发送，不需要返回

	data, _ := json.Marshal(mail)

	a.redisMQClient.Pubscribe(mail.To, data)

}

func (a *RedisChannelActor) SendMailNeedBack(mail Mail) Mail {
	// 同步请求，需要回调

	log.Println("SendMailNeedBack", "mail", mail)

	data, _ := json.Marshal(&mail)

	log.Println("SendMailNeedBack", "data", data)

	a.callbacksMutex.Lock()
	channelmails, ok := a.callbacks[mail.To]
	if !ok {
		channelmails = make(map[int64]chan Mail)
		a.callbacks[mail.To] = channelmails

	}

	_, ok = channelmails[mail.Mark]
	if ok {
		log.Fatalln("err....., mark_id has already exits")

	}
	channelmails[mail.Mark] = make(chan Mail)

	log.Println("chan, make", channelmails[mail.Mark], mail.To, mail.Mark)
	a.callbacksMutex.Unlock()

	a.redisMQClient.Pubscribe(mail.To, data)

	log.Println("wait back....")
	a.callbacksMutex.Lock()
	log.Println("chan, receive", channelmails[mail.Mark])
	retmail := <-channelmails[mail.Mark]
	log.Println("receive back", retmail)

	delete(channelmails, mail.Mark)
	a.callbacksMutex.Unlock()

	return retmail

}

func (a *RedisChannelActor) OnMessage(channel string, data []byte) {
	log.Println("OnMessage<"+channel+">", data)
	if a.channel != channel {
		log.Fatalln("err...channel not equal")
		return
	}
	var mail Mail
	json.Unmarshal(data, &mail)

	a.ReceiveMail(mail)

}

func (a *RedisChannelActor) onSubscription(channel string, kind string, count int) {

}

func (a *RedisChannelActor) ReceiveMail(mail Mail) {

	log.Println("ReceiveMail", mail, mail.IsBack)

	if mail.IsBack {
		log.Println("ReceiveMail", "IsBack", mail)
		log.Println("chan, send", a.callbacks[mail.From][mail.Mark], mail.From, mail.Mark)
		//a.callbacks[mail.From][mail.Mark] <- mail
		channelmails, ok := a.callbacks[mail.From]
		if !ok {
			log.Fatalln("sendback not ok")
		}
		channelmail, ok := channelmails[mail.Mark]
		if !ok {
			log.Fatalln("channel mail mark not ok")
		}
		channelmail <- mail

		log.Println("return ......ReceiveMail")

		return
	}

	message := a.ReadMail(mail)

	if mail.NeedBack {
		backmail := NewMail(a.channel, mail.From, message, false, true, mail.Mark)
		data, _ := json.Marshal(backmail)
		a.redisMQClient.Pubscribe(backmail.To, data)

	}

}

func (a *RedisChannelActor) ReadMail(mail Mail) (message []byte) {
	// 读消息
	//message = mail.Message
	message = a.handler.OnRedisChannelMessage(mail.Message)
	return
}

func (a *RedisChannelActor) Wait() {
	log.Println("RedisChannelActor.Wait...")
	a.redisMQClient.Wait()
}
