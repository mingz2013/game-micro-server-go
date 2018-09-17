package net

type Handler interface {
	Serve(c Conn, buf []byte)
}
