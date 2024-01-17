package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/getchar"
)

func main() {
	ascii, arrow, err := getchar.GetChar()
	if err != nil {
		panic(err)
	}
	fmt.Println(ascii, arrow)
}
