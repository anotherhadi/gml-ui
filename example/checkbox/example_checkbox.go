package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/checkbox"
)

func main() {
	fmt.Println("First:")
	firstResult, err := checkbox.Checkbox(checkbox.Settings{
		Options: []string{"Option 1", "Option 2", "Option 3"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(firstResult)

	fmt.Println("Second:")
	secondResult, err := checkbox.Checkbox(checkbox.Settings{
		Options:        []string{"Option 1", "Option 2", "Option 3"},
		DefaultOptions: []bool{false, false, true},
		DontCleanup:    true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(secondResult)

}
