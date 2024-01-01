package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/selection"
)

func main() {
	fmt.Println("First:")
	firstResult, err := selection.Selection(selection.Settings{
		Options: []string{"Option 1", "Option 2", "Option 3"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(firstResult)

	fmt.Println("Second:")
	secondResult, err := selection.Selection(selection.Settings{
		Options:     []string{"Option 1", "Option 2", "Option 3"},
		DontCleanup: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(secondResult)

}
