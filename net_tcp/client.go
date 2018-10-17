package net_tcp

type Client struct {
	conn    Conn
	handler Handler
}

func (c *Client) Init() {
	c.conn = NewConn()
}

func (c *Client) SetHandler(handler Handler) {
	c.handler = handler
}

func NewClient() *Client {

	c := &Client{}
	c.Init()
	return c
}

func (c *Client) Connect(addr string) {
	c.conn.removeAddr = addr
	err := c.conn.Connect()
	if err == nil {
		c.conn.handler = c.handler
		c.conn.handler.OnConn(&c.conn)
		c.conn.Serve()
	}
}
