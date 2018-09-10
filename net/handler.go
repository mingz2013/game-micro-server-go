package net

import "bytes"

type Handler interface {
	Serve(c Conn, buffer bytes.Buffer)
}
