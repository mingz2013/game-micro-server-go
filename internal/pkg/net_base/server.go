package net_base

type Server interface {
	StartServer()
	CloseServer()
	SetHandler(handler Handler)
}

const (
	PROTO_TCP = "TCP"
	PROTO_WS  = "WS"
)
