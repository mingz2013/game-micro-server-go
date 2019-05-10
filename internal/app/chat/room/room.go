package room

import (
	"github.com/mingz2013/lib-go/msg"
	"log"
	"time"
)

type Room struct {
	RoomId int
	Name   string

	RoomPlayers

	Creator Player

	// 定义好所有的输入输出接口，就可以定义这个类内部的功能了
	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg
}

func NewRoom(roomId int, MsgIn <-chan msg.Msg, MsgOut chan<- msg.Msg) Room {
	return Room{RoomId: roomId, Name: "", MsgIn: MsgIn, MsgOut: MsgOut}
}

func (r *Room) Init() {
	r.RoomPlayers.Init()
}

func (r *Room) Close() {
	close(r.MsgOut)
}

func (r Room) Run() {
	log.Println("Room.Run...")
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

func (r *Room) onMsg(m msg.Msg) {
	cmd := m["cmd"]
	if cmd != "room" {
		return
	}

	params := m.GetParams()
	action := params["action"].(string)

	switch action {
	case "join":
		{
			r.onActionJoin(m)
		}
	case "leave":
		{
			r.onActionLeave(m)
		}
	case "chat":
		{
			r.onActionChat(m)
		}
	default:
		log.Println("error action", m)

	}

}

func (r *Room) onActionJoin(m msg.Msg) {
	userId := m["userId"].(int)
	p, ok := r.AddPlayer(userId)
	if !ok {
		log.Println("already in...")
		return
	}

	log.Println("join ok...")

	m.SetResults(map[string]interface{}{
		"action":  "join",
		"retcode": 0,
		"msg":     "join ok",
		"userId":  p.UserId,
		"name":    p.Name,
	})

	r.Broadcast(m)

}

func (r *Room) onActionLeave(m msg.Msg) {
	userId := m["userId"].(int)
	p, ok := r.RmPlayer(userId)
	if !ok {
		log.Println("not in...")
		return
	}
	log.Println("rm ok...")

	m.SetResults(map[string]interface{}{
		"action":  "leave",
		"retcode": 0,
		"msg":     "leave ok",
		"userId":  p.UserId,
		"name":    p.Name,
	})

	r.Broadcast(m)
}

func (r *Room) onActionChat(m msg.Msg) {
	userId := m["userId"].(int)
	msgStr := m["msg"].(string)
	//m := msg.Msg{"from": userId}

	p, ok := r.FindPlayerByUserId(userId)
	if !ok {
		log.Println("player not in room", m)
		return
	}

	m.SetResults(map[string]interface{}{
		"action":  "chat",
		"retcode": 0,
		"msg":     msgStr,
		"userId":  p.UserId,
		"name":    p.Name,
	})

	r.Broadcast(m)

}

func (r *Room) SendRes(userId int, m msg.Msg) {
	m.SetKey("userId", userId)
	r.MsgOut <- m
}

func (r *Room) Broadcast(m msg.Msg) {
	for i := 0; i < len(r.players); i++ {
		r.SendRes(r.players[i].UserId, m)
	}
}
