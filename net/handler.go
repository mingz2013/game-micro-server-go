package net

type Handler interface {
	Serve(c Conn, r Request)
}
