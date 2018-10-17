package net_tcp

import "bufio"
import (
	"log"
	"net"
)

//type Conn interface {
//	WriteBuffer(buf []byte)
//	GetExtra() int
//	SetExtra(int)
//	Serve() error
//}

type Conn struct {
	handler Handler

	removeAddr string

	r    *connReader
	bufr *bufio.Reader
	bufw *bufio.Writer

	rwc net.Conn

	extra int // 一个额外的指针类型的字段字段，可用于设置一些数据，做反向绑定
}

func NewConn() *Conn {
	c := &Conn{}
	return c
}

func (c *Conn) Connect() (err error) {
	c.rwc, err = net.Dial("tcp", c.removeAddr)
	if err != nil {
		return err
	}

	c.r = &connReader{conn: c}
	c.bufr = bufio.NewReader(c.r)
	c.bufw = bufio.NewWriterSize(c.rwc, 4<<10)

	return nil
}

func (c Conn) Serve() error {

	log.Println("new conn serve...")

	for {
		buffer, err := c.readBuffer()

		if err != nil {
			return err
		}

		c.handler.Serve(c, buffer)

	}

}

func (c *Conn) readBuffer() (b []byte, e error) {
	//n,e:=c.rwc.Read(b)
	//if
	c.bufr.Read(b)

	return
}

func (c Conn) WriteBuffer(buf []byte) {

}

//func (c *Conn) readRequest() (net2.Request, error) {
//	return net2.Request{}, nil
//}
//
//func (c Conn) WriteResponse(resp net2.Response) {
//
//}

func (c Conn) GetExtra() int {
	return c.extra
}

func (c Conn) SetExtra(e int) {
	c.extra = e
}
