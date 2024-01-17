package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/input"
)

func main() {
	userInput, err := input.Input("What's your name?")
	if err != nil {
		panic(err)
	}

	fmt.Println(userInput)
}
