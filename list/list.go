// https://github.com/anotherhadi/gml-ui
package list

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/asciimoji"
	"github.com/anotherhadi/gml-ui/getchar"
	"github.com/anotherhadi/gml-ui/settings"
)

func getRange(length, cursor, maxRows int) (start, end int, moreBefore, moreAfter bool) {
	if maxRows > length+1 {
		maxRows = length + 1
	}

	var incrStart, incrEnd int
	incrStart = int(maxRows / 2)
	incrEnd = int(maxRows / 2)
	if maxRows%2 == 0 {
		incrStart--
	}

	if cursor-incrStart < 0 {
		start = 0
		end = start + maxRows - 1
	} else if cursor+incrEnd > length {
		end = length
		start = end - maxRows + 1
	} else {
		start = cursor - incrStart
		end = cursor + incrEnd
	}

	if start != 0 {
		moreBefore = true
		start++
	}
	if end != length {
		moreAfter = true
		end--
	}

	return
}

func printOption(settings settings.Settings, selected bool, title string, description string) {
	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	if selected {
		fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
		fmt.Print("│ ")
	} else {
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print("  ")
	}

	fmt.Print(ansi.Bold)
	fmt.Println(title)
	fmt.Print(ansi.Reset)

	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	if selected {
		fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
		fmt.Print("│ ")
	} else {
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print("  ")
	}

	if len(description) > settings.MaxCols-settings.LeftPadding-2 {
		fmt.Println(description[:settings.MaxCols-settings.LeftPadding-2])
	} else {
		fmt.Println(description)
	}

	fmt.Print(ansi.Reset, "\n")

}

func printOptions(settings settings.Settings, options [][]string, selected int) {

	var maxRows int = settings.MaxRows - settings.TopPadding - settings.BottomPadding
	if maxRows > len(options)*3 {
		maxRows = len(options) * 3
	}

	var start, end int
	var moreAfter, moreBefore bool
	start, end, moreBefore, moreAfter = getRange(len(options), selected, maxRows/3+1)

	fmt.Print(ansi.CursorUpN(maxRows + settings.TopPadding + settings.BottomPadding - maxRows%3))
	fmt.Print(ansi.ScreenClearEnd())
	fmt.Print(strings.Repeat("\n", settings.TopPadding))

	if moreBefore {
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print("\n")
		fmt.Print(strings.Repeat(" ", settings.LeftPadding))
		fmt.Println(asciimoji.Up, asciimoji.Up)
		fmt.Print("\n")
	}

	for index, option := range options[start:end] {
		printOption(settings, index+start == selected, option[0], option[1])
	}

	if moreAfter {
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print("\n")
		fmt.Print(strings.Repeat(" ", settings.LeftPadding))
		fmt.Println(asciimoji.Down, asciimoji.Down)
		fmt.Print("\n")
	}

	fmt.Print(ansi.Reset)
	fmt.Print(strings.Repeat("\n", settings.BottomPadding))
}

func List(options [][]string, customSettings ...settings.Settings) (selected int, err error) {

	settings := settings.GetSettings(customSettings)

	selected = settings.DefaultInt

	var maxRows int = settings.MaxRows - settings.TopPadding - settings.BottomPadding
	if maxRows > len(options)*3 {
		maxRows = len(options) * 3
	}
	fmt.Print(strings.Repeat("\n", maxRows+settings.TopPadding+settings.BottomPadding-maxRows%3))

	fmt.Print(ansi.CursorInvisible())

	for {
		printOptions(settings, options, selected)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			return -1, err
		}

		if arrow == "down" || ascii == 106 { // Down arrow, J
			if selected < len(options)-1 {
				selected++
			} else if !settings.DontLoop {
				selected = 0
			}
		} else if arrow == "up" || ascii == 107 { // Up arrow, K
			if selected > 0 {
				selected--
			} else if !settings.DontLoop {
				selected = len(options) - 1
			}
		} else if ascii == 13 { // Enter
			fmt.Print(ansi.CursorVisible())
			return selected, nil
		} else {
			if settings.ExitOnUnknownKey {
				fmt.Print(ansi.CursorVisible())
				return -1, errors.New("Key not accepted")
			} else if ascii == 3 {
				fmt.Print(ansi.CursorVisible())
				return -1, errors.New("SIGINT")
			}
		}
	}
}
