package paragraph

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/utils"
)

func printParagraph(str string, settings Settings) {
	fmt.Print("\n")
	fmt.Print(ansi.FgRgb(settings.Foreground.Red, settings.Foreground.Green, settings.Foreground.Blue))
	if settings.Background != (RGBColor{}) {
		fmt.Print(ansi.BgRgb(settings.Background.Red, settings.Background.Green, settings.Background.Blue))
	}
	var splitedPrompt []string = utils.SplitPrompt(str, int(settings.MaxCols))
	for _, line := range splitedPrompt {
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))
		fmt.Println(line)
	}
	fmt.Print(ansi.Reset)
	fmt.Print("\n")
}

func Paragraph(str string, customSettings ...Settings) {

	var settings Settings

	if len(customSettings) > 0 {
		settings = combineSettings(customSettings[0])
	} else {
		settings = getDefaultSettings()
	}

	printParagraph(str, settings)
}
