package net_tcp

type connReader struct {
	conn *Conn
}

func (cr *connReader) Read(p []byte) (n int, err error) {
	return
}
