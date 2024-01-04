package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/rgba_picker"
)

func main() {
	rgba, err := rgba_picker.RgbaPicker()
	fmt.Print(rgba, err)
}
