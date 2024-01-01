package input

import (
	"bufio"
	"fmt"
	"os"

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

func Input(customSettings ...Settings) (result string, err error) {

	var settings Settings

	if len(customSettings) > 0 {
		settings = combineSettings(customSettings[0])
	} else {
		settings = getDefaultSettings()
	}

	printPrompt(settings)

	var input string

	fmt.Print(ansi.FgRgb(settings.InputForeground.Red, settings.InputForeground.Green, settings.InputForeground.Blue))
	inputReader := bufio.NewReader(os.Stdin)
	input, err = inputReader.ReadString('\n')
	fmt.Print(ansi.Reset)
	if err != nil {
		return "", err
	}
	if !settings.DontCleanup {
		ansi.CursorUp(2)
		ansi.ClearScreenEnd()
	} else {
		fmt.Print("\n")
	}
	if input == "\n" {
		return settings.Default, nil
	}
	return input[:len(input)-1], nil
}
