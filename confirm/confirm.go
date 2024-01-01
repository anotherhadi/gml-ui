package confirm

import (
	"errors"
	"fmt"
	"math"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/getchar"
	"github.com/anotherhadi/gml-ui/utils"
)

func printPrompt(settings Settings) {
	var promptLength int = len(settings.Prompt)
	var optionsLength int = len(settings.Affirmative) + len(settings.Negative) + 13
	fmt.Print("\n")
	fmt.Print(ansi.FgRgb(settings.PromptForeground.Red, settings.PromptForeground.Green, settings.PromptForeground.Blue))
	var splitedPrompt []string = utils.SplitPrompt(settings.Prompt, int(settings.MaxCols))
	for _, line := range splitedPrompt {
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
		fmt.Print(utils.Repeat(" ", (optionsLength-promptLength)/2))
		fmt.Println(line)
	}
	fmt.Print(ansi.Reset)
	fmt.Print("\n")
}

func printOptions(settings Settings, selection bool) {
	var promptLength int = int(math.Min(float64(len(settings.Prompt)), float64(settings.MaxCols)))
	var optionsLength int = len(settings.Affirmative) + len(settings.Negative) + 13

	fmt.Print("\r")
	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(utils.Repeat(" ", (promptLength-optionsLength)/2))

	if selection {
		fmt.Print(ansi.BgRgb(settings.SelectedColor.BackgroundColor.Red, settings.SelectedColor.BackgroundColor.Green, settings.SelectedColor.BackgroundColor.Blue))
		fmt.Print(ansi.FgRgb(settings.SelectedColor.ForegroundColor.Red, settings.SelectedColor.ForegroundColor.Green, settings.SelectedColor.ForegroundColor.Blue))
	} else {
		fmt.Print(ansi.BgRgb(settings.UnselectedColor.BackgroundColor.Red, settings.UnselectedColor.BackgroundColor.Green, settings.UnselectedColor.BackgroundColor.Blue))
		fmt.Print(ansi.FgRgb(settings.UnselectedColor.ForegroundColor.Red, settings.UnselectedColor.ForegroundColor.Green, settings.UnselectedColor.ForegroundColor.Blue))
	}
	fmt.Print("   ")
	fmt.Print(settings.Affirmative)
	fmt.Print("   ")

	fmt.Print(ansi.Reset, "  ")

	if !selection {
		fmt.Print(ansi.BgRgb(settings.SelectedColor.BackgroundColor.Red, settings.SelectedColor.BackgroundColor.Green, settings.SelectedColor.BackgroundColor.Blue))
		fmt.Print(ansi.FgRgb(settings.SelectedColor.ForegroundColor.Red, settings.SelectedColor.ForegroundColor.Green, settings.SelectedColor.ForegroundColor.Blue))
	} else {
		fmt.Print(ansi.BgRgb(settings.UnselectedColor.BackgroundColor.Red, settings.UnselectedColor.BackgroundColor.Green, settings.UnselectedColor.BackgroundColor.Blue))
		fmt.Print(ansi.FgRgb(settings.UnselectedColor.ForegroundColor.Red, settings.UnselectedColor.ForegroundColor.Green, settings.UnselectedColor.ForegroundColor.Blue))
	}
	fmt.Print("   ")
	fmt.Print(settings.Negative)
	fmt.Print("   ")

	fmt.Print(ansi.Reset)
}

func Confirm(customSettings ...Settings) (selected bool, err error) {

	var settings Settings

	if len(customSettings) > 0 {
		settings = combineSettings(customSettings[0])
	} else {
		settings = getDefaultSettings()
	}

	selected = !settings.DefaultToFalse

	var blankLine int = 5 + int(len(settings.Prompt)/int(settings.MaxCols))
	fmt.Print(utils.Repeat("\n", blankLine))
	ansi.CursorUp(uint8(blankLine))

	ansi.CursorSave()
	ansi.CursorInvisible()

	printPrompt(settings)

	for {
		printOptions(settings, selected)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			utils.Cleanup(2, !settings.DontCleanup)
			return selected, err
		}

		if arrow == "left" || ascii == 104 || ascii == 121 { // Left arrow, H or Y
			selected = true
		} else if arrow == "right" || ascii == 108 || ascii == 110 { // Right arrow, L or N
			selected = false
		} else if !(ascii == 13) {
			utils.Cleanup(2, !settings.DontCleanup)
			return selected, errors.New("Key not accepted")
		}

		if ascii == 13 || ascii == 121 || ascii == 110 { // Enter, Y, N
			utils.Cleanup(2, !settings.DontCleanup)
			return selected, nil
		}
	}
}
