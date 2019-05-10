package net_tcp

import (
	"encoding/binary"
	"encoding/json"
	"github.com/mingz2013/lib-go/internal/pkg/net_base"
	"io"
	"log"
	"net"
)

//type Conn interface {
//	WriteBuffer(buf []byte)
//	GetExtra() int
//	SetExtra(int)
//	Serve() error
//}

const (
	headSize = 2
)

type Conn struct {
	handler net_base.Handler

	removeAddr string

	//r    *connReader
	//bufr *bufio.Reader
	//bufw *bufio.Writer

	rwc net.Conn

	extra int // 一个额外的指针类型的字段字段，可用于设置一些数据，做反向绑定
}

func NewConn() Conn {
	c := Conn{}
	//c.r = &connReader{conn: c}
	//c.bufr = bufio.NewReaderSize(c.rwc, 4<<10)
	//c.bufw = bufio.NewWriterSize(c.rwc, 4<<10)
	log.Println("NewConn...", c)
	return c
}

func (c *Conn) Connect() (err error) {
	c.rwc, err = net.Dial("tcp", c.removeAddr)
	if err != nil {
		return err
	}

	return nil
}

func (c *Conn) Close() {
	c.rwc.Close()
}

func (c *Conn) Serve() (err error) {

	log.Println("new conn serve...")

	for {
		buffer, err := c.Read()

		if err != nil {
			return err
		}

		log.Println("Conn.Serve", "c", c, &c, "buffer", buffer, &buffer)
		c.handler.Serve(c, buffer)

	}

	return
}

func (c *Conn) Read() (b []byte, e error) {

	size := make([]byte, headSize)

	if _, err := io.ReadFull(c.rwc, size); err != nil {
		return nil, err
	}

	//hb, err := c.read(headSize)
	//if err != nil{
	//	e = err
	//	return
	//}

	//var n int32
	//bf := bytes.NewBuffer(hb)
	//binary.Read(bf, binary.BigEndian, &n)

	n := binary.BigEndian.Uint16(size)

	message := make([]byte, n)
	copy(message, size)

	if _, err := io.ReadFull(c.rwc, message[headSize:]); err != nil {
		return nil, err
	}
	return message[headSize:], nil
	//b, e = c.read(int(n))

	//return
}

func (c *Conn) Write(buf []byte) error {
	//n := len(buf)
	//
	//
	//tmp:= int32(n)
	//bytesBuffer:= bytes.NewBuffer([]byte{})
	//binary.Write(bytesBuffer, binary.BigEndian, tmp)
	//binary.Write(bytesBuffer, binary.BigEndian, buf)
	//
	//c.bufw.Write(bytesBuffer.Bytes())

	size := len(buf) + headSize
	message := make([]byte, size)
	binary.BigEndian.PutUint16(message, uint16(size))
	//binary.BigEndian.PutUint16(message[2:], mcmd)
	//binary.BigEndian.PutUint16(message[4:], scmd)
	copy(message[headSize:], buf)

	if _, err := c.rwc.Write(message); err != nil {
		return err
	}

	return nil

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
