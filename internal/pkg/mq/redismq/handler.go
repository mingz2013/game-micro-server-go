package redismq

type Handler interface {
	OnMessage(channel string, data []byte)
	onSubscription(channel string, kind string, count int)
}
