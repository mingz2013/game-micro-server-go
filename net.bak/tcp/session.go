package tcp

import (
	"net"
)

type Session struct {
	conn       net.Conn
	recvBuffer []byte
	sendBuffer []byte
}

func NewSession(conn net.Conn) *Session {
	s := &Session{}
	s.Init(conn)
	return s
}

func (s *Session) Init(conn net.Conn) {
	s.conn = conn
}

func (s *Session) Start() {

}

func (s *Session) Write(b []byte) {
	s.conn.Write(b)
}

func (s *Session) Read() {

}

func (s *Session) SendMsg() {

}

func (s *Session) onMsg() {

}

func (s *Session) sendloop() {

}

func (s *Session) recvloop() {
	//for {

	//readnum, err := io.ReadAtLeast(s.conn, s.recvBuffer, 1024)

	//}
}

func (s *Session) RemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}
