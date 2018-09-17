package tcp

import "github.com/mingz2013/lib-go/net"

type connReader struct {
	conn *net.Conn
}

func (cr *connReader) Read(p []byte) (n int, err error) {
	return
}
