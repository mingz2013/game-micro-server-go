package robot

import (
	"github.com/mingz2013/lib-go/msg"
	"log"
	"time"
)

type Robot struct {
	UserId int
	Name   string

	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg
}

func (r *Robot) Init() {

}

func NewRobot(userId int) Robot {
	r := Robot{UserId: userId}
	r.Init()
	return r
}

func (r Robot) Run() {
	log.Println("Robot.Run...")
	for {
		select {
		case m, ok := <-r.MsgIn:
			{
				if !ok {
					continue
				}
				r.onMsg(m)
			}
		case <-time.After(time.Second * 1):
			continue

		}
	}
}

func (r *Robot) onMsg(m msg.Msg) {
	// TODO 如果有人加入，随机说hello
	// TODO 如果有人离开，说拜拜
	// TODO 如果有人Hi，回复Hello
	// TODO
}

func (r *Robot) createRoom() {

}
