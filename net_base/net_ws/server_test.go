package net_ws

import (
	"github.com/mingz2013/lib-go/net_base"
	"log"
	"testing"
)

type ServerHandler struct {
	net_base.Handler
}

func (h *ServerHandler) Serve(c *Conn, buf []byte) {
	log.Println("on serve...")
	s := string(buf)
	log.Println("receive msg:", s)
	c.WriteString(s)
}

func (h *ServerHandler) OnConn(c *Conn) (err error) {
	log.Println("on conn...")
	//c.WriteString("hello")
	return
}

func (h *ServerHandler) OnClose(c *Conn) (err error) {
	return
}

func NewHandler() net_base.Handler {
	h := &ServerHandler{}
	return h
}

func TestNewServer(t *testing.T) {
	s := NewServer("localhost:8002")
	s.SetHandler(NewHandler())
	s.StartServer()
}
