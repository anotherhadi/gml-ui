package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/number_picker"
)

func main() {
	fmt.Println("First:")
	firstResult, err := number_picker.NumberPicker()
	if err != nil {
		panic(err)
	}
	fmt.Println(firstResult)

	fmt.Println("Second:")
	secondResult, err := number_picker.NumberPicker(number_picker.Settings{
		Prompt:      "Choose a float",
		DontCleanup: true,
		Decimal:     true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(secondResult)
}
