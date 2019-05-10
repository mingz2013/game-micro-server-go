package databases

import "strconv"

// 用户游戏数据

func GetGameDataKey(gameId, userId int) (key string) {
	key = "gamedata:" + strconv.Itoa(gameId) + ":" + strconv.Itoa(userId)
	return
}
