package net_tcp

import "testing"

type ServerHandler struct {
	Handler
}

func (h *ServerHandler) Serve(c Conn, buf []byte) {

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
