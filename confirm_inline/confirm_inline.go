package confirm_inline

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/utils"
)

func printPrompt(settings Settings) {
	fmt.Print("\n")
	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(ansi.FgRgb(settings.PromptForeground.Red, settings.PromptForeground.Green, settings.PromptForeground.Blue))
	fmt.Print(settings.Prompt, " ")
	fmt.Print(ansi.Reset)
}

func printConfirm(settings Settings) {
	fmt.Print(ansi.FgRgb(settings.ConfirmForeground.Red, settings.ConfirmForeground.Green, settings.ConfirmForeground.Blue))
	fmt.Print("[")
	if !settings.DefaultToFalse {
		fmt.Print("Y/n")
	} else {
		fmt.Print("y/N")
	}
	fmt.Print("] ")
	fmt.Print(ansi.Reset)
}

func ConfirmInline(customSettings ...Settings) (result bool, err error) {

	var settings Settings

	if len(customSettings) > 0 {
		settings = combineSettings(customSettings[0])
	} else {
		settings = getDefaultSettings()
	}

	printPrompt(settings)
	printConfirm(settings)

	var input string

	n, err := fmt.Scanf("%s", &input)
	if !settings.DontCleanup {
		ansi.CursorUp(2)
		ansi.ClearScreenEnd()
	} else {
		fmt.Print("\n")
	}
	if err != nil && n != 0 {
		return
	}
	if n == 0 {
		return !settings.DefaultToFalse, nil
	}
	if strings.ToLower(input) == "yes" || strings.ToLower(input) == "y" {
		return true, nil
	}
	if strings.ToLower(input) == "no" || strings.ToLower(input) == "n" {
		return false, nil
	}
	return false, errors.New("Unexpected Input")
}
