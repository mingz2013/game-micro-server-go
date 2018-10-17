package net_tcp

import (
	"log"
	"net"
	"sync"
)

type Server struct {
	Addr     string
	listener net.Listener
	handler  Handler
	conns    map[Conn]interface{}

	mutex     sync.Mutex
	waitgroup sync.WaitGroup
}

func (s *Server) addConn(conn Conn) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.conns == nil {
		return false
	}
	s.conns[conn] = nil
	log.Println("addConn success...")
	return true
}

func (s *Server) removeConn(conn Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.conns != nil {
		delete(s.conns, conn)
		conn.Close()
	}
}

func NewServer(address string) *Server {
	s := &Server{}
	s.Init(address)
	return s
}

func (s *Server) Init(address string) {
	s.Addr = address
	s.conns = make(map[Conn]interface{})
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
		log.Println("after new", c)
		if !s.addConn(c) {
			c.Close()
			continue
		}
		log.Println("before handleConn", c)
		go s.handleConn(c)
		log.Println("end for....", c)

	}

}

func (s *Server) newConn(rwc net.Conn) Conn {
	c := NewConn()
	c.rwc = rwc
	c.handler = s.handler
	log.Println("newConn...", c)
	return c
}

func (s *Server) handleConn(conn Conn) {
	log.Println("handleConn.....", conn)
	s.waitgroup.Add(1)
	defer s.waitgroup.Done()
	//if !s.addConn(conn) {
	//	conn.Close()
	//	return
	//}
	defer s.removeConn(conn)

	s.handler.OnConn(&conn)
	log.Println("handleConn...", conn)
	conn.Serve()

	s.handler.OnClose(&conn)

}
