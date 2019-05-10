package net_ws

import (
	"github.com/gorilla/websocket"
	"github.com/mingz2013/lib-go/internal/pkg/net_base"
	"log"
	"net/http"
)

type Server struct {
	rwc     websocket.Conn
	handler net_base.Handler
	Addr    string
}

func NewServer(addr string) *Server {
	s := &Server{}
	s.Init(addr)
	return s
}

func (s *Server) Init(address string) {
	s.Addr = address
}

func (s *Server) StartServer() {

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(s.handler, w, r)
	})
	log.Println("Server.Start...Addr:", s.Addr)
	err := http.ListenAndServe(s.Addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func (s *Server) CloseServer() {
	s.rwc.Close()
}

func (s *Server) SetHandler(handler net_base.Handler) {
	s.handler = handler
}

//var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs(handler net_base.Handler, w http.ResponseWriter, r *http.Request) {
	rw, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conn := &Conn{handler: handler, rwc: rw}
	//conn.server.register <- client

	go conn.Serve()

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	//go client.writePump()
	//go client.readPump()
}
