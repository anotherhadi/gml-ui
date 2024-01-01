package ansi

import (
	"fmt"
)

const ESC = "\033"

// Text Styling
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

// Text Color
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
func Fg256(color uint8) string {
	return fmt.Sprintf("%s[38;5;%dm", ESC, color)
}

// RGB Foreground Color
// Ex: fmt.Println(ansi.FgRgb(30, 30, 255) + "Blue text")
func FgRgb(red, green, blue uint8) string {
	return fmt.Sprintf("%s[38;2;%d;%d;%dm", ESC, red, green, blue)
}

// Background Color
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
func Bg256(color uint8) string {
	return fmt.Sprintf("%s[48;5;%dm", ESC, color)
}

// RGB Background Color
// Ex: fmt.Println(ansi.BgRgb(30, 30, 255) + "Text with a blue background")
func BgRgb(red, green, blue uint8) string {
	return fmt.Sprintf("%s[48;2;%d;%d;%dm", ESC, red, green, blue)
}

//// Cursor Movement

// Move cursor to {line}, {col}
// Ex: ansi.CursorMove(10,20)
func CursorMove(line, col uint8) {
	fmt.Printf("%s[%d;%dH", ESC, line, col)
}

// Move cursor {line} up
// Ex: ansi.CursorUp(10)
func CursorUp(line uint8) {
	fmt.Printf("%s[%dA", ESC, line)
}

// Move cursor {line} down
// Ex: ansi.CursorDown(10)
func CursorDown(line uint8) {
	fmt.Printf("%s[%dB", ESC, line)
}

// Move cursor {col} right
// Ex: ansi.CursorRight(10)
func CursorRight(col uint8) {
	fmt.Printf("%s[%dC", ESC, col)
}

// Move cursor {col} left
// Ex: ansi.CursorLeft(10)
func CursorLeft(col uint8) {
	fmt.Printf("%s[%dD", ESC, col)
}

// Move cursor to {col}
// Ex: ansi.CursorCol(10)
func CursorCol(col uint8) {
	fmt.Printf("%s[%dG", ESC, col)
}

// Move cursor to Home (0,0)
// Ex: ansi.CursorHome()
func CursorHome() {
	fmt.Printf("%s[H", ESC)
}

// Save cursor position
func CursorSave() {
	fmt.Printf("%s[s", ESC)
}

// Restore cursor position
func CursorRestore() {
	fmt.Printf("%s[u", ESC)
}

// Make cursor invisible
func CursorInvisible() {
	fmt.Printf("%s[?25l", ESC)
}

// Make cursor visible
func CursorVisible() {
	fmt.Printf("%s[?25h", ESC)
}

// Save the screen
func ScreenSave() {
	fmt.Printf("%s[?47h", ESC)
}

// Restore the screen
func ScreenRestore() {
	fmt.Printf("%s[?47l", ESC)
}

// Enables the alternative buffer
func EnableAlternativeBuffer() {
	fmt.Printf("%s[?1049h", ESC)
}

// Disables the alternative buffer
func DisableAlternativeBuffer() {
	fmt.Printf("%s[?1049h", ESC)
}

//// Clear/Erase

// Clear the full screen
func ClearScreen() {
	fmt.Printf("%s[2J", ESC)
}

// Clear screen up
func ClearScreenUp() {
	fmt.Printf("%s[1J", ESC)
}

// Clear screen down
func ClearScreenDown() {
	fmt.Printf("%s[0J", ESC)
}

// Clear screen to the end
func ClearScreenEnd() {
	fmt.Printf("%s[J", ESC)
}

// Clear the entire line
func ClearLine() {
	fmt.Printf("%s[2K", ESC)
}

// Clear start of line to the cursor
func ClearLineStart() {
	fmt.Printf("%s[1K", ESC)
}

// Clear from cursor to end of line
func ClearLineEnd() {
	fmt.Printf("%s[K", ESC)
}
