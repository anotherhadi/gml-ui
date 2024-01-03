package selection

import "errors"

type RGBColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Settings struct {
	SelectedForeground   RGBColor
	UnselectedForeground RGBColor
	PromptForeground     RGBColor
	DontCleanup          bool
	LeftPadding          uint8
	MaxCols              uint8
	Prompt               string
	Options              []string
	UnknownKeysErr       bool
}

func getDefaultSettings() Settings {
	return Settings{
		PromptForeground:     RGBColor{230, 240, 255},
		SelectedForeground:   RGBColor{180, 190, 254},
		UnselectedForeground: RGBColor{69, 71, 90},
		DontCleanup:          false,
		LeftPadding:          3,
		MaxCols:              40,
		Prompt:               "Select an option:",
		Options:              []string{},
		UnknownKeysErr:       false,
	}
}

func combineSettings(customSettings Settings) (Settings, error) {
	defaultSettings := getDefaultSettings()

	if len(customSettings.Options) > 0 {
		defaultSettings.Options = customSettings.Options
	} else {
		return Settings{}, errors.New("No Options provided")
	}

	if customSettings.PromptForeground != (RGBColor{}) {
		defaultSettings.PromptForeground = customSettings.PromptForeground
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

	if customSettings.Prompt != "" {
		defaultSettings.Prompt = customSettings.Prompt
	}

	if customSettings.UnknownKeysErr {
		defaultSettings.UnknownKeysErr = customSettings.UnknownKeysErr
	}

	return defaultSettings, nil
}
