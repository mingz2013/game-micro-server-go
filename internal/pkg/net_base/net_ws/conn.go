package net_ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/mingz2013/lib-go/internal/pkg/net_base"
	"log"
)

type Conn struct {
	handler net_base.Handler

	removeAddr string

	//r    *connReader
	//bufr *bufio.Reader
	//bufw *bufio.Writer

	rwc *websocket.Conn

	extra int // 一个额外的指针类型的字段字段，可用于设置一些数据，做反向绑定
}

func (c *Conn) Serve() error {

	for {
		b, err := c.Read()

		if err != nil {
			return err
		}

		c.handler.Serve(c, b)

	}

	return nil

}

func (c *Conn) Read() (buf []byte, err error) {
	t, buf, err := c.rwc.ReadMessage()
	log.Println(t)
	return
}

func (c *Conn) Write(buf []byte) (err error) {
	c.rwc.WriteMessage(1, buf)
	return
}

func (c *Conn) WriteString(s string) (err error) {

	data := []byte(s)
	err = c.Write(data)

	return
}

func (c *Conn) WriteJson(js interface{}) (err error) {

	data, err := json.Marshal(js)
	if err != nil {
		return err
	}
	err = c.Write(data)
	return
}

func (c *Conn) GetExtra() int {
	return c.extra
}

func (c *Conn) SetExtra(e int) {
	c.extra = e
}
