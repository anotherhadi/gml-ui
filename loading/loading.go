package loading

import (
	"fmt"
	"time"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/utils"
)

func Loading(loadingChan chan bool, customSettings ...Settings) {

	var settings Settings

	if len(customSettings) > 0 {
		settings = combineSettings(customSettings[0])
	} else {
		settings = getDefaultSettings()
	}
	var frames []string = []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "}
	var frame_index uint8 = 0

	fmt.Print("\n")
	for {
		select {
		case <-loadingChan:
			if !settings.DontCleanup {
				fmt.Print("\r")
				fmt.Print(utils.Repeat(" ", 2+len(settings.Message)+int(settings.LeftPadding)))
				fmt.Print("\r")
				ansi.CursorUp(1)
			} else {
				fmt.Print("\n\n")
			}
			return

		default:
			fmt.Print("\r")
			fmt.Print(utils.Repeat(" ", int(settings.LeftPadding)))

			fmt.Print(ansi.FgRgb(settings.LoadingForeground.Red, settings.LoadingForeground.Green, settings.LoadingForeground.Blue))
			fmt.Print(frames[frame_index])
			frame_index++
			if frame_index >= uint8(len(frames)) {
				frame_index = 0
			}
			fmt.Print(ansi.FgRgb(settings.MessageForeground.Red, settings.MessageForeground.Green, settings.MessageForeground.Blue))
			fmt.Print(settings.Message)

			fmt.Print(ansi.Reset)
			time.Sleep(time.Second / time.Duration(settings.FPS))
		}
	}
}
