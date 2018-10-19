package myredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
	"testing"
)

func TestRedisClient(t *testing.T) {

	conf := map[string]interface{}{
		"host": "localhost",
		"port": "6379",
		"db":   0,
	}

	confByte, _ := json.Marshal(conf)
	confStr := string(confByte)

	log.Println("confStr", confStr)

	client := NewRedisClient(confStr)

	client.Do("SET", "test-key", "test-value")
	reply, err := client.String(client.Do("GET", "test-key"))
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("value:...", reply)

	// 存入json数据
	key := "test-key-2"
	imap := map[string]string{"key1": "111", "key2": "222"}
	// 将map转换成json数据
	value1, _ := json.Marshal(imap)
	// 存入redis
	n, err := client.Do("SETNX", key, value1)
	if err != nil {
		log.Println(err)
	}
	if n == int64(1) {
		log.Println("setnx key success", key, value1)
	}

	// 取json数据
	// 先声明imap用来装数据
	var imap1 map[string]string
	// json数据在go中是[]byte类型，所以此处用redis.Bytes转换
	value2, err := redis.Bytes(client.Do("GET", key))
	if err != nil {
		log.Println(err)
	}

	// 将json解析成map类型
	errShal := json.Unmarshal(value2, &imap1)
	if errShal != nil {
		log.Println(errShal)

	}

	log.Println("get imap1", imap1)

	ch := make(chan int)
	<-ch

}
