package db

import "github.com/gomodule/redigo/redis"

var pool *redis.Pool

func GetRedisPool() *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:   3, /*最大的空闲连接数*/
		MaxActive: 8, /*最大的激活连接数*/
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "链接地址，例如127.0.0.1:6379", redis.DialPassword("密码"))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	return pool
}

func DoConn() {
	c := pool.Get()
	defer c.Close()

	////存值,
	//_, err := c.Do("SET", "key", "value")
	////设置过期时间
	//_, err := c.Do("SET", "key", "value"，"EX",360)
	////存int
	//_, err := c.Do("SET", "key", 2)
	//
	////取值
	//v,err:=redis.String(c.Do("GET","key"))
	//bytes, err := redis.Bytes(c.Do("GET", "key"))

}
