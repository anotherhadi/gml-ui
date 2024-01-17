// https://github.com/anotherhadi/gml-ui
package getsize

import (
	"golang.org/x/term"
	"os"
)

func GetSize() (cols, rows int, err error) {
	var fd uintptr = os.Stdin.Fd()
	cols, rows, err = term.GetSize(int(fd))
	return
}
