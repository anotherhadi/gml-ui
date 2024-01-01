package confirm_inline

type RGBColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Settings struct {
	PromptForeground  RGBColor
	ConfirmForeground RGBColor
	DontCleanup       bool
	DefaultToFalse    bool
	LeftPadding       uint8
	Prompt            string
}

func getDefaultSettings() Settings {
	return Settings{
		PromptForeground:  RGBColor{230, 240, 255},
		ConfirmForeground: RGBColor{69, 71, 90},
		DefaultToFalse:    false,
		DontCleanup:       false,
		LeftPadding:       3,
		Prompt:            "Are you sure?",
	}
}

func combineSettings(customSettings Settings) Settings {
	defaultSettings := getDefaultSettings()

	if customSettings.PromptForeground != (RGBColor{}) {
		defaultSettings.PromptForeground = customSettings.PromptForeground
	}

	if customSettings.ConfirmForeground != (RGBColor{}) {
		defaultSettings.ConfirmForeground = customSettings.ConfirmForeground
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

	if customSettings.Prompt != "" {
		defaultSettings.Prompt = customSettings.Prompt
	}

	return defaultSettings
}
