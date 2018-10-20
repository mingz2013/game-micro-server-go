package redismq

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mingz2013/lib-go/db/myredis"
	"log"
	"sync"
)

type RedisMQClient struct {
	myredis.RedisClient

	wgSub sync.WaitGroup

	channelMap map[string]redis.PubSubConn

	handler Handler
}

func (c *RedisMQClient) Init(conf string) {
	c.RedisClient.Init(conf)
	c.channelMap = make(map[string]redis.PubSubConn)
}

func NewRedisMQClient(conf string) *RedisMQClient {
	client := &RedisMQClient{}

	client.Init(conf)

	return client
}

func (c *RedisMQClient) Subscribe(channel string, onMessage func(channel string, data []byte), onSubscription func(channel string, kind string, count int)) {
	//redisChannel := "redChatRoom"
	conn := c.Pool.Get()
	psc := redis.PubSubConn{conn}
	psc.Subscribe(channel)
	c.channelMap[channel] = psc

	c.wgSub.Add(1)
	go func() {
		defer func() {
			log.Println("chan close...")
			conn.Close()
			psc.Unsubscribe(channel)
			c.wgSub.Done()
		}()

		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				log.Println("messages<", v.Channel, ">:", v.Data)
				onMessage(v.Channel, v.Data)
			case redis.Subscription:
				log.Println(v.Channel, v.Kind, v.Count)

				switch v.Kind {
				case "subscribe":
					onSubscription(v.Channel, v.Kind, v.Count)
				case "unsubscribe":
					return
				case "psubscribe":
					onSubscription(v.Channel, v.Kind, v.Count)
				case "punsubscribe":
					return
				default:
					return

				}

				continue
			case error:
				log.Println(v)
				return

			}
		}

	}()
}

func (c *RedisMQClient) Unsubscribe(channel string) {
	conn, ok := c.channelMap[channel]
	if !ok {
		return
	}
	conn.Unsubscribe(channel)
	conn.Close()
	delete(c.channelMap, channel)
}

func (c *RedisMQClient) Pubscribe(channel string, data []byte) (err error) {

	log.Println("pub msg", data)
	//redisChannel := "redChatRoom"
	conn := c.Pool.Get()

	defer conn.Close()

	_, err = c.Do("PUBLISH", channel, data)
	if err != nil {
		log.Println("pub err:", err)
	}
	return err

}
