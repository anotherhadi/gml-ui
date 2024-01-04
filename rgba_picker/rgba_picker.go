package rgbapicker

import (
	"errors"
	"fmt"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/asciimoji"
	"github.com/anotherhadi/gml-ui/getchar"
	"github.com/anotherhadi/gml-ui/utils"
)

type boxStyle struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Verticaly   string
	Horizontaly string
}

func getBoxStyle(s string) boxStyle {
	styles := map[string]boxStyle{
		"-": {
			TopLeft:     "┌",
			TopRight:    "┐",
			BottomLeft:  "└",
			BottomRight: "┘",
			Verticaly:   "│",
			Horizontaly: "─",
		},
		"=": {
			TopLeft:     "╔",
			TopRight:    "╗",
			BottomLeft:  "╚",
			BottomRight: "╝",
			Verticaly:   "║",
			Horizontaly: "═",
		},
		"none": {
			TopLeft:     " ",
			TopRight:    " ",
			BottomLeft:  " ",
			BottomRight: " ",
			Verticaly:   " ",
			Horizontaly: " ",
		},
	}
	return styles[s]
}

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

func printRgbaPicker(settings Settings, rgba [4]int, selected int) {
	ansi.ClearScreenEnd()
	length := 5*4 + 3
	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
	boxStyle := getBoxStyle(settings.BoxStyle)

	fmt.Print("   ")
	for i := 0; i < 4; i++ {
		if i == selected {
			fmt.Print(ansi.FgRgb(settings.InputForeground.Red, settings.InputForeground.Green, settings.InputForeground.Blue))
		}
		fmt.Print(asciimoji.Up)
		fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
		fmt.Print(utils.Repeat(" ", 5))
	}
	fmt.Print("\n")
	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(boxStyle.TopLeft)
	fmt.Print(utils.Repeat(boxStyle.Horizontaly, length))
	fmt.Print(boxStyle.TopRight)
	fmt.Print("   ")
	fmt.Print(ansi.BgRgb(uint8(rgba[0]), uint8(rgba[1]), uint8(rgba[2])))
	fmt.Print("   ")
	fmt.Print(ansi.Reset)
	fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
	fmt.Print("\n")

	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(boxStyle.Verticaly)
	for i := 0; i < 4; i++ {
		fmt.Print(" ")
		if i == selected {
			fmt.Print(ansi.FgRgb(settings.InputForeground.Red, settings.InputForeground.Green, settings.InputForeground.Blue))
		}
		fmt.Printf("%3d", rgba[i])
		fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
		fmt.Print(" ")
		fmt.Print(boxStyle.Verticaly)
	}
	fmt.Print("   ")
	fmt.Print(ansi.BgRgb(uint8(rgba[0]), uint8(rgba[1]), uint8(rgba[2])))
	fmt.Print("   ")
	fmt.Print(ansi.Reset)
	fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))

	fmt.Print("\n")
	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(boxStyle.BottomLeft)
	fmt.Print(utils.Repeat(boxStyle.Horizontaly, length))
	fmt.Print(boxStyle.BottomRight)
	fmt.Print("   ")
	fmt.Print(ansi.BgRgb(uint8(rgba[0]), uint8(rgba[1]), uint8(rgba[2])))
	fmt.Print("   ")
	fmt.Print(ansi.Reset)
	fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
	fmt.Print("\n")
	fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print("   ")
	for i := 0; i < 4; i++ {
		if i == selected {
			fmt.Print(ansi.FgRgb(settings.InputForeground.Red, settings.InputForeground.Green, settings.InputForeground.Blue))
		}
		fmt.Print(asciimoji.Down)
		fmt.Print(ansi.FgRgb(settings.BorderForeground.Red, settings.BorderForeground.Green, settings.BorderForeground.Blue))
		fmt.Print(utils.Repeat(" ", 5))
	}
	fmt.Print("\n")

	fmt.Print(ansi.Reset)
	ansi.CursorUp(5)
}

func RgbaPicker(customSettings ...Settings) (rgba [4]int, err error) {

	var settings Settings

	if len(customSettings) > 0 {
		settings = combineSettings(customSettings[0])
	} else {
		settings = getDefaultSettings()
	}

	rgba = settings.Default
	var selected int = 0

	var blankLine int = 6
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
		printRgbaPicker(settings, rgba, selected)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			utils.Cleanup(6, !settings.DontCleanup)
			return [4]int{}, err
		}

		if arrow == "left" || ascii == 104 { // Left arrow, Down arrow, H or J
			if selected != 0 {
				selected--
			}
		} else if arrow == "right" || ascii == 108 { // Right arrow, Up arrow, L or K
			if selected != 3 {
				selected++
			}
		} else if arrow == "down" || ascii == 106 {
			if rgba[selected] != 0 {
				rgba[selected]--
			}
		} else if arrow == "up" || ascii == 107 {
			if rgba[selected] != 255 {
				rgba[selected]++
			}
		} else if ascii >= 48 && ascii <= 57 { // Manually input a number
			number := int(ascii) - 48
			if rgba[selected]*10+number <= 255 {
				rgba[selected] = rgba[selected]*10 + number
			}
		} else if ascii == 127 { // Del
			rgba[selected] = int(rgba[selected] / 10)
		} else if ascii == 13 { // CR
			utils.Cleanup(6, !settings.DontCleanup)
			return rgba, nil
		} else {
			if settings.UnknownKeysErr {
				utils.Cleanup(6, !settings.DontCleanup)
				return [4]int{}, errors.New("Key not accepted")
			} else if ascii == 3 {
				utils.Cleanup(6, !settings.DontCleanup)
				return [4]int{}, errors.New("SIGINT")
			}
		}
	}
}
