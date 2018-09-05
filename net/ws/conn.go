package ws

import (
	"bufio"
	"github.com/gorilla/websocket"
	"github.com/mingz2013/lib-go/net"
)

type Conn struct {
	server *Server

	removeAddr string

	r    *connReader
	bufr *bufio.Reader
	bufw *bufio.Writer

	rwc *websocket.Conn

	extra int // 一个额外的指针类型的字段字段，可用于设置一些数据，做反向绑定
}

func (c Conn) Serve() error {

	for {
		r, err := c.readRequest()

		if err != nil {
			return err
		}

		c.server.Handler.Serve(c, r)

	}

	return nil

}

func (c *Conn) readRequest() (net.Request, error) {
	c.rwc.ReadMessage()
	return net.Request{}, nil
}

func (c Conn) WriteResponse(resp net.Response) {
	//c.rwc.WriteMessage()
}

func (c Conn) GetExtra() int {
	return c.extra
}

func (c Conn) SetExtra(e int) {
	c.extra = e
}
