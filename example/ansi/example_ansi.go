package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/ansi"
)

func main() {
	ansi.CursorRight(5)
	fmt.Print(ansi.Red, "Hello ")
	fmt.Print(ansi.FgRgb(69, 71, 90), "World!\n")
	ansi.CursorRight(3)
	fmt.Print(ansi.BrightCyan, ansi.Bold, ansi.Underline, "I use ANSI escape code\n")
}
