// https://github.com/anotherhadi/gml-ui
package getchar

import (
	"github.com/pkg/term"
)

const (
	escape     = 27
	leftSquare = 91
	upArrow    = 65
	downArrow  = 66
	rightArrow = 67
	leftArrow  = 68
)

// GetChar returns either an ASCII code or an arrow ("up", "down", "right", "left").
func GetChar() (ascii uint8, arrow string, err error) {
	t, err := term.Open("/dev/tty")
	if err != nil {
		return
	}
	defer t.Close()

	err = term.RawMode(t)
	if err != nil {
		return
	}
	defer func() {
		err = t.Restore()
	}()

	bytes := make([]byte, 3)
	numRead, err := t.Read(bytes)
	if err != nil {
		return
	}

	if numRead == 3 && bytes[0] == escape && bytes[1] == leftSquare {
		// Three-character control sequence, beginning with "ESC-[".
		// Since there are no ASCII codes for arrow keys, returning strings
		switch bytes[2] {
		case upArrow:
			arrow = "up"
		case downArrow:
			arrow = "down"
		case rightArrow:
			arrow = "right"
		case leftArrow:
			arrow = "left"
		}
	} else if numRead == 1 {
		ascii = uint8(bytes[0])
	}

	return
}
