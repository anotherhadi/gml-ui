package main

import "github.com/anotherhadi/gml-ui/table"

func main() {
	var elements [][]string

	elements = append(elements, []string{"Setting", "Type", "Default"})
	elements = append(elements, []string{"Colors..", "Color", "Cool colors"})
	elements = append(elements, []string{"Style", "string(-,=, )", "-"})
	elements = append(elements, []string{"MaxCols", "int", "80"})
	elements = append(elements, []string{"LeftPadding", "int", "3"})

	err := table.Table(elements)
	if err != nil {
		panic(err)
	}
}
