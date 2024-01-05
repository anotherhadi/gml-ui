package list

import (
	"errors"
	"fmt"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/asciimoji"
	"github.com/anotherhadi/gml-ui/getchar"
	"github.com/anotherhadi/gml-ui/utils"
)

func printOptions(settings Settings, selected int, maxOptions int) {
	var goUp int

	var startIndex int = 0
	var endIndex int = startIndex + maxOptions + 1
	if endIndex > len(settings.Options)-1 {
		endIndex--
	}

	if selected >= endIndex {
		startIndex = selected - maxOptions + 1
		endIndex = selected + 1
	}

	ansi.ClearScreenEnd()

	if startIndex > 0 {
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)+2))
		fmt.Print(ansi.FgRgb(settings.UnselectedForeground.Red, settings.UnselectedForeground.Green, settings.UnselectedForeground.Blue))
		fmt.Print(asciimoji.Up, "   ", asciimoji.Up, "\n\n\n")
		goUp += 3
	}

	for index, option := range settings.Options[startIndex:endIndex] {
		goUp += 3
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
		if index+startIndex == selected {
			fmt.Print(ansi.FgRgb(settings.SelectedTitleForeground.Red, settings.SelectedTitleForeground.Green, settings.SelectedTitleForeground.Blue))
			fmt.Print("│ ")
		} else {
			fmt.Print(ansi.FgRgb(settings.UnselectedForeground.Red, settings.UnselectedForeground.Green, settings.UnselectedForeground.Blue))
			fmt.Print("  ")
		}
		fmt.Print(ansi.Bold)
		fmt.Print(option.Title)
		fmt.Print(ansi.Reset, "\n")
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
		if index+startIndex == selected {
			fmt.Print(ansi.FgRgb(settings.SelectedTitleForeground.Red, settings.SelectedTitleForeground.Green, settings.SelectedTitleForeground.Blue))
			fmt.Print("│ ")
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
	if endIndex < len(settings.Options) {
		fmt.Print("\n")
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)+2))
		fmt.Print(ansi.FgRgb(settings.UnselectedForeground.Red, settings.UnselectedForeground.Green, settings.UnselectedForeground.Blue))
		fmt.Print(asciimoji.Down, "   ", asciimoji.Down, "\n")
		goUp += 2
	}
	ansi.CursorUp(goUp)
	ansi.CursorCol(0)

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
	maxOptions := int(settings.MaxRows/3) - 2
	var blankLine int
	if maxOptions > len(settings.Options) {
		maxOptions = len(settings.Options)
		blankLine = int(maxOptions) * 3
	} else {
		blankLine = settings.MaxRows + 3
	}
	if maxOptions == 0 {
		return -1, errors.New("Need more rows")
	}

	fmt.Print(utils.Repeat("\n", blankLine))
	ansi.CursorUp(blankLine)

	ansi.CursorSave()
	ansi.CursorInvisible()

	for {
		printOptions(settings, selected, maxOptions)

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
			if settings.UnknownKeysErr {
				utils.Cleanup(uint8(len(settings.Options))*3, !settings.DontCleanup)
				return -1, errors.New("Key not accepted")
			} else if ascii == 3 {
				utils.Cleanup(uint8(len(settings.Options))*3, !settings.DontCleanup)
				return -1, errors.New("SIGINT")
			}
		}

		if ascii == 13 { // Enter
			utils.Cleanup(uint8(len(settings.Options))*3, !settings.DontCleanup)
			return selected, nil
		}
	}
}
