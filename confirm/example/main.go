package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/confirm"
)

func main() {
	result, err := confirm.Confirm()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
