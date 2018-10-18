package actor

import "github.com/mingz2013/lib-go/net_base"

type Config struct {
	address  string
	protocol string
}

func NewConfig() *Config {
	c := &Config{}
	c.Init()
	return c
}

func (c *Config) Init() {
	c.address = "localhost:8000"
	c.protocol = net_base.PROTO_TCP
	c.protocol = net_base.PROTO_WS
}

func (c *Config) ParseFromStr(conf string) {

}
