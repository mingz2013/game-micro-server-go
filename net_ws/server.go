package net_ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Server struct {
	rwc     websocket.Conn
	Handler Handler
	Addr    string
}

func NewServer(address string) *Server {
	s := &Server{}
	s.Init(address)
	return s
}

func (s *Server) Init(address string) {
	s.Addr = address
}

func (s *Server) Start() {

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(s, w, r)
	})
	log.Println("Server.Start...Addr:", s.Addr)
	err := http.ListenAndServe(s.Addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func (s *Server) Close() {
	s.rwc.Close()
}

func (s *Server) SetHandler(handler Handler) {
	s.Handler = handler
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

func serveWs(server *Server, w http.ResponseWriter, r *http.Request) {
	rw, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conn := &Conn{server: server, rwc: rw}
	//conn.server.register <- client

	go conn.Serve()

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	//go client.writePump()
	//go client.readPump()
}
