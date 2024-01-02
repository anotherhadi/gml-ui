package main

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/list"
)

func main() {
	selected, err := list.List(list.Settings{
		Options: []list.Options{
			{
				Title:       "First Option",
				Description: "There's a description",
			},
			{
				Title:       "Second one",
				Description: "The next one will not have a description",
			},
			{
				Title: "The last one",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(selected)
}
