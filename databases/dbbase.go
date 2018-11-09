package databases

import "github.com/mingz2013/lib-go/db/myredis"

var (
	RedisClient *myredis.RedisClient
)

func Init() {
	RedisClient = myredis.NewRedisClient("")

}
