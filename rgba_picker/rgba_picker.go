// https://github.com/anotherhadi/gml-ui
package rgba_picker

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/asciimoji"
	"github.com/anotherhadi/gml-ui/getchar"
	"github.com/anotherhadi/gml-ui/settings"
)

type boxStyle struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
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
			Verticaly:   "│",
			Horizontaly: "─",
		},
		"=": {
			TopLeft:     "╔",
			TopRight:    "╗",
			BottomLeft:  "╚",
			BottomRight: "╝",
			Verticaly:   "║",
			Horizontaly: "═",
		},
		"none": {
			TopLeft:     " ",
			TopRight:    " ",
			BottomLeft:  " ",
			BottomRight: " ",
			Verticaly:   " ",
			Horizontaly: " ",
		},
	}
	return styles[s]
}

func printRgbaPicker(settings settings.Settings, rgba [4]int, selected int) {
	fmt.Print(ansi.CursorUpN(settings.TopPadding + 5 + settings.BottomPadding))
	fmt.Print(ansi.ScreenClearDown())
	fmt.Print(strings.Repeat("\n", settings.TopPadding))
	length := 5*4 + 3
	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
	boxStyle := getBoxStyle(settings.Style)

	fmt.Print("   ")
	for i := 0; i < 4; i++ {
		if i == selected {
			fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
		}
		fmt.Print(asciimoji.Up)
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print(strings.Repeat(" ", 5))
	}
	fmt.Print("\n")
	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(boxStyle.TopLeft)
	fmt.Print(strings.Repeat(boxStyle.Horizontaly, length))
	fmt.Print(boxStyle.TopRight)
	fmt.Print("   ")
	fmt.Print(ansi.BgRgb(uint8(rgba[0]), uint8(rgba[1]), uint8(rgba[2])))
	fmt.Print("   ")
	fmt.Print(ansi.Reset)
	fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
	fmt.Print("\n")

	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(boxStyle.Verticaly)
	for i := 0; i < 4; i++ {
		fmt.Print(" ")
		if i == selected {
			fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
		}
		fmt.Printf("%3d", rgba[i])
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print(" ")
		fmt.Print(boxStyle.Verticaly)
	}
	fmt.Print("   ")
	fmt.Print(ansi.BgRgb(uint8(rgba[0]), uint8(rgba[1]), uint8(rgba[2])))
	fmt.Print("   ")
	fmt.Print(ansi.Reset)
	fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))

	fmt.Print("\n")
	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(boxStyle.BottomLeft)
	fmt.Print(strings.Repeat(boxStyle.Horizontaly, length))
	fmt.Print(boxStyle.BottomRight)
	fmt.Print("   ")
	fmt.Print(ansi.BgRgb(uint8(rgba[0]), uint8(rgba[1]), uint8(rgba[2])))
	fmt.Print("   ")
	fmt.Print(ansi.Reset)
	fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
	fmt.Print("\n")
	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print("   ")
	for i := 0; i < 4; i++ {
		if i == selected {
			fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
		}
		fmt.Print(asciimoji.Down)
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print(strings.Repeat(" ", 5))
	}
	fmt.Print("\n")

	fmt.Print(strings.Repeat("\n", settings.BottomPadding))

	fmt.Print(ansi.Reset)
}

func RgbaPicker(customSettings ...settings.Settings) (rgba [4]int, err error) {

	settings := settings.GetSettings(customSettings)

	rgba = [4]int{int(settings.DefaultColor.Red), int(settings.DefaultColor.Green), int(settings.DefaultColor.Blue), 255}
	var selected int = 0

	fmt.Print(ansi.CursorInvisible())

	fmt.Print(strings.Repeat("\n", settings.TopPadding+5+settings.BottomPadding))

	for {
		printRgbaPicker(settings, rgba, selected)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			fmt.Print(ansi.CursorVisible())
			return [4]int{}, err
		}

		if arrow == "left" || ascii == 104 { // Left arrow or H
			if selected != 0 {
				selected--
			}
		} else if arrow == "right" || ascii == 108 { // Right arrow or L
			if selected != 3 {
				selected++
			}
		} else if arrow == "down" || ascii == 106 { // Down arrow or J
			if rgba[selected] != 0 {
				rgba[selected]--
			}
		} else if arrow == "up" || ascii == 107 { // Up Arrow or K
			if rgba[selected] != 255 {
				rgba[selected]++
			}
		} else if ascii >= 48 && ascii <= 57 { // Manually input a number
			number := int(ascii) - 48
			if rgba[selected]*10+number <= 255 {
				rgba[selected] = rgba[selected]*10 + number
			}
		} else if ascii == 127 { // Del
			rgba[selected] = int(rgba[selected] / 10)
		} else if ascii == 13 { // CR
			fmt.Print(ansi.CursorVisible())
			return rgba, nil
		} else {
			if settings.ExitOnUnknownKey {
				fmt.Print(ansi.CursorVisible())
				return [4]int{}, errors.New("Key not accepted")
			} else if ascii == 3 {
				fmt.Print(ansi.CursorVisible())
				return [4]int{}, errors.New("SIGINT")
			}
		}
	}
}
