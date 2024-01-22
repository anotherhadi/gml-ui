// https://github.com/anotherhadi/gml-ui
package checkbox

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

func printOption(settings settings.Settings, selected bool, option string, checked bool) {
	fmt.Print(strings.Repeat(" ", settings.LeftPadding))
	fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor) + "[")
	if checked {
		fmt.Print(ansi.FgRgbSettings(settings.AccentColor) + "*")
	} else {
		fmt.Print(" ")
	}
	fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor) + "] ")
	if selected {
		fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
	}
	if len(option) > settings.MaxCols-settings.LeftPadding-4 {
		fmt.Print(option[:settings.MaxCols-settings.LeftPadding-4])
	} else {
		fmt.Print(option)
	}
	fmt.Print(ansi.Reset + "\n")
}

func printOptions(settings settings.Settings, selected int, options []string, checked []bool) {
	var maxRows int = settings.MaxRows - settings.TopPadding - settings.BottomPadding
	if maxRows > len(options) {
		maxRows = len(options)
	}

	start, end, moreBefore, moreAfter := getRange(len(options), selected, maxRows+1)

	fmt.Print(ansi.CursorUpN(maxRows + settings.TopPadding + settings.BottomPadding))
	fmt.Print(ansi.ScreenClearEnd())
	fmt.Print(strings.Repeat("\n", settings.TopPadding))

	if moreBefore {
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print(strings.Repeat(" ", settings.LeftPadding))
		fmt.Println(asciimoji.Up, asciimoji.Up)
	}

	for i, option := range options[start:end] {
		printOption(settings, i+start == selected, option, checked[i+start])
	}

	if moreAfter {
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print(strings.Repeat(" ", settings.LeftPadding))
		fmt.Println(asciimoji.Down, asciimoji.Down)
	}

	fmt.Print(ansi.Reset)
	fmt.Print(strings.Repeat("\n", settings.BottomPadding))
}

func Checkbox(options []string, customSettings ...settings.Settings) (checked []bool, err error) {

	settings := settings.GetSettings(customSettings)

	var selected int = 0
	if settings.DefaultChecked != nil && len(settings.DefaultChecked) != len(options) {
		return nil, errors.New("Length of Options and DefaultChecked not equal")
	} else if settings.DefaultChecked != nil {
		checked = settings.DefaultChecked
	} else {
		checked = make([]bool, len(options))
	}

	fmt.Print(ansi.CursorInvisible())

	if len(options) > settings.MaxRows-settings.TopPadding-settings.BottomPadding {
		fmt.Print(strings.Repeat("\n", settings.MaxRows))
	} else {
		fmt.Print(strings.Repeat("\n", settings.TopPadding+len(options)+settings.BottomPadding))
	}

	for {
		printOptions(settings, selected, options, checked)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			fmt.Print(ansi.CursorVisible())
			return nil, err
		}

		if arrow == "down" || ascii == 106 { // Down arrow, J
			selected++
			if settings.DontLoop && selected > len(options)-1 {
				selected = len(options) - 1
			} else if selected > len(options)-1 {
				selected = 0
			}
		} else if arrow == "up" || ascii == 107 { // Up arrow, K
			selected--
			if settings.DontLoop && selected < 0 {
				selected = 0
			} else if selected < 0 {
				selected = len(options) - 1
			}
		} else if ascii == 32 { // Space
			checked[selected] = !checked[selected]
		} else if ascii == 13 { // Enter
			nselected := 0
			for _, check := range checked {
				if check {
					nselected++
				}
			}
			if nselected <= int(settings.Maximum) && nselected >= int(settings.Minimum) {
				fmt.Print(ansi.CursorVisible())
				return checked, nil
			}
		} else {
			if settings.ExitOnUnknownKey {
				fmt.Print(ansi.CursorVisible())
				return nil, errors.New("Key not accepted")
			} else if ascii == 3 { // Ctrl C
				fmt.Print(ansi.CursorVisible())
				return nil, errors.New("SIGINT")
			}
		}
	}
}
