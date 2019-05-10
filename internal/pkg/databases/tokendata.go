package databases

import "strconv"

// 用户的token数据

func GetUserId2TokenKey(userId int) (key string) {
	return "userid2token:" + strconv.Itoa(userId)
}

func GetToken2UserIdKey(token string) (key string) {
	return "token2userid:" + token
}

func GetTokenByUserId(userId int) (token string) {
	return
}

func GetUserIdByToken(token string) (userId int) {
	return
}
