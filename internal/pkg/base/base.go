package base

import (
	"log"
	"sync"
)

type Processor interface {
	Run()
}

func RunProcessor(wg *sync.WaitGroup, p Processor) {
	log.Println("RunProcessor...")
	wg.Add(1)

	go func() {
		log.Println("RunProcessor.func...")
		defer wg.Done()

		p.Run()

	}()
}
