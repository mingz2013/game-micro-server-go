package net_base

type Handler interface {
	Serve(c Conn, buf []byte)
	OnConn(c Conn) error
	OnClose(c Conn) error
}
