package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/selection_filter"
)

func main() {
	fmt.Println("First:")
	firstResult, err := selection_filter.SelectionFilter(selection_filter.Settings{
		Options: []string{"Option 1", "Option 2", "Option 3"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(firstResult)

	fmt.Println("Second:")
	secondResult, err := selection_filter.SelectionFilter(selection_filter.Settings{
		Options:       []string{"Option 1", "Option 2", "Option 3", "option 4", "Another one", "Again", "The last one"},
		MaxRows:       6,
		DontCleanup:   true,
		CaseSensitive: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(secondResult)

}
