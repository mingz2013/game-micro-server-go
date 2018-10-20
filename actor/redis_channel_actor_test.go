package actor

import (
	"encoding/json"
	"testing"
)

func TestNewRedisChannelActor(t *testing.T) {

	conf := map[string]interface{}{
		"host": "localhost",
		"port": "6379",
		"db":   0,
	}

	data, _ := json.Marshal(conf)
	dataStr := string(data)

	a := NewRedisChannelActor(dataStr)

	message := "hello , my name is actor redis"

	mail := NewMail(a.channel, a.channel, []byte(message), true, false)

	a.SendMailNeedBack(*mail)

}
