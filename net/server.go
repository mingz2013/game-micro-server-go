package net

type Server interface {
	Close()
	Start()
	SetHandler(handler Handler)
}

const (
	PROTO_TCP = "TCP"
	PROTO_WS  = "WS"
)
