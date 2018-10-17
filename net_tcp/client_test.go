package net_tcp

import (
	"log"
	"testing"
)

type ClientHandler struct {
	Handler
}

func (h *ClientHandler) Serve(c *Conn, buf []byte) {
	log.Println("on serve...")
	s := string(buf)
	log.Println("receive msg:", s)

}

func (h *ClientHandler) OnConn(c *Conn) (err error) {
	log.Println("on conn...")
	c.WriteString("hello")
	return
}

func (h *ClientHandler) OnClose(c *Conn) (err error) {
	return
}

func NewClientHandler() Handler {
	h := &ClientHandler{}
	return h
}

func TestNewClient(t *testing.T) {
	c := NewClient()
	c.SetHandler(NewClientHandler())
	c.Connect("localhost:8080")

}
