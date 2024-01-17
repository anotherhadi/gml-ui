// https://github.com/anotherhadi/gml-ui
package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/settings"
)

func printPrompt(prompt string, settings settings.Settings) {
	fmt.Print(strings.Repeat("\n", settings.TopPadding))
	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(ansi.FgRgbSettings(settings.TextColor))
	fmt.Print(prompt, " ")
	fmt.Print(ansi.Reset)
}

func Input(prompt string, customSettings ...settings.Settings) (result string, err error) {

	settings := settings.GetSettings(customSettings)

	printPrompt(prompt, settings)

	var input string

	fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
	inputReader := bufio.NewReader(os.Stdin)
	input, err = inputReader.ReadString('\n')
	fmt.Print(ansi.Reset)
	if err != nil {
		return "", err
	}
	fmt.Print(strings.Repeat("\n", settings.BottomPadding))
	if input == "\n" {
		return settings.DefaultString, nil
	}
	return input[:len(input)-1], nil
}
