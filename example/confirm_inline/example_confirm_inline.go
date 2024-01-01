package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/confirm_inline"
)

func main() {
	fmt.Println("First:")
	firstResult, err := confirm_inline.ConfirmInline()
	if err != nil {
		panic(err)
	}
	fmt.Println(firstResult)

	fmt.Println("Second:")
	secondResult, err := confirm_inline.ConfirmInline(confirm_inline.Settings{
		Prompt:         "Do you like this library?",
		DefaultToFalse: true,
		DontCleanup:    true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(secondResult)
}
