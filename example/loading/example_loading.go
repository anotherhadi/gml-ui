package main

import (
	"fmt"
	"time"

	"github.com/anotherhadi/gml-ui/loading"
)

func main() {

	loadingChan := make(chan bool)

	go loading.Loading(loadingChan, loading.Settings{
		Message:     "Are you waiting? Yes you are",
		DontCleanup: true,
	})

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
	}

	loadingChan <- false

	close(loadingChan)
	fmt.Println("Work done!")
}
