package table

import (
	"errors"
	"fmt"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/utils"
)

type boxStyle struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	ToLeft      string
	ToRight     string
	ToUp        string
	ToDown      string
	Middle      string
	Verticaly   string
	Horizontaly string
}

func getBoxStyle(s string) boxStyle {
	styles := map[string]boxStyle{
		"-": {
			TopLeft:     "┌",
			TopRight:    "┐",
			BottomLeft:  "└",
			BottomRight: "┘",
			ToLeft:      "┤",
			ToRight:     "├",
			ToDown:      "┬",
			ToUp:        "┴",
			Middle:      "┼",
			Verticaly:   "│",
			Horizontaly: "─",
		},
		"=": {
			TopLeft:     "╔",
			TopRight:    "╗",
			BottomLeft:  "╚",
			BottomRight: "╝",
			ToLeft:      "╣",
			ToRight:     "╠",
			ToUp:        "╩",
			ToDown:      "╦",
			Middle:      "╬",
			Verticaly:   "║",
			Horizontaly: "═",
		},
		"none": {
			TopLeft:     " ",
			TopRight:    " ",
			BottomLeft:  " ",
			BottomRight: " ",
			ToLeft:      " ",
			ToRight:     " ",
			ToUp:        " ",
			ToDown:      " ",
			Middle:      " ",
			Verticaly:   " ",
			Horizontaly: " ",
		},
	}
	return styles[s]
}

// TODO: Elements wrapping, label decoration, Text alignment
func Table(customSettings ...Settings) error {

	var settings Settings
	var err error

	if len(customSettings) > 0 {
		settings, err = combineSettings(customSettings[0])
		if err != nil {
			return err
		}
	} else {
		return errors.New("No Elements")
	}

	var cols_length []int = make([]int, len(settings.Elements[0]))
	var boxStyle boxStyle = getBoxStyle(settings.Style)
	// var rows_length []int = make([]int, len(settings.Elements))

	for _, row := range settings.Elements {
		for icol, col := range row {
			if len(col) > int(settings.MaxLengthsCol) {
				cols_length[icol] = int(settings.MaxLengthsCol)
				break
			}
			if len(col) > cols_length[icol] {
				cols_length[icol] = len(col)
			}
			// size := int(len(col) / int(settings.MaxLengthsCol))
			// if len(col)%int(settings.MaxLengthsCol) != 0 {
			// 	size++
			// }
			// if size > rows_length[irow] {
			// 	rows_length[irow] = size
			// }
		}
	}

	// var table_width int
	// for _, col := range cols_length {
	// 	table_width += col + 3
	// }
	// table_width -= 1
	// var table_height int
	// for _, row := range rows_length {
	// 	table_height += row + 1
	// }
	// table_height -= 1

	fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(boxStyle.TopLeft)
	for i, col := range cols_length {
		fmt.Print(utils.Repeat(boxStyle.Horizontaly, col+2))
		if i < len(cols_length)-1 {
			fmt.Print(boxStyle.ToDown)
		}
	}
	fmt.Print(boxStyle.TopRight)
	fmt.Print("\n")

	for irow, row := range settings.Elements {
		fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
		fmt.Print(boxStyle.Verticaly)
		for icol, col := range row {
			fmt.Print(" ")
			var length int
			fmt.Print(ansi.FgRgb(settings.TextForeground.Red, settings.TextForeground.Green, settings.TextForeground.Blue))
			if irow == 0 {
				fmt.Print(ansi.FgRgb(settings.LabelForeground.Red, settings.LabelForeground.Green, settings.LabelForeground.Blue))
			}
			if len(col) > int(settings.MaxLengthsCol) {
				length = int(settings.MaxLengthsCol)
				fmt.Print(col[:settings.MaxLengthsCol-2])
				fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
				fmt.Print("..")
			} else {
				length = len(col)
				fmt.Print(col)
			}
			fmt.Print(utils.Repeat(" ", cols_length[icol]-length))
			fmt.Print(" ")
			if icol < len(row)-1 {
				fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
				fmt.Print(boxStyle.Verticaly)
			}
		}
		fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
		fmt.Print(boxStyle.Verticaly)
		fmt.Print("\n")
		if settings.Separator && irow != len(settings.Elements)-1 {
			fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
			fmt.Print(boxStyle.ToRight)
			for i, col := range cols_length {
				fmt.Print(utils.Repeat(boxStyle.Horizontaly, col+2))
				if i < len(cols_length)-1 {
					fmt.Print(boxStyle.Middle)
				}
			}
			fmt.Print(boxStyle.ToLeft)
			fmt.Print("\n")
		}
	}

	fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(boxStyle.BottomLeft)
	for i, col := range cols_length {
		fmt.Print(utils.Repeat(boxStyle.Horizontaly, col+2))
		if i < len(cols_length)-1 {
			fmt.Print(boxStyle.ToUp)
		}
	}
	fmt.Print(boxStyle.BottomRight)
	fmt.Print("\n")
	fmt.Print(ansi.Reset)

	return nil
}
