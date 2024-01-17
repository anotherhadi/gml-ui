// https://github.com/anotherhadi/gml-ui
package selection

import (
	"bytes"
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
func filterStringsByPrefix(original []string, prefix string, caseSensitive bool) []string {
	if prefix == "" {
		return original
	}

	var filtered []string

	if caseSensitive {
		for _, str := range original {
			if strings.HasPrefix(str, prefix) {
				filtered = append(filtered, str)
			}
		}
	} else {
		var lowerPrefix string = strings.ToLower(prefix)
		var lowerStr string
		for _, str := range original {
			lowerStr = strings.ToLower(str)
			if strings.HasPrefix(lowerStr, lowerPrefix) {
				filtered = append(filtered, str)
			}
		}
	}

	return filtered
}
func printOption(settings settings.Settings, selected bool, option string) {
	fmt.Print(strings.Repeat(" ", settings.LeftPadding))

	if selected {
		fmt.Print(ansi.FgRgbSettings(settings.AccentColor) + "> ")
	} else {
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor) + "  ")
	}
	if len(option) > settings.MaxCols-settings.LeftPadding-2 {
		fmt.Print(option[:settings.MaxCols-settings.LeftPadding-2])
	} else {
		fmt.Print(option)
	}
	fmt.Print(ansi.Reset + "\n")
}

func printOptions(settings settings.Settings, options []string, selected int, filter string, previousLength int) {

	var maxRows int = settings.MaxRows - settings.TopPadding - settings.BottomPadding
	if maxRows > previousLength {
		maxRows = previousLength
	}
	ansi.CursorUpN(maxRows + settings.TopPadding + settings.BottomPadding)
	maxRows = settings.MaxRows - settings.TopPadding - settings.BottomPadding
	if maxRows > len(options) {
		maxRows = len(options)
	}

	if settings.Filter {
		ansi.CursorUpN(1)
	}
	ansi.ScreenClearEnd()

	fmt.Print(strings.Repeat("\n", settings.TopPadding))
	if settings.Filter {
		fmt.Print(strings.Repeat(" ", settings.LeftPadding))
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print("Filter: ")
		fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
		ansi.LineClearEnd()
		if filter == "" {
			fmt.Print("<type to add filter>")
		} else {
			fmt.Print(filter)
		}
		fmt.Print("\n")
	}

	start, end, moreBefore, moreAfter := getRange(len(options), selected, maxRows+1)

	if end == -1 {
		end = 0
	}
	if moreBefore {
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print(strings.Repeat(" ", settings.LeftPadding))
		fmt.Println(asciimoji.Up, asciimoji.Up)
	}

	for i, option := range options[start:end] {
		printOption(settings, i+start == selected, option)
	}

	if moreAfter {
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
		fmt.Print(strings.Repeat(" ", settings.LeftPadding))
		fmt.Println(asciimoji.Down, asciimoji.Down)
	}

	fmt.Print(ansi.Reset)
	fmt.Print(strings.Repeat("\n", settings.BottomPadding))

}

func Selection(options []string, customSettings ...settings.Settings) (selected int, err error) {

	settings := settings.GetSettings(customSettings)

	selected = settings.DefaultInt
	var filterBuffer bytes.Buffer
	var previousLength int = len(options)

	if len(options) > settings.MaxRows-settings.TopPadding-settings.BottomPadding {
		fmt.Print(strings.Repeat("\n", settings.MaxRows))
	} else {
		fmt.Print(strings.Repeat("\n", settings.TopPadding+len(options)+settings.BottomPadding))
	}
	if settings.Filter {
		fmt.Print("\n")
	}

	ansi.CursorInvisible()

	for {
		var filteredOptions []string = filterStringsByPrefix(options, filterBuffer.String(), settings.CaseSensitive)
		printOptions(settings, filteredOptions, selected, filterBuffer.String(), previousLength)
		previousLength = len(filteredOptions)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			ansi.CursorVisible()
			return selected, err
		}

		if arrow == "down" { // Down arrow
			selected++
			if settings.DontLoop && selected > len(options)-1 {
				selected = len(options) - 1
			} else if selected > len(options)-1 {
				selected = 0
			}
		} else if arrow == "up" { // Up arrow
			selected--
			if settings.DontLoop && selected < 0 {
				selected = 0
			} else if selected < 0 {
				selected = len(options) - 1
			}
		} else if ((ascii >= 65 && ascii <= 90) || (ascii >= 97 && ascii <= 122) || (ascii >= 48 && ascii <= 57) || ascii == 32) && settings.Filter { // Letters, Numbers, Space
			selected = 0
			filterBuffer.WriteString(string(ascii))
		} else if ascii == 127 { // DEL
			selected = 0
			if filterBuffer.Len() > 0 {
				filterBuffer.Truncate(filterBuffer.Len() - 1)
			}
		} else if ascii == 13 { // CR
			ansi.CursorVisible()
			if len(filteredOptions) == 0 {
				return -1, errors.New("Returned nothing")
			}
			for index, str := range options {
				if str == filteredOptions[selected] {
					return index, nil
				}
			}
			return -1, errors.New("Unable to find in list")
		} else {
			if settings.ExitOnUnknownKey {
				ansi.CursorVisible()
				return -1, errors.New("Key not accepted")
			} else if ascii == 3 {
				ansi.CursorVisible()
				return -1, errors.New("SIGINT")
			}
		}

	}
}
