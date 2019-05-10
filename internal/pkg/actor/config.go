package actor

import "github.com/mingz2013/game-micro-server-go/internal/pkg/net_base"

type Config struct {
	host     string
	port     string
	protocol string
}

func NewConfig() *Config {
	c := &Config{}
	c.Init()
	return c
}

func (c *Config) Init() {
	c.host = "localhost"
	c.port = "8000"
	c.protocol = net_base.PROTO_TCP
	//c.protocol = net_base.PROTO_WS
}

func (c *Config) ParseFromStr(conf string) {

}
