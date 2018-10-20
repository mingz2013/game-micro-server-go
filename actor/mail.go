package actor

type Mail struct {
	from      string // channel
	to        string // channel
	message   []byte
	needBack  bool  // 是否需要回信
	timestamp int64 // 时间戳
	mark      int64 // 标记，用于回信的时候，标明回的哪一封信
}

//func Init(){
//	time.Now().Unix()
//}
