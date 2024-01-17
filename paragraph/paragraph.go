// https://github.com/anotherhadi/gml-ui
package paragraph

import (
	"fmt"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/settings"
)

func splitPrompt(prompt string, maxCols int) []string {
	var result []string

	words := strings.Fields(prompt)

	var currentLine string
	for _, word := range words {
		if len(currentLine)+len(word)+1 <= maxCols {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		} else {
			result = append(result, currentLine)
			currentLine = word
		}
	}

	if currentLine != "" {
		result = append(result, currentLine)
	}

	return result
}

func printParagraph(str string, settings settings.Settings) {
	fmt.Print(strings.Repeat("\n", int(settings.TopPadding)))
	fmt.Print(ansi.FgRgbSettings(settings.TextColor))
	var splitedPrompt []string = splitPrompt(str, settings.MaxCols)

	for _, line := range splitedPrompt {
		fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))
		fmt.Println(line)
	}
	fmt.Print(ansi.Reset)

	fmt.Print(strings.Repeat("\n", int(settings.BottomPadding)))
}

func Paragraph(str string, customSettings ...settings.Settings) {

	settings := settings.GetSettings(customSettings)

	printParagraph(str, settings)
}
