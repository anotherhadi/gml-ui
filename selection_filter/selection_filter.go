package selection_filter

import (
	"bytes"
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

func printOptions(settings Settings, options []string, selected int, filter string) {
	ansi.ClearScreenEnd()
	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(ansi.FgRgb(settings.PromptForeground.Red, settings.PromptForeground.Green, settings.PromptForeground.Blue))
	fmt.Print("Filter: ")
	fmt.Print(ansi.FgRgb(settings.UnselectedForeground.Red, settings.UnselectedForeground.Green, settings.UnselectedForeground.Blue))
	fmt.Print(filter)
	fmt.Print("\n")

	var startIndex uint8 = 0
	var endIndex uint8 = startIndex + settings.MaxRows

	if selected >= int(endIndex) {
		startIndex = uint8(selected) - settings.MaxRows + 1
		endIndex = uint8(selected) + 1
	}

	if startIndex > 0 {
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)+2))
		fmt.Print(ansi.FgRgb(settings.UnselectedForeground.Red, settings.UnselectedForeground.Green, settings.UnselectedForeground.Blue))
		fmt.Print("^...^\n")
	}
	for index, option := range options {
		if index >= int(startIndex) && index < int(endIndex) {
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
	}
	if len(options) > int(endIndex) {
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)+2))
		fmt.Print(ansi.FgRgb(settings.UnselectedForeground.Red, settings.UnselectedForeground.Green, settings.UnselectedForeground.Blue))
		fmt.Print("v...v\n")
		fmt.Print(ansi.Reset)
		ansi.CursorUp(1)
	}
	if startIndex > 0 {
		ansi.CursorUp(1)
	}
	if settings.MaxRows < uint8(len(options)) {
		ansi.CursorUp(settings.MaxRows + 1)
	} else {
		ansi.CursorUp(uint8(len(options)) + 1)
	}

}

func SelectionFilter(customSettings ...Settings) (selected int, err error) {

	var settings Settings

	if len(customSettings) > 0 {
		settings = combineSettings(customSettings[0])
	} else {
		settings = getDefaultSettings()
	}

	selected = 0
	var filterBuffer bytes.Buffer

	var blankLine int
	if int(settings.MaxRows) < int(len(settings.Options)) {
		blankLine = 3 + int(settings.MaxRows)
	} else {
		blankLine = 2 + int(len(settings.Options))
	}
	if settings.Prompt != "noprompt" {
		blankLine += int(len(settings.Prompt)/int(settings.MaxCols)) + 3
	}
	fmt.Print(utils.Repeat("\n", blankLine))
	ansi.CursorUp(uint8(blankLine))

	ansi.CursorSave()
	ansi.CursorInvisible()

	if settings.Prompt != "noprompt" {
		printPrompt(settings)
	}

	for {
		var filteredOptions []string = utils.FilterStringsByPrefix(settings.Options, filterBuffer.String(), settings.CaseSensitive)
		printOptions(settings, filteredOptions, selected, filterBuffer.String())

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			utils.Cleanup(uint8(len(filteredOptions))+2, !settings.DontCleanup)
			return selected, err
		}

		if arrow == "down" { // Down arrow
			if selected < len(filteredOptions)-1 {
				selected++
			}
		} else if arrow == "up" { // Up arrow
			if selected > 0 {
				selected--
			}
		} else if (ascii >= 65 && ascii <= 90) || (ascii >= 97 && ascii <= 122) || (ascii >= 48 && ascii <= 57) || ascii == 32 { // Letters, Numbers, Space
			selected = 0
			filterBuffer.WriteString(string(ascii))
		} else if ascii == 127 { // DEL
			selected = 0
			if filterBuffer.Len() > 0 {
				filterBuffer.Truncate(filterBuffer.Len() - 1)
			}
		} else if ascii == 13 { // CR
			utils.Cleanup(uint8(len(filteredOptions))+2, !settings.DontCleanup)
			if len(filteredOptions) == 0 {
				return -1, errors.New("Returned nothing")
			}
			for index, str := range settings.Options {
				if str == filteredOptions[selected] {
					return index, nil
				}
			}
			return -1, errors.New("Unable to find in list")
		} else {
			if settings.UnknownKeysErr {
				utils.Cleanup(uint8(len(filteredOptions))+2, !settings.DontCleanup)
				return -1, errors.New("Key not accepted")
			} else if ascii == 3 {
				utils.Cleanup(uint8(len(filteredOptions))+2, !settings.DontCleanup)
				return -1, errors.New("SIGINT")
			}
		}

	}
}
