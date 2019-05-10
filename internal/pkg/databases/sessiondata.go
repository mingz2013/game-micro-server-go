package databases

import "strconv"

// userdata, 由于数据量太大，不好查询，所以建立sessiondata，存储热数据cache，用户上线后，建立sessiondata，

func GetSessionDataKey(userId int) (key string) {
	key = "sessiondata:" + strconv.Itoa(userId)
	return
}
