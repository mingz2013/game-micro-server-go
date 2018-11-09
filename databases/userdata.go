package databases

import "strconv"

// user，每个userId存储一条信息，hash，

// 存储所有的user数据

func GetUserDataKey(userId int) (key string) {
	key = "userdata:" + strconv.Itoa(userId)
	return
}

func ParamsMapToList(datas map[string]interface{}) (lst []interface{}) {
	for k, v := range datas {
		lst = append(lst, k)
		lst = append(lst, v)
	}
	return
}

func GetAttrs(userId int, fields ...string) (reply interface{}, err error) {
	key := GetUserDataKey(userId)
	return RedisClient.Do("HMGET", key, fields)

}

func SetAttrs(userId int, datas map[string]interface{}) (reply interface{}, err error) {
	key := GetUserDataKey(userId)
	params := ParamsMapToList(datas)
	return RedisClient.Do("HMSET", key, params)

}

func DelAttr(userId int, field string) {
	key := GetUserDataKey(userId)
	RedisClient.Do("HDEL", key, field)
}

func IncrAttr(userId int, field string, value interface{}) {
	key := GetUserDataKey(userId)
	RedisClient.Do("HINCRBY", key, field, value)

}
