package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/rgba_picker"
)

func main() {
	rgba, err := rgba_picker.RgbaPicker()
	if err != nil {
		panic(err)
	}
	fmt.Println(rgba)
}
