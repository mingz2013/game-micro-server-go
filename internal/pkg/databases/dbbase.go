package databases

import "github.com/mingz2013/game-micro-server-go/internal/pkg/db/myredis"

var (
	RedisClient *myredis.RedisClient
)

func Init() {
	RedisClient = myredis.NewRedisClient("")

}
