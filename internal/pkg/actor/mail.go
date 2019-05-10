package actor

import "time"

var (
	MARK_ID int64
)

type Mail struct {
	From      string `json:from` // channel
	To        string `json:to`   // channel
	Message   []byte `json:message`
	NeedBack  bool   `json:needback`  // 是否需要回信
	IsBack    bool   `json:isback`    // 是否是回信
	Timestamp int64  `json:timestamp` // 时间戳
	Mark      int64  `json:mark`      // 标记，用于回信的时候，标明回的哪一封信
}

func Init() {
	//time.Now().Unix()
	MARK_ID = 0
}

func NewMail(from string, to string, message []byte, needBack bool, isBack bool, mark int64) *Mail {

	m := &Mail{
		From:      from,
		To:        to,
		Message:   message,
		NeedBack:  needBack,
		IsBack:    isBack,
		Timestamp: time.Now().Unix(),
		Mark:      mark,
	}

	if isBack {

	} else {
		MARK_ID += 1
		m.Mark = MARK_ID
	}

	return m
}
