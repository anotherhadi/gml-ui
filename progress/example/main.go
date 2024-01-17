package main

import (
	"fmt"
	"time"

	"github.com/anotherhadi/gml-ui/progress"
)

func main() {
	percentageChan := make(chan int)

	go progress.ProgressBar(percentageChan)

	for i := 0; i < 10; i++ {
		percentageChan <- (i + 1) * 10
		time.Sleep(time.Second)
	}

	close(percentageChan)
	fmt.Println("Work done!")
}
