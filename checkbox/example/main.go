package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/checkbox"
	"github.com/anotherhadi/gml-ui/settings"
)

func main() {
	checked, err := checkbox.Checkbox([]string{"Option 1", "Option 2", "Option 3", "Option 4", "Option 5"})
	if err != nil {
		panic(err)
	}
	fmt.Println(checked)

	// With custom settings
	checked, err = checkbox.Checkbox(
		[]string{"Option 1", "Option 2", "Option 3", "Option 4", "Option 5"},
		settings.Settings{
			TopPadding:     1,
			BottomPadding:  1,
			LeftPadding:    3,
			MaxCols:        80,
			MaxRows:        6,
			DefaultChecked: []bool{true, false, false, true, false},
			DontLoop:       true,
			Maximum:        3,
			Minimum:        1,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(checked)
}
