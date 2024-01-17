// https://github.com/anotherhadi/gml-ui
package table

import (
	"fmt"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/settings"
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

func Table(table [][]string, customSettings ...settings.Settings) error {

	settings := settings.GetSettings(customSettings)

	var cols_length []int = make([]int, len(table[0]))
	var boxStyle boxStyle = getBoxStyle(settings.Style)

	for _, row := range table {
		for icol, col := range row {
			if len(col) > int(settings.MaxCols) {
				cols_length[icol] = int(settings.MaxCols)
				break
			}
			if len(col) > cols_length[icol] {
				cols_length[icol] = len(col)
			}
		}
	}

	fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(boxStyle.TopLeft)
	for i, col := range cols_length {
		fmt.Print(strings.Repeat(boxStyle.Horizontaly, col+2))
		if i < len(cols_length)-1 {
			fmt.Print(boxStyle.ToDown)
		}
	}
	fmt.Print(boxStyle.TopRight)
	fmt.Print("\n")

	for irow, row := range table {
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))

		fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
		fmt.Print(boxStyle.Verticaly)
		for icol, col := range row {
			fmt.Print(" ")
			var length int
			fmt.Print(ansi.FgRgbSettings(settings.TextColor))
			if irow == 0 {
				fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
			}
			if len(col) > int(settings.MaxCols) {
				length = int(settings.MaxCols)
				fmt.Print(col[:settings.MaxCols-2])
				fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
				fmt.Print("..")
			} else {
				length = len(col)
				fmt.Print(col)
			}
			fmt.Print(strings.Repeat(" ", cols_length[icol]-length))
			fmt.Print(" ")
			if icol < len(row)-1 {
				fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
				fmt.Print(boxStyle.Verticaly)
			}
		}
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print(boxStyle.Verticaly)
		fmt.Print("\n")
		if irow != len(table)-1 {
			fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
			fmt.Print(boxStyle.ToRight)
			for i, col := range cols_length {
				fmt.Print(strings.Repeat(boxStyle.Horizontaly, col+2))
				if i < len(cols_length)-1 {
					fmt.Print(boxStyle.Middle)
				}
			}
			fmt.Print(boxStyle.ToLeft)
			fmt.Print("\n")
		}
	}

	fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(boxStyle.BottomLeft)
	for i, col := range cols_length {
		fmt.Print(strings.Repeat(boxStyle.Horizontaly, col+2))
		if i < len(cols_length)-1 {
			fmt.Print(boxStyle.ToUp)
		}
	}
	fmt.Print(boxStyle.BottomRight)
	fmt.Print("\n")
	fmt.Print(ansi.Reset)

	return nil
}
