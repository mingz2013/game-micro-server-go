package net_ws

import (
	"github.com/gorilla/websocket"
	"log"
)

type Conn struct {
	server *Server

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

		c.server.Handler.Serve(c, b)

	}

	return nil

}

func (c *Conn) Read() (buf []byte, err error) {
	t, buf, err := c.rwc.ReadMessage()
	log.Println(t)
	return
}

func (c *Conn) Write(buf []byte) {
	c.rwc.WriteMessage(1, buf)
}

func (c *Conn) GetExtra() int {
	return c.extra
}

func (c *Conn) SetExtra(e int) {
	c.extra = e
}
