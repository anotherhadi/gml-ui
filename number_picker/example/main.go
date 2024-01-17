package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/number_picker"
	"github.com/anotherhadi/gml-ui/settings"
)

func main() {

	// Int
	number, err := number_picker.NumberPicker()
	if err != nil {
		panic(err)
	}
	fmt.Println(number)

	// Float
	// With settings
	number, err = number_picker.NumberPicker(settings.Settings{
		Decimal:   2,
		Minimum:   -10,
		Maximum:   100,
		Increment: 0.1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(number)
}
