// https://github.com/anotherhadi/gml-ui
package confirm_inline

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/settings"
)

func printConfirm(prompt string, settings settings.Settings) {
	fmt.Print(ansi.FgRgbSettings(settings.TextColor))
	fmt.Print(prompt, " ")
	fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
	fmt.Print("[")
	if settings.DefaultBool {
		fmt.Print("Y/n")
	} else {
		fmt.Print("y/N")
	}
	fmt.Print("] ")
	fmt.Print(ansi.Reset)
}

func ConfirmInline(prompt string, customSettings ...settings.Settings) (result bool, err error) {

	settings := settings.GetSettings(customSettings)

	fmt.Print(strings.Repeat("\n", settings.TopPadding))
	printConfirm(prompt, settings)

	var input string

	n, err := fmt.Scanf("%s", &input)
	fmt.Print(strings.Repeat("\n", settings.BottomPadding))
	if err != nil && n != 0 {
		return
	}
	if n == 0 {
		return settings.DefaultBool, nil
	}
	if strings.ToLower(input) == "yes" || strings.ToLower(input) == "y" {
		return true, nil
	}
	if strings.ToLower(input) == "no" || strings.ToLower(input) == "n" {
		return false, nil
	}
	return false, errors.New("Unexpected Input")
}
