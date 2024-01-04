package main

import (
	"fmt"

	rgbapicker "github.com/anotherhadi/gml-ui/rgba_picker"
)

func main() {
	rgba, err := rgbapicker.RgbaPicker()
	fmt.Print(rgba, err)
}
