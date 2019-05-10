package models

import "time"

type Session struct {
	SessionID string `json:"sessionId" bson:"sessionId"`
	//User       *User         `json:"-" bson:"user"`
	UserType   string    `json:"userType" bson:"userType"`
	NickName   string    `json:"nickName" bson:"nickName"`
	CreateTime time.Time `json:"-" bson:"createTime"`
	UpdateTime time.Time `json:"-" bson:"updateTime"`
	Expires    time.Time `json:"-" bson:"expires"`
	Locale     string    `json:"-" bson:"locale"` // default is zh_CN
}
