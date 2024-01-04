package number_picker

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

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

func printInputPicker(settings Settings, number string, maxLength uint8) {
	ansi.ClearScreenEnd()
	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))

	fmt.Print("  ╔")
	fmt.Print(utils.Repeat("═", int(maxLength)+2))
	fmt.Print("╗\n")

	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	if utils.StringToFloat(number) > settings.Minimum {
		fmt.Print("< ")
	} else {
		fmt.Print("  ")
	}
	fmt.Print("║ ")

	fmt.Print(ansi.FgRgb(settings.InputForeground.Red, settings.InputForeground.Green, settings.InputForeground.Blue))
	fmt.Print(utils.Repeat(" ", int(maxLength)-len(number)))
	fmt.Printf("%s", number)

	fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
	fmt.Print(" ║")
	if utils.StringToFloat(number) < settings.Maximum {
		fmt.Print(" > ")
	} else {
		fmt.Print("  ")
	}

	fmt.Print("\n")
	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))

	fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
	fmt.Print("  ╚")
	fmt.Print(utils.Repeat("═", int(maxLength)+2))
	fmt.Print("╝\n")
	fmt.Print(ansi.Reset)
	ansi.CursorUp(3)
}

func NumberPicker(customSettings ...Settings) (number float64, err error) {

	var settings Settings

	if len(customSettings) > 0 {
		settings = combineSettings(customSettings[0])
	} else {
		settings = getDefaultSettings()
	}

	var manualInputBuffer bytes.Buffer
	manualInputBuffer.WriteString(utils.FloatToString(settings.Default))
	var maxLength uint8 = uint8(len(utils.FloatToString(settings.Maximum)))
	if maxLength < uint8(len(utils.FloatToString(settings.Minimum))) {
		maxLength = uint8(len(utils.FloatToString(settings.Minimum)))
	}
	if settings.Decimal {
		maxLength += settings.Round
	}

	var blankLine int = 5
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
		printInputPicker(settings, manualInputBuffer.String(), maxLength)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			utils.Cleanup(4, !settings.DontCleanup)
			return -1, err
		}

		if arrow == "left" || arrow == "down" || ascii == 104 || ascii == 106 { // Left arrow, Down arrow, H or J
			temp := utils.StringToFloat(manualInputBuffer.String())
			temp -= settings.Increment
			temp = utils.RoundTo(temp, settings.Round)
			if temp <= settings.Maximum && temp >= settings.Minimum {
				manualInputBuffer.Reset()
				manualInputBuffer.WriteString(utils.FloatToString(temp))
			}
		} else if arrow == "right" || arrow == "up" || ascii == 108 || ascii == 107 { // Right arrow, Up arrow, L or K
			temp := utils.StringToFloat(manualInputBuffer.String())
			temp += settings.Increment
			temp = utils.RoundTo(temp, settings.Round)
			if temp <= settings.Maximum && temp >= settings.Minimum {
				manualInputBuffer.Reset()
				manualInputBuffer.WriteString(utils.FloatToString(temp))
			}

		} else if ascii >= 48 && ascii <= 57 { // Manually input a number
			if manualInputBuffer.String() == utils.FloatToString(settings.Default) {
				manualInputBuffer.Reset()
			}
			manualInputBuffer.WriteString(string(ascii))
			if utils.StringToFloat(manualInputBuffer.String()) > settings.Maximum || utils.StringToFloat(manualInputBuffer.String()) < settings.Minimum {
				if manualInputBuffer.Len() > 0 {
					manualInputBuffer.Truncate(manualInputBuffer.Len() - 1)
				}
			} else if utils.CountDigitsAfterDecimal(utils.StringToFloat(manualInputBuffer.String())) > int(settings.Round) {
				manualInputBuffer.Truncate(manualInputBuffer.Len() - 1)
			}
		} else if ascii == 46 { // Dot
			if !settings.Decimal {
				utils.Cleanup(4, !settings.DontCleanup)
				return -1, errors.New("Key not accepted")
			}
			if !strings.Contains(manualInputBuffer.String(), ".") {
				manualInputBuffer.WriteString(".")
			}
		} else if ascii == 127 { // Del
			if manualInputBuffer.Len() > 0 {
				manualInputBuffer.Truncate(manualInputBuffer.Len() - 1)
			}
		} else if ascii == 45 { // Dash
			if !strings.Contains(manualInputBuffer.String(), "-") && manualInputBuffer.Len() == 0 {
				manualInputBuffer.WriteString("-")
			}
		} else if ascii == 13 { // CR
			utils.Cleanup(4, !settings.DontCleanup)
			number = utils.StringToFloat(manualInputBuffer.String())
			number = utils.RoundTo(number, settings.Round)
			return number, nil
		} else {
			if settings.UnknownKeysErr {
				utils.Cleanup(4, !settings.DontCleanup)
				return -1, errors.New("Key not accepted")
			} else if ascii == 3 {
				utils.Cleanup(4, !settings.DontCleanup)
				return -1, errors.New("SIGINT")
			}
		}
	}
}
