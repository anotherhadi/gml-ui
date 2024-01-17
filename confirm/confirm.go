// https://github.com/anotherhadi/gml-ui
package confirm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/getchar"
	"github.com/anotherhadi/gml-ui/settings"
)

func printOptions(settings settings.Settings, selection bool) {
	ansi.CursorUpN(1 + settings.TopPadding + settings.BottomPadding)
	fmt.Print(strings.Repeat("\n", settings.TopPadding))
	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))

	if selection {
		fmt.Print(ansi.BgRgbSettings(settings.AccentColor))
		fmt.Print(ansi.FgRgbSettings(settings.AccentBackgroundColor))
	} else {
		fmt.Print(ansi.BgRgbSettings(settings.SecondaryColor))
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryBackgroundColor))
	}
	fmt.Print("   ")
	fmt.Print(settings.Affirmative)
	fmt.Print("   ")

	fmt.Print(ansi.Reset, "  ")

	if !selection {
		fmt.Print(ansi.BgRgbSettings(settings.AccentColor))
		fmt.Print(ansi.FgRgbSettings(settings.AccentBackgroundColor))
	} else {
		fmt.Print(ansi.BgRgbSettings(settings.SecondaryColor))
		fmt.Print(ansi.FgRgbSettings(settings.SecondaryBackgroundColor))
	}
	fmt.Print("   ")
	fmt.Print(settings.Negative)
	fmt.Print("   ")

	fmt.Print(ansi.Reset, "\n")
	fmt.Print(strings.Repeat("\n", settings.BottomPadding))
}

func Confirm(customSettings ...settings.Settings) (result bool, err error) {

	settings := settings.GetSettings(customSettings)

	result = settings.DefaultBool

	fmt.Print(strings.Repeat("\n", settings.TopPadding+1+settings.BottomPadding))
	ansi.CursorInvisible()

	for {
		printOptions(settings, result)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			ansi.CursorVisible()
			return false, err
		}

		if arrow == "left" || ascii == 104 || ascii == 121 { // Left arrow, H or Y
			result = true
		} else if arrow == "right" || ascii == 108 || ascii == 110 { // Right arrow, L or N
			result = false
		} else if !(ascii == 13) {
			if settings.ExitOnUnknownKey {
				ansi.CursorVisible()
				return false, errors.New("Key not accepted")
			} else if ascii == 3 {
				ansi.CursorVisible()
				return false, errors.New("SIGINT")
			}
		}

		if ascii == 13 || ascii == 121 || ascii == 110 { // Enter, Y, N
			ansi.CursorVisible()
			return result, nil
		}
	}
}
