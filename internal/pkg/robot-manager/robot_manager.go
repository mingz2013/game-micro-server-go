package robot_manager

import (
	"github.com/mingz2013/lib-go/internal/pkg/base"
	"github.com/mingz2013/lib-go/internal/pkg/msg"
	majiangrobot "github.com/mingz2013/mahjong-table-go/robot"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

type RobotSession struct {
	Robot  Robot
	MsgIn  chan msg.Msg
	MsgOut chan msg.Msg
}

type RobotManager struct {
	SessionMap map[int]RobotSession

	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg

	RobotIds []int
}

func (r *RobotManager) Init() {

	r.SessionMap = map[int]RobotSession{}

	r.RobotIds = []int{}
	for i := 1; i <= 1000; i++ {
		r.RobotIds = append(r.RobotIds, i)
	}

	for i := 0; i < 4; i++ {
		r.NewRobot()
	}

}

func NewRobotManager(msgIn chan msg.Msg, msgOut chan msg.Msg) *RobotManager {
	r := RobotManager{MsgIn: msgIn, MsgOut: msgOut}
	r.Init()
	return &r
}

func (r *RobotManager) NewTestRobots(num int) {
	log.Println("RobotManager.NewTestRobots", num)
	defer wg.Done()
	for i := 0; i < num; i++ {
		r.NewRobot()
	}
}

func (r RobotManager) Run() {
	log.Println("RobotManager.Run...")
	wg.Add(1)
	go r.NewTestRobots(4)

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

	wg.Wait()
}

func (r *RobotManager) onMsg(m msg.Msg) {
	log.Println("RobotManager.onMsg", m)

	userId := m.GetKey("id").(int)

	//session := r.SessionMap[userId]
	session := r.findRobotByUserId(userId)

	session.MsgIn <- m

}

func (r *RobotManager) findRobotByUserId(userId int) RobotSession {
	return r.SessionMap[userId]
}

func (r *RobotManager) NewRobot() {
	userId, ok := r.findOneUserId()
	if !ok {
		return
	}

	msgOut := make(chan msg.Msg)
	msgIn := make(chan msg.Msg)

	robot := majiangrobot.NewRobot(userId, "", msgIn, msgOut)

	session := RobotSession{}
	session.Robot = robot
	session.MsgIn = msgIn
	session.MsgOut = msgOut
	r.SessionMap[userId] = session

	base.RunProcessor(&wg, robot)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case m, ok := <-session.MsgOut:
				{
					if !ok {
						continue
					}
					r.MsgOut <- m
				}
			case <-time.After(time.Second * 1):
				continue

			}
		}
	}()

}

func (r *RobotManager) findOneUserId() (userId int, ok bool) {
	if len(r.RobotIds) == 0 {
		return 0, false
	}
	userId = r.RobotIds[0]
	r.RobotIds = r.RobotIds[1:]
	return userId, true
}
