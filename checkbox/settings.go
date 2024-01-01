package checkbox

import (
	"errors"
)

type RGBColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Settings struct {
	PromptForeground     RGBColor
	CheckedForeground    RGBColor
	SelectedForeground   RGBColor
	UnselectedForeground RGBColor
	DontCleanup          bool
	LeftPadding          uint8
	MaxCols              uint8
	Minimum              uint8
	Maximum              uint8
	Prompt               string
	Options              []string
	DefaultOptions       []bool
}

func getDefaultSettings() Settings {
	return Settings{
		PromptForeground:     RGBColor{230, 240, 255},
		CheckedForeground:    RGBColor{180, 190, 254},
		SelectedForeground:   RGBColor{180, 190, 254},
		UnselectedForeground: RGBColor{69, 71, 90},
		DontCleanup:          false,
		LeftPadding:          3,
		MaxCols:              40,
		Minimum:              0,
		Maximum:              0,
		Prompt:               "Select options:",
		Options:              []string{},
		DefaultOptions:       []bool{},
	}
}

func combineSettings(customSettings Settings) (Settings, error) {
	var defaultSettings Settings = getDefaultSettings()

	if len(customSettings.Options) > 0 {
		defaultSettings.Options = customSettings.Options
	} else {
		return Settings{}, errors.New("No Options provided")
	}

	if len(customSettings.DefaultOptions) > 0 {
		if len(customSettings.Options) == len(customSettings.DefaultOptions) {
			defaultSettings.DefaultOptions = customSettings.DefaultOptions
		} else {
			return Settings{}, errors.New("No DefaultOptions for every Options")
		}
	} else {
		defaultSettings.DefaultOptions = make([]bool, len(defaultSettings.Options))
	}

	if customSettings.PromptForeground != (RGBColor{}) {
		defaultSettings.PromptForeground = customSettings.PromptForeground
	}

	if customSettings.CheckedForeground != (RGBColor{}) {
		defaultSettings.CheckedForeground = customSettings.CheckedForeground
	}

	if customSettings.SelectedForeground != (RGBColor{}) {
		defaultSettings.SelectedForeground = customSettings.SelectedForeground
	}

	if customSettings.UnselectedForeground != (RGBColor{}) {
		defaultSettings.UnselectedForeground = customSettings.UnselectedForeground
	}

	if customSettings.DontCleanup {
		defaultSettings.DontCleanup = customSettings.DontCleanup
	}

	if customSettings.LeftPadding != 0 {
		defaultSettings.LeftPadding = customSettings.LeftPadding
	}

	if customSettings.MaxCols != 0 {
		defaultSettings.MaxCols = customSettings.MaxCols
	}

	if customSettings.Minimum != 0 {
		defaultSettings.Minimum = customSettings.Minimum
	}

	if customSettings.Maximum != 0 {
		defaultSettings.Maximum = customSettings.Maximum
	}

	if customSettings.Prompt != "" {
		defaultSettings.Prompt = customSettings.Prompt
	}

	return defaultSettings, nil
}
