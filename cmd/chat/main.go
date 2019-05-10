package main

import (
	"github.com/mingz2013/lib-go/base"
	"github.com/mingz2013/lib-go/msg"
	"github.com/mingz2013/lib-go/robot-manager"
	"github.com/mingz2013/lib-go/table-manager"
	"log"
	"runtime"
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

func main() {
	log.Println("main,,,NumCPU", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	StartLocalTest()

}
