package checkbox

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

func printOptions(settings Settings, selected int, checked []bool) {
	for index, option := range settings.Options {
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))

		fmt.Print(ansi.BrightBlack)
		fmt.Print("[")
		if checked[index] {
			fmt.Print(ansi.FgRgb(settings.CheckedForeground.Red, settings.CheckedForeground.Green, settings.CheckedForeground.Blue))
			fmt.Print("*")

		} else {
			fmt.Print(" ")
		}
		fmt.Print(ansi.BrightBlack)
		fmt.Print("] ")

		if index == selected {
			fmt.Print(ansi.FgRgb(settings.SelectedForeground.Red, settings.SelectedForeground.Green, settings.SelectedForeground.Blue))
		} else {
			fmt.Print(ansi.FgRgb(settings.UnselectedForeground.Red, settings.UnselectedForeground.Green, settings.UnselectedForeground.Blue))
		}
		fmt.Print(option)
		fmt.Print(ansi.Reset)
		fmt.Print("\n")
	}
	ansi.CursorUp(uint8(len(settings.Options)))

}

// TODO: Add minimum/maximum feature
func Checkbox(customSettings ...Settings) (checked []bool, err error) {

	var settings Settings

	if len(customSettings) > 0 {
		settings, err = combineSettings(customSettings[0])
		if err != nil {
			return
		}
	} else {
		settings = getDefaultSettings()
	}

	var selected int = 0
	checked = settings.DefaultOptions

	var blankLine int = 3 + int(len(settings.Options))
	fmt.Print(utils.Repeat("\n", blankLine))
	ansi.CursorUp(uint8(blankLine))

	ansi.CursorSave()
	ansi.CursorInvisible()

	printPrompt(settings)

	for {
		printOptions(settings, selected, checked)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			utils.Cleanup(uint8(len(settings.Options))+1, !settings.DontCleanup)
			return checked, err
		}

		if arrow == "down" || ascii == 106 { // Down arrow, J
			if selected < len(settings.Options)-1 {
				selected++
			}
		} else if arrow == "up" || ascii == 107 { // Up arrow, K
			if selected > 0 {
				selected--
			}
		} else if ascii == 32 {
			checked[selected] = !checked[selected]
		} else if !(ascii == 13) {
			utils.Cleanup(uint8(len(settings.Options))+1, !settings.DontCleanup)
			return checked, errors.New("Key not accepted")
		}

		if ascii == 13 { // Enter
			utils.Cleanup(uint8(len(settings.Options))+1, !settings.DontCleanup)
			return checked, nil
		}
	}
}
