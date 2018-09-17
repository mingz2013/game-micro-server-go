package net

type Conn interface {
	WriteBuffer(buf []byte)
	GetExtra() int
	SetExtra(int)
	Serve() error
}
