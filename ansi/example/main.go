package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/ansi"
)

func main() {
	ansi.ScreenClear()
	ansi.CursorHome()
	fmt.Println(ansi.Red, ansi.Bold, "Hello world!", ansi.Reset)
	ansi.CursorMove(4, 4)
	fmt.Println(ansi.Cyan, "I'm Hadi")
}
