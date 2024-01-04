package rgbapicker

type RGBColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Settings struct {
	PromptForeground RGBColor
	BorderForeground RGBColor
	InputForeground  RGBColor
	DontCleanup      bool
	MaxCols          uint8
	LeftPadding      uint8
	Prompt           string
	Default          [4]int
	UnknownKeysErr   bool
	BoxStyle         string
}

func getDefaultSettings() Settings {
	return Settings{
		PromptForeground: RGBColor{230, 240, 255},
		BorderForeground: RGBColor{69, 71, 90},
		InputForeground:  RGBColor{180, 190, 254},
		DontCleanup:      false,
		MaxCols:          40,
		LeftPadding:      3,
		Prompt:           "Choose a color (RGBA):",
		Default:          [4]int{180, 190, 254, 255},
		UnknownKeysErr:   false,
		BoxStyle:         "-",
	}
}

func combineSettings(customSettings Settings) Settings {
	defaultSettings := getDefaultSettings()

	if customSettings.PromptForeground != (RGBColor{}) {
		defaultSettings.PromptForeground = customSettings.PromptForeground
	}

	if customSettings.BorderForeground != (RGBColor{}) {
		defaultSettings.BorderForeground = customSettings.BorderForeground
	}

	if customSettings.InputForeground != (RGBColor{}) {
		defaultSettings.InputForeground = customSettings.InputForeground
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

	if customSettings.Default != ([4]int{}) {
		defaultSettings.Default = customSettings.Default
	}

	if customSettings.UnknownKeysErr {
		defaultSettings.UnknownKeysErr = customSettings.UnknownKeysErr
	}

	if customSettings.BoxStyle != "" {
		defaultSettings.BoxStyle = customSettings.BoxStyle
	}

	return defaultSettings
}
