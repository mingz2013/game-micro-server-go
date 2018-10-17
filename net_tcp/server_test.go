package net_tcp

import (
	"log"
	"testing"
)

type ServerHandler struct {
	Handler
}

func (h *ServerHandler) Serve(c *Conn, buf []byte) {
	log.Println("on serve...")
	s := string(buf)
	log.Println("receive msg:", s)
	c.WriteString(s)
}

func (h *ServerHandler) OnConn(c *Conn) (err error) {
	log.Println("on conn...")
	c.WriteString("hello")
	return
}

func (h *ServerHandler) OnClose(c *Conn) (err error) {
	return
}

func NewHandler() Handler {
	h := &ServerHandler{}
	return h
}

func TestNewServer(t *testing.T) {
	s := NewServer("localhost:8080")
	s.SetHandler(NewHandler())
	s.Start()
}
