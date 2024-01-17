// https://github.com/anotherhadi/gml-ui
package progress

import (
	"fmt"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/settings"
)

func adjustColorValue(base, scale, value int) int {
	result := base + scale*value
	if result < 0 {
		return 0
	} else if result > 255 {
		return 255
	}
	return result
}

func ProgressBar(percentageChan chan int, customSettings ...settings.Settings) {

	settings := settings.GetSettings(customSettings)

	var red int
	var redScale int = (int(settings.SecondaryColor.Red) - int(settings.AccentColor.Red)) / int(settings.MaxCols)
	var green int
	var greenScale int = (int(settings.SecondaryColor.Green) - int(settings.AccentColor.Green)) / int(settings.MaxCols)
	var blue int
	var blueScale int = (int(settings.SecondaryColor.Blue) - int(settings.AccentColor.Blue)) / int(settings.MaxCols)

	fmt.Print("\n")
	for percentage := range percentageChan {
		percentScaled := percentage * int(settings.MaxCols) / 100
		fmt.Print("\r")
		fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))

		for i := 0; i < percentScaled; i++ {
			red = adjustColorValue(int(settings.AccentColor.Red), redScale, i)
			green = adjustColorValue(int(settings.AccentColor.Green), greenScale, i)
			blue = adjustColorValue(int(settings.AccentColor.Blue), blueScale, i)
			fmt.Print(ansi.BgRgb(uint8(red), uint8(green), uint8(blue)))
			fmt.Print(" ")
		}

		fmt.Print(ansi.BgRgbSettings(settings.TextBackgroundColor))
		fmt.Print(strings.Repeat(" ", int(settings.MaxCols)-percentScaled))

		fmt.Print(ansi.Reset)
		fmt.Print(ansi.FgRgbSettings(settings.TextColor))
		fmt.Print(" ", percentage, "% ")
		fmt.Print(ansi.Reset)
		if percentage == 100 {
			fmt.Print("\n\n")
		}
	}
}
