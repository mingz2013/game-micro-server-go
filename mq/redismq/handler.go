package redismq

type Handler interface {
	OnMessage(channel string, data []byte)
}
