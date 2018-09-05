package tcp

import (
	net2 "github.com/mingz2013/lib-go/net"
	"log"
	"net"
)

type Server struct {
	Addr     string
	listener net.Listener
	Handler  net2.Handler
}

func NewServer(address string) *Server {
	s := &Server{}
	s.Init(address)
	return s
}

func (s *Server) Init(address string) {
	s.Addr = address
}

func (s *Server) SetHandler(handler net2.Handler) {
	s.Handler = handler
}

func (s *Server) Start() {
	log.Println("Server.Start...")
	s.ListenAndServe()

}

func (s *Server) Close() {
	s.listener.Close()
	s.listener = nil

}

func (s *Server) ListenAndServe() error {
	addr := s.Addr
	if addr == "" {
		addr = "localhost:8000"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.Serve(ln)
}

func (s *Server) Serve(l net.Listener) error {
	defer l.Close()

	for {
		rw, e := l.Accept()
		if e != nil {
			return e
		}

		c := s.newConn(rw)
		go c.Serve()

	}

}

func (s *Server) newConn(rwc net.Conn) *Conn {
	c := &Conn{
		server: s,
		rwc:    rwc,
	}
	return c
}
