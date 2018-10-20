package actor

import "time"

var (
	MARK_ID int64
)

type Mail struct {
	from      string // channel
	to        string // channel
	message   []byte
	needBack  bool  // 是否需要回信
	isBack    bool  // 是否是回信
	timestamp int64 // 时间戳
	mark      int64 // 标记，用于回信的时候，标明回的哪一封信
}

func Init() {
	//time.Now().Unix()
	MARK_ID = 0
}

func NewMail(from string, to string, message []byte, needBack bool, isBack bool) *Mail {
	MARK_ID += 1
	m := &Mail{
		from:      from,
		to:        to,
		message:   message,
		needBack:  needBack,
		isBack:    isBack,
		timestamp: time.Now().Unix(),
		mark:      MARK_ID,
	}

	return m
}
