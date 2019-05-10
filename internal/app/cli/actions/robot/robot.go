package robot

import (
	"log"
	"net/http"
)

func Start() error {
	resp, err := http.Get("http://localhost:8006/robot/start/")
	if err != nil {
		return err
	}

	log.Println("resp", resp)

	return nil
}
