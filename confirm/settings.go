package confirm

type RGBColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type OptionColors struct {
	BackgroundColor RGBColor
	ForegroundColor RGBColor
}

type Settings struct {
	SelectedColor    OptionColors
	UnselectedColor  OptionColors
	PromptForeground RGBColor
	DontCleanup      bool
	DefaultToFalse   bool
	LeftPadding      uint8
	MaxCols          uint8
	Prompt           string
	Affirmative      string
	Negative         string
	UnknownKeysErr   bool
}

func getDefaultSettings() Settings {
	return Settings{
		SelectedColor: OptionColors{
			BackgroundColor: RGBColor{180, 190, 254},
			ForegroundColor: RGBColor{17, 17, 27},
		},
		UnselectedColor: OptionColors{
			BackgroundColor: RGBColor{69, 71, 90},
			ForegroundColor: RGBColor{230, 240, 255},
		},
		PromptForeground: RGBColor{230, 240, 255},
		DefaultToFalse:   false,
		DontCleanup:      false,
		MaxCols:          40,
		LeftPadding:      3,
		Prompt:           "Are you sure?",
		Affirmative:      "Yes",
		Negative:         "No",
		UnknownKeysErr:   false,
	}
}

func combineSettings(customSettings Settings) Settings {
	defaultSettings := getDefaultSettings()

	if customSettings.SelectedColor.BackgroundColor != (RGBColor{}) {
		defaultSettings.SelectedColor.BackgroundColor = customSettings.SelectedColor.BackgroundColor
	}

	if customSettings.SelectedColor.ForegroundColor != (RGBColor{}) {
		defaultSettings.SelectedColor.ForegroundColor = customSettings.SelectedColor.ForegroundColor
	}

	if customSettings.UnselectedColor.BackgroundColor != (RGBColor{}) {
		defaultSettings.UnselectedColor.BackgroundColor = customSettings.UnselectedColor.BackgroundColor
	}

	if customSettings.UnselectedColor.ForegroundColor != (RGBColor{}) {
		defaultSettings.UnselectedColor.ForegroundColor = customSettings.UnselectedColor.ForegroundColor
	}

	if customSettings.PromptForeground != (RGBColor{}) {
		defaultSettings.PromptForeground = customSettings.PromptForeground
	}

	if customSettings.DontCleanup {
		defaultSettings.DontCleanup = customSettings.DontCleanup
	}

	if customSettings.DefaultToFalse {
		defaultSettings.DefaultToFalse = customSettings.DefaultToFalse
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

	if customSettings.Affirmative != "" {
		defaultSettings.Affirmative = customSettings.Affirmative
	}

	if customSettings.Negative != "" {
		defaultSettings.Negative = customSettings.Negative
	}

	if customSettings.UnknownKeysErr {
		defaultSettings.UnknownKeysErr = customSettings.UnknownKeysErr
	}

	return defaultSettings
}
