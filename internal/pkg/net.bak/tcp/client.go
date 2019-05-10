package tcp

import "net"

func (c *Conn) Connect() (err error) {
	c.rwc, err = net.Dial("tcp", c.removeAddr)
	if err != nil {
		return err
	}
	return nil
}

//func (c *Conn) ReadResponse() {
//
//}
