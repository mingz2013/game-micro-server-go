package actor

import (
	"github.com/mingz2013/lib-go/internal/pkg/net_base"
	"log"
	"testing"
)

type ActorHandler struct {
	net_base.Handler
}

func (h *ActorHandler) Serve(c net_base.Conn, buf []byte) {
	log.Println("on serve...", c, buf)
	s := string(buf)
	log.Println("receive msg:", s)
	c.WriteString(s)
}

func (h *ActorHandler) OnConn(c net_base.Conn) (err error) {
	log.Println("on conn...")
	c.WriteString("hello")
	return
}

func (h *ActorHandler) OnClose(c net_base.Conn) (err error) {
	return
}

func NewHandler() *ActorHandler {
	h := &ActorHandler{}
	return h
}

func TestNewActor(t *testing.T) {
	a := NewConnectorActor("")
	a.SetHandler(NewHandler())
	a.Start()

}
