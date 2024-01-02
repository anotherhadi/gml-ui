package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/asciitext"
)

func main() {
	str := asciitext.AsciiText("Hello world!")
	fmt.Print(str)
}
