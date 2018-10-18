package net_tcp

import (
	"github.com/mingz2013/lib-go/net_base"
	"log"
	"testing"
)

type ClientHandler struct {
	net_base.Handler
}

func (h *ClientHandler) Serve(c net_base.Conn, buf []byte) {
	log.Println("on serve...")
	s := string(buf)
	log.Println("receive msg:", s)

}

func (h *ClientHandler) OnConn(c net_base.Conn) (err error) {
	log.Println("on conn...")
	c.WriteString("hello")
	return
}

func (h *ClientHandler) OnClose(c net_base.Conn) (err error) {
	return
}

func NewClientHandler() net_base.Handler {
	h := &ClientHandler{}
	return h
}

func TestNewClient(t *testing.T) {
	c := NewClient()
	c.SetHandler(NewClientHandler())
	c.Connect("localhost:8000")

}
