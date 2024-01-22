// https://github.com/anotherhadi/gml-ui
package number_picker

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/getchar"
	"github.com/anotherhadi/gml-ui/settings"
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

func printInputPicker(settings settings.Settings, number string, maxLength int) {

	boxStyle := getBoxStyle(settings.Style)
	fmt.Print(ansi.CursorUpN(settings.TopPadding + 3 + settings.BottomPadding))
	fmt.Print(ansi.ScreenClearEnd())

	fmt.Print(strings.Repeat("\n", settings.TopPadding))

	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))

	fmt.Print("  " + boxStyle.TopLeft)
	fmt.Print(strings.Repeat(boxStyle.Horizontaly, int(maxLength)+2))
	fmt.Print(boxStyle.TopRight + "\n")

	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
	if stringToFloat(number) > settings.Minimum {
		fmt.Print("< ")
	} else {
		fmt.Print("  ")
	}
	fmt.Print(boxStyle.Verticaly + " ")

	fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
	fmt.Print(strings.Repeat(" ", int(maxLength)-len(number)))
	fmt.Printf("%s", number)

	fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
	fmt.Print(" " + boxStyle.Verticaly)
	if stringToFloat(number) < settings.Maximum {
		fmt.Print(" > ")
	} else {
		fmt.Print("  ")
	}

	fmt.Print("\n")
	fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))

	fmt.Print(ansi.FgRgbSettings(settings.SecondaryColor))
	fmt.Print("  " + boxStyle.BottomLeft)
	fmt.Print(strings.Repeat(boxStyle.Horizontaly, int(maxLength)+2))
	fmt.Print(boxStyle.BottomRight + "\n")
	fmt.Print(ansi.Reset)
	fmt.Print(strings.Repeat("\n", settings.BottomPadding))
}

func stringToFloat(str string) float64 {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return result
}

func floatToString(n float64) string {
	return strconv.FormatFloat(n, 'f', -1, 64)
}

func roundTo(n float64, decimals int) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func countDigitsAfterDecimal(num float64) int {
	strNum := fmt.Sprintf("%g", num)
	decimalPos := strings.Index(strNum, ".")
	if decimalPos != -1 {
		return len(strNum) - decimalPos - 1
	}
	return 0
}

func NumberPicker(customSettings ...settings.Settings) (number float64, err error) {

	settings := settings.GetSettings(customSettings)

	var manualInputBuffer bytes.Buffer
	if settings.Decimal == 0 {
		manualInputBuffer.WriteString(floatToString(float64(settings.DefaultInt)))
	} else {
		manualInputBuffer.WriteString(floatToString(settings.DefaultFloat))
	}
	var maxLength int = len(floatToString(settings.Maximum))
	if maxLength < len(floatToString(settings.Minimum)) {
		maxLength = len(floatToString(settings.Minimum))
	}
	if settings.Decimal != 0 {
		maxLength += settings.Decimal
	}

	fmt.Print(ansi.CursorInvisible())
	fmt.Print(strings.Repeat("\n", settings.TopPadding+3+settings.BottomPadding))

	for {
		printInputPicker(settings, manualInputBuffer.String(), maxLength)

		ascii, arrow, err := getchar.GetChar()
		if err != nil {
			fmt.Print(ansi.CursorVisible())
			return -1, err
		}

		if arrow == "left" || arrow == "down" || ascii == 104 || ascii == 106 { // Left arrow, Down arrow, H or J
			temp := stringToFloat(manualInputBuffer.String())
			temp -= settings.Increment
			temp = roundTo(temp, settings.Decimal)
			if temp <= settings.Maximum && temp >= settings.Minimum {
				manualInputBuffer.Reset()
				manualInputBuffer.WriteString(floatToString(temp))
			}
		} else if arrow == "right" || arrow == "up" || ascii == 108 || ascii == 107 { // Right arrow, Up arrow, L or K
			temp := stringToFloat(manualInputBuffer.String())
			temp += settings.Increment
			temp = roundTo(temp, settings.Decimal)
			if temp <= settings.Maximum && temp >= settings.Minimum {
				manualInputBuffer.Reset()
				manualInputBuffer.WriteString(floatToString(temp))
			}

		} else if ascii >= 48 && ascii <= 57 { // Manually input a number
			manualInputBuffer.WriteString(string(ascii))
			if stringToFloat(manualInputBuffer.String()) > settings.Maximum || stringToFloat(manualInputBuffer.String()) < settings.Minimum {
				if manualInputBuffer.Len() > 0 {
					manualInputBuffer.Truncate(manualInputBuffer.Len() - 1)
				}
			} else if countDigitsAfterDecimal(stringToFloat(manualInputBuffer.String())) > int(settings.Decimal) {
				manualInputBuffer.Truncate(manualInputBuffer.Len() - 1)
			}
			if strings.HasPrefix(manualInputBuffer.String(), "0") && len(manualInputBuffer.String()) > 1 {
				temp := manualInputBuffer.String()
				manualInputBuffer.Reset()
				manualInputBuffer.WriteString(temp[1:])
			}
		} else if ascii == 46 { // Dot
			if settings.Decimal == 0 {
				if settings.ExitOnUnknownKey {
					fmt.Print(ansi.CursorVisible())
					return -1, errors.New("Key not accepted")
				} else {
					continue
				}
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
			number = stringToFloat(manualInputBuffer.String())
			number = roundTo(number, settings.Decimal)
			fmt.Print(ansi.CursorVisible())
			return number, nil
		} else {
			if settings.ExitOnUnknownKey {
				fmt.Print(ansi.CursorVisible())
				return -1, errors.New("Key not accepted")
			} else if ascii == 3 {
				fmt.Print(ansi.CursorVisible())
				return -1, errors.New("SIGINT")
			}
		}
	}
}
