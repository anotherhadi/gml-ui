package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/confirm"
)

func main() {
	fmt.Println("First:")
	firstResult, err := confirm.Confirm()
	if err != nil {
		panic(err)
	}
	fmt.Println(firstResult)

	fmt.Println("Second:")
	secondResult, err := confirm.Confirm(confirm.Settings{
		Prompt:         "Do you like this library?",
		DefaultToFalse: true,
		DontCleanup:    true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(secondResult)
}
