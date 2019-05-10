package actor

type RedisChannelActorHandler interface {
	OnRedisChannelMessage(message []byte) (retMsg []byte)
}
