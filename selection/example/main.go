package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/selection"
	"github.com/anotherhadi/gml-ui/settings"
)

func main() {
	selected, err := selection.Selection([]string{"Option 1", "Option 2", "Option 3", "Option 4", "Another option", "Last option"})
	if err != nil {
		panic(err)
	}
	fmt.Println(selected)

	// With settings and a filter
	selected, err = selection.Selection(
		[]string{"Option 1", "Option 2", "Option 3", "Option 4", "Another option", "Last option"},
		settings.Settings{
			Filter:   true,
			DontLoop: true,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(selected)
}
