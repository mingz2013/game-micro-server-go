package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
	"time"
)

type RedisClient struct {
	pool  *redis.Pool
	host  string
	port  string
	db    int // select
	wgSub sync.WaitGroup
}

func NewRedisClient(conf string) *RedisClient {
	client := &RedisClient{}

	client.Init(conf)
	return client

}

func (c *RedisClient) Init(conf string) {
	log.Println(conf)
	var confJs map[string]interface{}
	err := json.Unmarshal([]byte(conf), &confJs)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(confJs)

	c.host = confJs["host"].(string)
	c.port = confJs["port"].(string)
	c.db = int(confJs["db"].(float64))

	log.Println(c.host, c.port, c.db)

	c.pool = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   1024,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {

			conn, err := redis.Dial("tcp", c.host+":"+c.port)
			if err != nil {
				return nil, err
			}

			log.Println("conn success...")

			// select db

			conn.Do("SELECT", c.db)

			return conn, nil
		},
	}

}

func (c *RedisClient) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := c.pool.Get()
	defer conn.Close()

	return conn.Do(commandName, args...)
}

func (c *RedisClient) Subscribe(channel string, onMessage func(channel string, data []byte), onSubscription func(channel string, kind string, count int)) {
	//redisChannel := "redChatRoom"
	conn := c.pool.Get()
	psc := redis.PubSubConn{conn}
	psc.Subscribe(channel)

	c.wgSub.Add(1)
	go func() {
		defer func() {
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
				onSubscription(v.Channel, v.Kind, v.Count)
				continue
			case error:
				log.Println(v)
				return

			}
		}

	}()
}

func (c *RedisClient) Pubscribe(channel string, data []byte) (err error) {

	log.Println("pub msg", data)
	//redisChannel := "redChatRoom"
	conn := c.pool.Get()

	defer conn.Close()

	_, err = c.Do("PUBLISH", channel, data)
	if err != nil {
		log.Println("pub err:", err)
	}
	return err

}

func (c *RedisClient) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

func (c *RedisClient) Int(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}
