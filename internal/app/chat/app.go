package chat

import (
	"github.com/mingz2013/game-micro-server-go/internal/pkg/base"
	"github.com/mingz2013/game-micro-server-go/internal/pkg/msg"
	"github.com/mingz2013/game-micro-server-go/internal/pkg/robot-manager"
	"github.com/mingz2013/game-micro-server-go/internal/pkg/table-manager"
	"sync"
)

func StartLocalTest() {
	var wg sync.WaitGroup

	msgIn := make(chan msg.Msg)
	msgOut := make(chan msg.Msg)

	robotManager := robot_manager.NewRobotManager(msgIn, msgOut)
	roomManager := table_manager.NewTableManager("", msgOut, msgIn)

	base.RunProcessor(&wg, roomManager)
	base.RunProcessor(&wg, robotManager)

	wg.Wait()
}
