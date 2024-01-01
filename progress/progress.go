package progress

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/utils"
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

func ProgressBar(percentageChan chan int, customSettings ...Settings) {

	var settings Settings

	if len(customSettings) > 0 {
		settings = combineSettings(customSettings[0])
	} else {
		settings = getDefaultSettings()
	}

	var red int
	var redScale int = (int(settings.ProgressSecondForeground.Red) - int(settings.ProgressForeground.Red)) / int(settings.Width)
	var green int
	var greenScale int = (int(settings.ProgressSecondForeground.Green) - int(settings.ProgressForeground.Green)) / int(settings.Width)
	var blue int
	var blueScale int = (int(settings.ProgressSecondForeground.Blue) - int(settings.ProgressForeground.Blue)) / int(settings.Width)

	fmt.Print("\n")
	for percentage := range percentageChan {
		percentScaled := percentage * int(settings.Width) / 100
		fmt.Print("\r")
		fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))

		for i := 0; i < percentScaled; i++ {
			red = adjustColorValue(int(settings.ProgressForeground.Red), redScale, i)
			green = adjustColorValue(int(settings.ProgressForeground.Green), greenScale, i)
			blue = adjustColorValue(int(settings.ProgressForeground.Blue), blueScale, i)
			fmt.Print(ansi.BgRgb(uint8(red), uint8(green), uint8(blue)))
			fmt.Print(" ")
		}

		fmt.Print(ansi.BgRgb(settings.EmptyForeground.Red, settings.EmptyForeground.Green, settings.EmptyForeground.Blue))
		fmt.Print(utils.Repeat(" ", int(settings.Width)-percentScaled))

		fmt.Print(ansi.Reset)
		fmt.Print(ansi.FgRgb(settings.PercentageForeground.Red, settings.PercentageForeground.Green, settings.PercentageForeground.Blue))
		fmt.Print(" ", percentage, "% ")
		fmt.Print(ansi.Reset)
		if percentage == 100 {
			if !settings.DontCleanup {
				fmt.Print("\r")
				fmt.Print(utils.Repeat(" ", int(settings.Width)+int(settings.LeftPadding)+5))
				fmt.Print("\r")
				ansi.CursorUp(1)
			} else {
				fmt.Print("\n\n")
			}
		}
	}
}
