package net_tcp

type Handler interface {
	Serve(c Conn, buf []byte)
}

//func (h*Handler)Serve(c Conn, buf []byte){
//
//}
