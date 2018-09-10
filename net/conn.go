package net

import "bytes"

type Conn interface {
	WriteBuffer(buffer bytes.Buffer)
	GetExtra() int
	SetExtra(int)
	Serve() error
}
