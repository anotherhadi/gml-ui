// https://github.com/anotherhadi/gml-ui
package ansi

import (
	"fmt"

	"github.com/anotherhadi/gml-ui/settings"
)

// Escape code
const ESC = "\033"

//
//        Text Styling
//

// Ex: fmt.Println(ansi.Underline + "Underlined text")
const (
	Reset      = ESC + "[0m"
	Bold       = ESC + "[1m"
	Faint      = ESC + "[2m"
	Italic     = ESC + "[3m"
	Underline  = ESC + "[4m"
	Blink      = ESC + "[5m"
	Reverse    = ESC + "[7m"
	CrossedOut = ESC + "[9m"
)

//
//        Foreground color
//

// Ex: fmt.Println(ansi.Blue + "Blue text")
const (
	Black   = ESC + "[30m"
	Red     = ESC + "[31m"
	Green   = ESC + "[32m"
	Yellow  = ESC + "[33m"
	Blue    = ESC + "[34m"
	Magenta = ESC + "[35m"
	Cyan    = ESC + "[36m"
	White   = ESC + "[37m"

	BrightBlack   = ESC + "[90m"
	BrightRed     = ESC + "[91m"
	BrightGreen   = ESC + "[92m"
	BrightYellow  = ESC + "[93m"
	BrightBlue    = ESC + "[94m"
	BrightMagenta = ESC + "[95m"
	BrightCyan    = ESC + "[96m"
	BrightWhite   = ESC + "[97m"
)

// 256 Foreground Color
// Ex: fmt.Println(ansi.Fg256(27) + "Blue text")
func Fg256(color int) string {
	return fmt.Sprintf("%s[38;5;%dm", ESC, color)
}

// RGB Foreground Color
// Ex: fmt.Println(ansi.FgRgb(30, 30, 255) + "Blue text")
func FgRgb(red, green, blue uint8) string {
	return fmt.Sprintf("%s[38;2;%d;%d;%dm", ESC, red, green, blue)
}

func FgRgbSettings(s settings.Color) string {
	return fmt.Sprintf("%s[38;2;%d;%d;%dm", ESC, s.Red, s.Green, s.Blue)
}

//
//        Background color
//

// Ex: fmt.Println(ansi.BgBlue + "Text with a blue background")
const (
	BgBlack   = ESC + "[40m"
	BgRed     = ESC + "[41m"
	BgGreen   = ESC + "[42m"
	BgYellow  = ESC + "[43m"
	BgBlue    = ESC + "[44m"
	BgMagenta = ESC + "[45m"
	BgCyan    = ESC + "[46m"
	BgWhite   = ESC + "[47m"

	BgBrightBlack   = ESC + "[100m"
	BgBrightRed     = ESC + "[101m"
	BgBrightGreen   = ESC + "[102m"
	BgBrightYellow  = ESC + "[103m"
	BgBrightBlue    = ESC + "[104m"
	BgBrightMagenta = ESC + "[105m"
	BgBrightCyan    = ESC + "[106m"
	BgBrightWhite   = ESC + "[107m"
)

// 256 Background Color
// Ex: fmt.Println(ansi.Bg256(27) + "Text with a blue background")
func Bg256(color int) string {
	return fmt.Sprintf("%s[48;5;%dm", ESC, color)
}

// RGB Background Color
// Ex: fmt.Println(ansi.BgRgb(30, 30, 255) + "Text with a blue background")
func BgRgb(red, green, blue uint8) string {
	return fmt.Sprintf("%s[48;2;%d;%d;%dm", ESC, red, green, blue)
}

func BgRgbSettings(s settings.Color) string {
	return fmt.Sprintf("%s[48;2;%d;%d;%dm", ESC, s.Red, s.Green, s.Blue)
}

//
//        Cursor Movement
//

// Move cursor to {line}, {col}
// Ex: fmt.Print(ansi.CursorMove(10,20))
func CursorMove(line, col int) string {
	return fmt.Sprintf("%s[%d;%dH", ESC, line, col)
}

// Move cursor {line} up
// Ex: fmt.Print(ansi.CursorUp(10))
func CursorUpN(line int) string {
	return fmt.Sprintf("%s[%dA", ESC, line)
}

// Move cursor {line} down
// Ex: fmt.Print(ansi.CursorDown(10))
func CursorDownN(line int) string {
	return fmt.Sprintf("%s[%dB", ESC, line)
}

// Move cursor {col} right
// Ex: fmt.Print(ansi.CursorRight(10))
func CursorRightN(col int) string {
	return fmt.Sprintf("%s[%dC", ESC, col)
}

// Move cursor {col} left
// Ex: fmt.Print(ansi.CursorLeft(10))
func CursorLeftN(col int) string {
	return fmt.Sprintf("%s[%dD", ESC, col)
}

// Move cursor to {col}
// Ex: fmt.Print(ansi.CursorCol(10))
func CursorCol(col int) string {
	return fmt.Sprintf("%s[%dG", ESC, col)
}

// Move cursor to Home (0,0)
// Ex: fmt.Print(ansi.CursorHome())
func CursorHome() string {
	return fmt.Sprintf("%s[H", ESC)
}

// Save cursor position
func CursorSave() string {
	return fmt.Sprintf("%s[s", ESC)
}

// Restore cursor position
func CursorRestore() string {
	return fmt.Sprintf("%s[u", ESC)
}

// Make cursor visible
func CursorVisible() string {
	return fmt.Sprintf("%s[?25h", ESC)
}

// Make cursor invisible
func CursorInvisible() string {
	return fmt.Sprintf("%s[?25l", ESC)
}

//
//        Alternative buffer
//

// Enables the alternative buffer
func AlternativeBufferEnable() string {
	return fmt.Sprintf("%s[?1049h", ESC)
}

// Disables the alternative buffer
func AlternativeBufferDisable() string {
	return fmt.Sprintf("%s[?1049l", ESC)
}

//
//        Screen
//

// Save the screen
func ScreenSave() string {
	return fmt.Sprintf("%s[?47h", ESC)
}

// Restore the screen
func ScreenRestore() string {
	return fmt.Sprintf("%s[?47l", ESC)
}

// Clear the entire screen
func ScreenClear() string {
	return fmt.Sprintf("%s[2J", ESC)
}

// Clear screen up
func ScreenClearUp() string {
	return fmt.Sprintf("%s[1J", ESC)
}

// Clear screen down
func ScreenClearDown() string {
	return fmt.Sprintf("%s[0J", ESC)
}

// Clear screen to the end
func ScreenClearEnd() string {
	return fmt.Sprintf("%s[J", ESC)
}

//
//        Clear Lines
//

// Clear the entire line
func LineClear() string {
	return fmt.Sprintf("%s[2K", ESC)
}

// Clear start of line to the cursor
func LineClearStart() string {
	return fmt.Sprintf("%s[1K", ESC)
}

// Clear from cursor to end of line
func LineClearEnd() string {
	return fmt.Sprintf("%s[K", ESC)
}
