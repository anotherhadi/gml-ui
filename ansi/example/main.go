package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/ansi"
)

func main() {
	fmt.Print(ansi.ScreenClear())
	fmt.Print(ansi.CursorHome())
	fmt.Println(ansi.Red, ansi.Bold, "Hello world!", ansi.Reset)
	fmt.Print(ansi.CursorMove(4, 4))
	fmt.Println(ansi.Cyan, "I'm Hadi")
}
