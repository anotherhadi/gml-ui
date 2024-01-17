// https://github.com/anotherhadi/gml-ui
package loading

import (
	"fmt"
	"strings"
	"time"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/settings"
)

func Loading(loadingChan chan bool, customSettings ...settings.Settings) {

	settings := settings.GetSettings(customSettings)

	var frames []string = []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "}
	var fps int = 10
	var frame_index uint8 = 0

	fmt.Print(strings.Repeat("\n", settings.TopPadding))
	for {
		select {
		case <-loadingChan:
			fmt.Print("\r")
			ansi.LineClear()
			return

		default:
			fmt.Print("\r")
			fmt.Print(strings.Repeat(" ", int(settings.LeftPadding)))

			fmt.Print(ansi.FgRgbSettings(settings.AccentColor))
			fmt.Print(frames[frame_index])
			frame_index++
			if frame_index >= uint8(len(frames)) {
				frame_index = 0
			}
			fmt.Print(ansi.FgRgbSettings(settings.TextColor))
			fmt.Print("Loading...")

			fmt.Print(ansi.Reset)
			time.Sleep(time.Second / time.Duration(fps))
		}
	}
}
