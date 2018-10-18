package net_bak

type Handler interface {
	Serve(c Conn, buf []byte)
}
