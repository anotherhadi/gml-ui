package list

import (
	"errors"
	"fmt"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/getchar"
	"github.com/anotherhadi/gml-ui/utils"
)

func printOptions(settings Settings, selected int) {
	for index, option := range settings.Options {
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
		if index == selected {
			fmt.Print(ansi.FgRgb(settings.SelectedTitleForeground.Red, settings.SelectedTitleForeground.Green, settings.SelectedTitleForeground.Blue))
			fmt.Print("| ")
		} else {
			fmt.Print(ansi.FgRgb(settings.UnselectedForeground.Red, settings.UnselectedForeground.Green, settings.UnselectedForeground.Blue))
			fmt.Print("  ")
		}
		fmt.Print(ansi.Bold)
		fmt.Print(option.Title)
		fmt.Print(ansi.Reset, "\n")
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
		if index == selected {
			fmt.Print(ansi.FgRgb(settings.SelectedTitleForeground.Red, settings.SelectedTitleForeground.Green, settings.SelectedTitleForeground.Blue))
			fmt.Print("| ")
			fmt.Print(ansi.FgRgb(settings.SelectedDescriptionForeground.Red, settings.SelectedDescriptionForeground.Green, settings.SelectedDescriptionForeground.Blue))
		} else {
			fmt.Print(ansi.FgRgb(settings.UnselectedForeground.Red, settings.UnselectedForeground.Green, settings.UnselectedForeground.Blue))
			fmt.Print("  ")
		}
		fmt.Print(option.Description)
		fmt.Print(ansi.Reset)
		fmt.Print("\n")
		fmt.Print("\n")
	}
	ansi.CursorUp(uint8(len(settings.Options)) * 3)

}

func List(customSettings ...Settings) (selected int, err error) {

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

	var blankLine int = int(len(settings.Options)*3 + 1)
	fmt.Print(utils.Repeat("\n", blankLine))
	ansi.CursorUp(uint8(blankLine))

	ansi.CursorSave()
	ansi.CursorInvisible()
	fmt.Print("\n")

	for {
		printOptions(settings, selected)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			utils.Cleanup(uint8(len(settings.Options))*3, !settings.DontCleanup)
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
			utils.Cleanup(uint8(len(settings.Options))*3, !settings.DontCleanup)
			return -1, errors.New("Key not accepted")
		}

		if ascii == 13 { // Enter
			utils.Cleanup(uint8(len(settings.Options))*3, !settings.DontCleanup)
			return selected, nil
		}
	}
}
