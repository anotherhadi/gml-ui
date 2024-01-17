// https://github.com/anotherhadi/gml-ui
package asciitext

import (
	"strings"

	"github.com/anotherhadi/gml-ui/settings"
)

func getLetter(letter rune) []string {
	letters := make(map[string][]string)
	letters["a"] = []string{
		"        ",
		"   __ _ ",
		"  / _` |",
		" | (_| |",
		"  \\__,_|",
		"        ",
	}
	letters["b"] = []string{
		"  _     ",
		" | |__  ",
		" | '_ \\ ",
		" | |_) |",
		" |_.__/ ",
		"        ",
	}
	letters["c"] = []string{
		"       ",
		"   ___ ",
		"  / __|",
		" | (__ ",
		"  \\___|",
		"       ",
	}
	letters["d"] = []string{
		"      _ ",
		"   __| |",
		"  / _` |",
		" | (_| |",
		"  \\__,_|",
		"        ",
	}
	letters["e"] = []string{
		"       ",
		"   ___ ",
		"  / _ \\",
		" |  __/",
		"  \\___|",
		"       ",
	}
	letters["f"] = []string{
		"   __ ",
		"  / _|",
		" | |_ ",
		" |  _|",
		" |_|  ",
		"      ",
	}
	letters["g"] = []string{
		"        ",
		"   __ _ ",
		"  / _` |",
		" | (_| |",
		"  \\__, |",
		"  |___/ ",
	}
	letters["h"] = []string{
		"  _     ",
		" | |__  ",
		" | '_ \\ ",
		" | | | |",
		" |_| |_|",
		"        ",
	}
	letters["i"] = []string{
		"  _ ",
		" (_)",
		" | |",
		" | |",
		" |_|",
		"    ",
	}
	letters["j"] = []string{
		"    _ ",
		"   (_)",
		"   | |",
		"   | |",
		"  _/ |",
		" |__/ ",
	}
	letters["k"] = []string{
		"  _    ",
		" | | __",
		" | |/ /",
		" |   < ",
		" |_|\\_\\",
		"       ",
	}
	letters["l"] = []string{
		"  _ ",
		" | |",
		" | |",
		" | |",
		" |_|",
		"    ",
	}
	letters["m"] = []string{
		"            ",
		"  _ __ ___  ",
		" | '_ ` _ \\ ",
		" | | | | | |",
		" |_| |_| |_|",
		"            ",
	}
	letters["n"] = []string{
		"        ",
		"  _ __  ",
		" | '_ \\ ",
		" | | | |",
		" |_| |_|",
		"        ",
	}
	letters["o"] = []string{
		"        ",
		"   ___  ",
		"  / _ \\ ",
		" | (_) |",
		"  \\___/ ",
		"        ",
	}
	letters["p"] = []string{
		"        ",
		"  _ __  ",
		" | '_ \\ ",
		" | |_) |",
		" | .__/ ",
		" |_|    ",
	}
	letters["q"] = []string{
		"        ",
		"   __ _ ",
		"  / _` |",
		" | (_| |",
		"  \\__, |",
		"     |_|",
	}
	letters["r"] = []string{
		"       ",
		"  _ __ ",
		" | '__|",
		" | |   ",
		" |_|   ",
		"       ",
	}
	letters["s"] = []string{
		"      ",
		"  ___ ",
		" / __|",
		" \\__ \\",
		" |___/",
		"      ",
	}
	letters["t"] = []string{
		"  _   ",
		" | |_ ",
		" | __|",
		" | |_ ",
		"  \\__|",
		"      ",
	}
	letters["u"] = []string{
		"        ",
		"  _   _ ",
		" | | | |",
		" | |_| |",
		"  \\__,_|",
		"        ",
	}
	letters["v"] = []string{
		"        ",
		" __   __",
		" \\ \\ / /",
		"  \\ V / ",
		"   \\_/  ",
		"        ",
	}
	letters["w"] = []string{
		"           ",
		" __      __",
		" \\ \\ /\\ / /",
		"  \\ V  V / ",
		"   \\_/\\_/  ",
		"           ",
	}
	letters["x"] = []string{
		"       ",
		" __  __",
		" \\ \\/ /",
		"  >  < ",
		" /_/\\_\\",
		"       ",
	}
	letters["y"] = []string{
		"        ",
		"  _   _ ",
		" | | | |",
		" | |_| |",
		"  \\__, |",
		"  |___/ ",
	}
	letters["z"] = []string{
		"      ",
		"  ____",
		" |_  /",
		"  / / ",
		" /___|",
		"      ",
	}
	letters["A"] = []string{
		"     _    ",
		"    / \\   ",
		"   / _ \\  ",
		"  / ___ \\ ",
		" /_/   \\_\\",
		"          ",
	}
	letters["B"] = []string{
		"  ____  ",
		" | __ ) ",
		" |  _ \\ ",
		" | |_) |",
		" |____/ ",
		"        ",
	}
	letters["C"] = []string{
		"   ____ ",
		"  / ___|",
		" | |    ",
		" | |___ ",
		"  \\____|",
		"        ",
	}
	letters["D"] = []string{
		"  ____  ",
		" |  _ \\ ",
		" | | | |",
		" | |_| |",
		" |____/ ",
		"        ",
	}
	letters["E"] = []string{
		"  _____ ",
		" | ____|",
		" |  _|  ",
		" | |___ ",
		" |_____|",
		"        ",
	}
	letters["F"] = []string{
		"  _____ ",
		" |  ___|",
		" | |_   ",
		" |  _|  ",
		" |_|    ",
		"        ",
	}
	letters["G"] = []string{
		"   ____ ",
		"  / ___|",
		" | |  _ ",
		" | |_| |",
		"  \\____|",
		"        ",
	}
	letters["H"] = []string{
		"  _   _ ",
		" | | | |",
		" | |_| |",
		" |  _  |",
		" |_| |_|",
		"        ",
	}
	letters["I"] = []string{
		"  ___ ",
		" |_ _|",
		"  | | ",
		"  | | ",
		" |___|",
		"      ",
	}
	letters["J"] = []string{
		"      _ ",
		"     | |",
		"  _  | |",
		" | |_| |",
		"  \\___/ ",
		"        ",
	}
	letters["K"] = []string{
		"  _  __",
		" | |/ /",
		" | ' / ",
		" | . \\ ",
		" |_|\\_\\",
		"       ",
	}
	letters["L"] = []string{
		"  _     ",
		" | |    ",
		" | |    ",
		" | |___ ",
		" |_____|",
		"        ",
	}
	letters["M"] = []string{
		"  __  __ ",
		" |  \\/  |",
		" | |\\/| |",
		" | |  | |",
		" |_|  |_|",
		"         ",
	}
	letters["N"] = []string{
		"  _   _ ",
		" | \\ | |",
		" |  \\| |",
		" | |\\  |",
		" |_| \\_|",
		"        ",
	}
	letters["O"] = []string{
		"   ___  ",
		"  / _ \\ ",
		" | | | |",
		" | |_| |",
		"  \\___/ ",
		"        ",
	}
	letters["P"] = []string{
		"  ____  ",
		" |  _ \\ ",
		" | |_) |",
		" |  __/ ",
		" |_|    ",
		"        ",
	}
	letters["Q"] = []string{
		"   ___  ",
		"  / _ \\ ",
		" | | | |",
		" | |_| |",
		"  \\__\\_\\",
		"        ",
	}
	letters["R"] = []string{
		"  ____  ",
		" |  _ \\ ",
		" | |_) |",
		" |  _ < ",
		" |_| \\_\\",
		"        ",
	}
	letters["S"] = []string{
		"  ____  ",
		" / ___| ",
		" \\___ \\ ",
		"  ___) |",
		" |____/ ",
		"        ",
	}
	letters["T"] = []string{
		"  _____ ",
		" |_   _|",
		"   | |  ",
		"   | |  ",
		"   |_|  ",
		"        ",
	}
	letters["U"] = []string{
		"  _   _ ",
		" | | | |",
		" | | | |",
		" | |_| |",
		"  \\___/ ",
		"        ",
	}
	letters["V"] = []string{
		" __     __",
		" \\ \\   / /",
		"  \\ \\ / / ",
		"   \\ V /  ",
		"    \\_/   ",
		"          ",
	}
	letters["W"] = []string{
		" __        __",
		" \\ \\      / /",
		"  \\ \\ /\\ / / ",
		"   \\ V  V /  ",
		"    \\_/\\_/   ",
		"             ",
	}
	letters["X"] = []string{
		" __  __",
		" \\ \\/ /",
		"  \\  / ",
		"  /  \\ ",
		" /_/\\_\\",
		"       ",
	}
	letters["Y"] = []string{
		" __   __",
		" \\ \\ / /",
		"  \\ V / ",
		"   | |  ",
		"   |_|  ",
		"        ",
	}
	letters["Z"] = []string{
		"  _____",
		" |__  /",
		"   / / ",
		"  / /_ ",
		" /____|",
		"       ",
	}
	letters["!"] = []string{
		"  _ ",
		" | |",
		" | |",
		" |_|",
		" (_)",
		"    ",
	}
	letters["?"] = []string{
		"  ___ ",
		" |__ \\",
		"   / /",
		"  |_| ",
		"  (_) ",
		"      ",
	}
	letters["."] = []string{
		"    ",
		"    ",
		"    ",
		"  _ ",
		" (_)",
		"    ",
	}
	letters[","] = []string{
		"    ",
		"    ",
		"    ",
		"  _ ",
		" ( )",
		" |/ ",
	}
	letters[";"] = []string{
		"    ",
		"  _ ",
		" (_)",
		"  _ ",
		" ( )",
		" |/ ",
	}
	letters[":"] = []string{
		"    ",
		"  _ ",
		" (_)",
		"  _ ",
		" (_)",
		"    ",
	}
	letters["/"] = []string{
		"     __",
		"    / /",
		"   / / ",
		"  / /  ",
		" /_/   ",
		"       ",
	}
	letters["!"] = []string{
		"  _ ",
		" | |",
		" | |",
		" |_|",
		" (_)",
		"    ",
	}
	letters["="] = []string{
		"        ",
		"  _____ ",
		" |_____|",
		" |_____|",
		"        ",
		"        ",
	}
	letters["+"] = []string{
		"        ",
		"    _   ",
		"  _| |_ ",
		" |_   _|",
		"   |_|  ",
		"        ",
	}
	letters["&"] = []string{
		"   ___   ",
		"  ( _ )  ",
		"  / _ \\/\\",
		" | (_>  <",
		"  \\___/\\/",
		"         ",
	}
	letters["#"] = []string{
		"    _  _   ",
		"  _| || |_ ",
		" |_  ..  _|",
		" |_      _|",
		"   |_||_|  ",
		"           ",
	}
	letters["-"] = []string{
		"        ",
		"        ",
		"  _____ ",
		" |_____|",
		"        ",
		"        ",
	}
	letters["_"] = []string{
		"        ",
		"        ",
		"        ",
		"        ",
		"  _____ ",
		" |_____|",
	}
	letters["("] = []string{
		"   __",
		"  / /",
		" | | ",
		" | | ",
		" | | ",
		"  \\_\\",
	}
	letters[")"] = []string{
		" __  ",
		" \\ \\ ",
		"  | |",
		"  | |",
		"  | |",
		" /_/ ",
	}
	letters["["] = []string{
		"  __ ",
		" | _|",
		" | | ",
		" | | ",
		" | | ",
		" |__|",
	}
	letters["]"] = []string{
		"  __ ",
		" |_ |",
		"  | |",
		"  | |",
		"  | |",
		" |__|",
	}
	letters["{"] = []string{
		"    __",
		"   / /",
		"  | | ",
		" < <  ",
		"  | | ",
		"   \\_\\",
	}
	letters["}"] = []string{
		" __   ",
		" \\ \\  ",
		"  | | ",
		"   > >",
		"  | | ",
		" /_/  ",
	}
	letters["|"] = []string{
		"  _ ",
		" | |",
		" | |",
		" | |",
		" | |",
		" |_|",
	}
	letters["'"] = []string{
		"  _ ",
		" ( )",
		" |/ ",
		"    ",
		"    ",
		"    ",
	}
	letters["%"] = []string{
		"  _  __",
		" (_)/ /",
		"   / / ",
		"  / /_ ",
		" /_/(_)",
		"       ",
	}
	letters["^"] = []string{
		"  /\\ ",
		" |/\\|",
		"     ",
		"     ",
		"     ",
		"     ",
	}
	letters["$"] = []string{
		"   _  ",
		"  | | ",
		" / __)",
		" \\__ \\",
		" (   /",
		"  |_| ",
	}
	letters["*"] = []string{
		"       ",
		" __/\\__",
		" \\    /",
		" /_  _\\",
		"   \\/  ",
		"       ",
	}
	letters["1"] = []string{
		"  _ ",
		" / |",
		" | |",
		" | |",
		" |_|",
		"    ",
	}
	letters["2"] = []string{
		"  ____  ",
		" |___ \\ ",
		"   __) |",
		"  / __/ ",
		" |_____|",
		"        ",
	}
	letters["3"] = []string{
		"  _____ ",
		" |___ / ",
		"   |_ \\ ",
		"  ___) |",
		" |____/ ",
		"        ",
	}
	letters["4"] = []string{
		"  _  _   ",
		" | || |  ",
		" | || |_ ",
		" |__   _|",
		"    |_|  ",
		"         ",
	}
	letters["5"] = []string{
		"  ____  ",
		" | ___| ",
		" |___ \\ ",
		"  ___) |",
		" |____/ ",
		"        ",
	}
	letters["6"] = []string{
		"   __   ",
		"  / /_  ",
		" | '_ \\ ",
		" | (_) |",
		"  \\___/ ",
		"        ",
	}
	letters["7"] = []string{
		"  _____ ",
		" |___  |",
		"    / / ",
		"   / /  ",
		"  /_/   ",
		"        ",
	}
	letters["8"] = []string{
		"   ___  ",
		"  ( _ ) ",
		"  / _ \\ ",
		" | (_) |",
		"  \\___/ ",
		"        ",
	}
	letters["9"] = []string{
		"   ___  ",
		"  / _ \\ ",
		" | (_) |",
		"  \\__, |",
		"    /_/ ",
		"        ",
	}
	letters["0"] = []string{
		"   ___  ",
		"  / _ \\ ",
		" | | | |",
		" | |_| |",
		"  \\___/ ",
		"        ",
	}
	letters[" "] = []string{
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
	}

	if _, exists := letters[string(letter)]; exists {
		return letters[string(letter)]
	} else {
		return letters["?"]
	}
}

func AsciiText(str string, customSettings ...settings.Settings) string {

	settings := settings.GetSettings(customSettings)

	var asciiString [6]string
	var asciiLetter [6]string
	var result string

	for _, letter := range str {
		asciiLetter = [6]string(getLetter(letter))
		for i := 0; i < 6; i++ {
			asciiString[i] += asciiLetter[i]
		}
	}

	result += strings.Repeat("\n", settings.TopPadding)
	for i := 0; i < 6; i++ {
		result += strings.Repeat(" ", int(settings.LeftPadding))
		if len(asciiString[i]) > settings.MaxCols-settings.LeftPadding {
			result += asciiString[i][:settings.MaxCols-settings.LeftPadding]
		} else {
			result += asciiString[i]
		}
		if i < 5 {
			result += "\n"
		}
	}
	result += strings.Repeat("\n", settings.BottomPadding)

	return result
}
