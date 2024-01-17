package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/confirm_inline"
)

func main() {
	result, err := confirm_inline.ConfirmInline("Are you sure?")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
