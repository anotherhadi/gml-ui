package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/asciitext"
	"github.com/anotherhadi/gml-ui/settings"
)

func main() {
	myBigText := asciitext.AsciiText("Hello world!")
	fmt.Print(myBigText)

	// With custom settings
	myBigText = asciitext.AsciiText("Hello world!", settings.Settings{
		TopPadding:    1,
		BottomPadding: 1,
		LeftPadding:   3,
	})
	fmt.Print(myBigText)
}
