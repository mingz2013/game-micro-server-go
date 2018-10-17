package net_tcp

import (
	"bufio"
	"log"
	"net"
)

type Server struct {
	Addr     string
	listener net.Listener
	handler  Handler
}

func NewServer(address string) *Server {
	s := &Server{}
	s.Init(address)
	return s
}

func (s *Server) Init(address string) {
	s.Addr = address
}

func (s *Server) SetHandler(handler Handler) {
	s.handler = handler
}

func (s *Server) GetHandler() Handler {
	return s.handler
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

	log.Println("Serve on", s.Addr)

	for {
		rw, e := l.Accept()
		if e != nil {
			return e
		}

		c := s.newConn(rw)

		//s.Handler.OnConn(c)

		go c.Serve()

	}

}

func (s *Server) newConn(rwc net.Conn) *Conn {
	c := &Conn{
		handler: s.handler,
		rwc:     rwc,
	}

	c.r = &connReader{conn: c}
	c.bufr = bufio.NewReader(c.r)
	c.bufw = bufio.NewWriterSize(c.rwc, 4<<10)

	return c
}
