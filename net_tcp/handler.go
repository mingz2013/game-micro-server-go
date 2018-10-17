package net_tcp

type Handler interface {
	Serve(c *Conn, buf []byte)
	OnConn(c *Conn) error
	OnClose(c *Conn) error
}

//func (h*Handler)Serve(c Conn, buf []byte){
//
//}
