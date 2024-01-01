package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
)

func Repeat(s string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

func SplitPrompt(prompt string, maxCols int) []string {
	var result []string

	for i := 0; i < len(prompt); i += maxCols {
		endIndex := i + maxCols
		if endIndex > len(prompt) {
			endIndex = len(prompt)
		}

		result = append(result, prompt[i:endIndex])
	}

	return result
}

func Cleanup(line uint8, cleanup bool) {
	ansi.CursorVisible()
	fmt.Print(Repeat("\n", int(line)))
	if cleanup {
		ansi.CursorRestore()
		ansi.ClearScreenEnd()
	}
}

func FilterStringsByPrefix(original []string, prefix string, caseSensitive bool) []string {
	if prefix == "" {
		return original
	}

	var filtered []string

	if caseSensitive {
		for _, str := range original {
			if strings.HasPrefix(str, prefix) {
				filtered = append(filtered, str)
			}
		}
	} else {
		var lowerPrefix string = strings.ToLower(prefix)
		var lowerStr string
		for _, str := range original {
			lowerStr = strings.ToLower(str)
			if strings.HasPrefix(lowerStr, lowerPrefix) {
				filtered = append(filtered, str)
			}
		}
	}

	return filtered
}

func CountDigits(n int) int {
	var str string = strconv.Itoa(n)
	return len(str)
}

func StringToFloat(str string) float64 {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return result
}

func FloatToString(n float64) string {
	return strconv.FormatFloat(n, 'f', -1, 64)
}

func RoundTo(n float64, decimals uint8) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func CountDigitsAfterDecimal(num float64) int {
	strNum := fmt.Sprintf("%g", num)
	decimalPos := strings.Index(strNum, ".")
	if decimalPos != -1 {
		return len(strNum) - decimalPos - 1
	}
	return 0
}
