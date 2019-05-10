package tcp

import "github.com/mingz2013/lib-go/internal/pkg/net.bak"

type connReader struct {
	conn *net_bak.Conn
}

func (cr *connReader) Read(p []byte) (n int, err error) {
	return
}
