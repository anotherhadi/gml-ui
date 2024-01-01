package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/input"
)

func main() {
	result, err := input.Input(input.Settings{
		Prompt:      "Do you like this library?",
		DontCleanup: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Input:", result)
}
