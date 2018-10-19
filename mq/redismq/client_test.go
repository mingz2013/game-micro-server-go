package redismq

import (
	"encoding/json"
	"log"
	"testing"
)

func TestNewRedisMQClient(t *testing.T) {

	conf := map[string]interface{}{
		"host": "localhost",
		"port": "6379",
		"db":   0,
	}

	confByte, _ := json.Marshal(conf)
	confStr := string(confByte)

	log.Println("confStr", confStr)

	client := NewRedisMQClient(confStr)

	client.Subscribe("test-channel", func(channel string, data []byte) {

		log.Println("channel", channel, "data", data)
		client.Unsubscribe("test-channel")

	}, func(channel string, kind string, count int) {
		log.Println("channel", channel, "kind", kind, "count", count)
	})

	client.Pubscribe("test-channel", []byte("hello"))

	ch := make(chan int)
	<-ch

}
