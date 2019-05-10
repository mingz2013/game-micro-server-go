package net_base

type Conn interface {
	Write(buf []byte) error
	WriteString(s string) (err error)
	WriteJson(js interface{}) (err error)
	GetExtra() int
	SetExtra(int)
}
