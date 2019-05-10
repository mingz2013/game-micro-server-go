package main

import (
	"github.com/mingz2013/game-micro-server-go/internal/app/chat"
	"log"
	"runtime"
)

func main() {
	log.Println("main,,,NumCPU", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	chat.StartLocalTest()

}
