package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/getsize"
)

func main() {
	cols, rows, err := getsize.GetSize()
	if err != nil {
		panic(err)
	}
	fmt.Println(cols, rows)
}
