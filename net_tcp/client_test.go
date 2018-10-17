package net_tcp

import "testing"

type ClientHandler struct {
	Handler
}

func (h *ClientHandler) Serve(c Conn, buf []byte) {

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
