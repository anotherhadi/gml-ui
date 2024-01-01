package input

type RGBColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Settings struct {
	PromptForeground RGBColor
	InputForeground  RGBColor
	DontCleanup      bool
	LeftPadding      uint8
	Prompt           string
	Default          string
}

func getDefaultSettings() Settings {
	return Settings{
		PromptForeground: RGBColor{230, 240, 255},
		InputForeground:  RGBColor{180, 190, 254},
		DontCleanup:      false,
		LeftPadding:      3,
		Prompt:           "Input:",
		Default:          "default",
	}
}

func combineSettings(customSettings Settings) Settings {
	defaultSettings := getDefaultSettings()

	if customSettings.PromptForeground != (RGBColor{}) {
		defaultSettings.PromptForeground = customSettings.PromptForeground
	}

	if customSettings.InputForeground != (RGBColor{}) {
		defaultSettings.InputForeground = customSettings.InputForeground
	}

	if customSettings.DontCleanup {
		defaultSettings.DontCleanup = customSettings.DontCleanup
	}

	if customSettings.Default != "" {
		defaultSettings.Default = customSettings.Default
	}

	if customSettings.LeftPadding != 0 {
		defaultSettings.LeftPadding = customSettings.LeftPadding
	}

	if customSettings.Prompt != "" {
		defaultSettings.Prompt = customSettings.Prompt
	}

	return defaultSettings
}
