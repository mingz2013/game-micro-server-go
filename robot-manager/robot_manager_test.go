package robot_manager

import (
	"github.com/mingz2013/lib-go/msg"
	"testing"
)

func TestNewRobotManager(t *testing.T) {

	MsgIn := make(chan msg.Msg)
	MsgOut := make(chan msg.Msg)

	m := NewRobotManager(MsgIn, MsgOut)

	m.Run()

}
