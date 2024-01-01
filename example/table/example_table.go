package main

import (
	"github.com/anotherhadi/gml-ui/table"
)

func main() {

	var elements [][]string

	elements = append(elements, []string{"Setting", "Type", "Default"})
	elements = append(elements, []string{"Colors..", "table.RGBColor", "Cool colors"})
	elements = append(elements, []string{"Style", "string(-,=, )", "-"})
	elements = append(elements, []string{"Separator", "bool", "false"})
	elements = append(elements, []string{"Align", "string(left,right,center)", "left"})
	elements = append(elements, []string{"LabelDecoration", "string(default, none, cols)", "default"})
	elements = append(elements, []string{"MaxLengthsCol", "int", "40"})
	elements = append(elements, []string{"LeftPadding", "int", "3"})
	elements = append(elements, []string{"Elements", "[][]string", "nil"})

	err := table.Table(table.Settings{
		Elements:  elements,
		Separator: true,
	})
	if err != nil {
		panic(err)
	}
}
