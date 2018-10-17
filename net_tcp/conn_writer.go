package net_tcp

type connWriter struct {
	conn *Conn
}

func (cr *connWriter) Write(p []byte) (n int, err error) {
	return
}
