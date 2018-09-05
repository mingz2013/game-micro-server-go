package net

type Conn interface {
	WriteResponse(resp Response)
	GetExtra() int
	SetExtra(int)
	Serve() error
}
