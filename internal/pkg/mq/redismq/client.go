package redismq

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mingz2013/game-micro-server-go/internal/pkg/db/myredis"
	"log"
	"sync"
)

type RedisMQClient struct {
	myredis.RedisClient

	wgSub sync.WaitGroup

	channelMap map[string]redis.PubSubConn
	mu         sync.Mutex

	handler Handler
}

func (c *RedisMQClient) Wait() {
	c.wgSub.Wait()
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
	log.Println("Subscribe, get conn from pool", conn)
	psc := redis.PubSubConn{conn}
	psc.Subscribe(channel)
	c.mu.Lock()
	c.channelMap[channel] = psc
	c.mu.Unlock()

	c.wgSub.Add(1)
	go func() {
		defer func() {
			log.Println("Subscribe close..., to Unsubscribe...")
			//conn.Close()
			c.Unsubscribe(channel)
			//psc.Close()
			c.wgSub.Done()
		}()

		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				log.Println("messages<"+v.Channel+">:", v.Data)
				onMessage(v.Channel, v.Data)
			case redis.Subscription:
				log.Println("Subscription<"+v.Channel+">:", v.Kind, v.Count)

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
				log.Println("error:", v)
				return
				//continue

			}
		}

	}()
}

func (c *RedisMQClient) Unsubscribe(channel string) {
	log.Println("Unsubscribe...", channel)
	c.mu.Lock()
	conn, ok := c.channelMap[channel]
	if !ok {
		return
	}
	conn.Unsubscribe(channel)
	conn.Close()
	delete(c.channelMap, channel)
	c.mu.Unlock()
}

func (c *RedisMQClient) Pubscribe(channel string, data []byte) (err error) {

	log.Println("Pubscribe msg", "channel", channel, "<-data", data)
	//redisChannel := "redChatRoom"
	conn := c.Pool.Get()
	log.Println("Pubscribe, get conn from pool", conn)

	//psc := redis.PubSubConn{conn}

	defer func() {
		log.Println("Pubscribe:close...")
		conn.Close()
	}()

	//psc.

	_, err = conn.Do("PUBLISH", channel, data)
	if err != nil {
		log.Println("pub err:", err)
	}
	return err

}
