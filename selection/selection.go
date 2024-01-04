package selection

import (
	"errors"
	"fmt"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/getchar"
	"github.com/anotherhadi/gml-ui/utils"
)

func printPrompt(settings Settings) {
	fmt.Print("\n")
	fmt.Print(ansi.FgRgb(settings.PromptForeground.Red, settings.PromptForeground.Green, settings.PromptForeground.Blue))
	var splitedPrompt []string = utils.SplitPrompt(settings.Prompt, int(settings.MaxCols))
	for _, line := range splitedPrompt {
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
		fmt.Println(line)
	}
	fmt.Print(ansi.Reset)
	fmt.Print("\n")
}

func printOptions(settings Settings, selected int) {
	for index, option := range settings.Options {
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
		if index == selected {
			fmt.Print(ansi.FgRgb(settings.SelectedForeground.Red, settings.SelectedForeground.Green, settings.SelectedForeground.Blue))
			fmt.Print("> ")
		} else {
			fmt.Print(ansi.FgRgb(settings.UnselectedForeground.Red, settings.UnselectedForeground.Green, settings.UnselectedForeground.Blue))
			fmt.Print("  ")
		}
		fmt.Print(option)
		fmt.Print(ansi.Reset)
		fmt.Print("\n")
	}
	ansi.CursorUp(len(settings.Options))

}

func Selection(customSettings ...Settings) (selected int, err error) {

	var settings Settings

	if len(customSettings) > 0 {
		settings, err = combineSettings(customSettings[0])
		if err != nil {
			return -1, err
		}
	} else {
		settings = getDefaultSettings()
	}

	selected = 0

	var blankLine int = int(len(settings.Options)) + 1
	if settings.Prompt != "noprompt" {
		blankLine += int(len(settings.Prompt)/int(settings.MaxCols)) + 2
	}
	fmt.Print(utils.Repeat("\n", blankLine))
	ansi.CursorUp(blankLine)

	ansi.CursorSave()
	ansi.CursorInvisible()

	if settings.Prompt != "noprompt" {
		printPrompt(settings)
	}

	for {
		printOptions(settings, selected)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			utils.Cleanup(uint8(len(settings.Options)), !settings.DontCleanup)
			return -1, err
		}

		if arrow == "down" || ascii == 106 { // Down arrow, J
			if selected < len(settings.Options)-1 {
				selected++
			}
		} else if arrow == "up" || ascii == 107 { // Up arrow, K
			if selected > 0 {
				selected--
			}
		} else if !(ascii == 13) {
			if settings.UnknownKeysErr {
				utils.Cleanup(uint8(len(settings.Options)), !settings.DontCleanup)
				return -1, errors.New("Key not accepted")
			} else if ascii == 3 {
				utils.Cleanup(uint8(len(settings.Options)), !settings.DontCleanup)
				return -1, errors.New("SIGINT")
			}
		}

		if ascii == 13 { // Enter
			utils.Cleanup(uint8(len(settings.Options)), !settings.DontCleanup)
			return selected, nil
		}
	}
}
