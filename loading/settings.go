package loading

type RGBColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Settings struct {
	LoadingForeground RGBColor
	MessageForeground RGBColor
	Message           string
	DontCleanup       bool
	FPS               uint8
	LeftPadding       uint8
}

func getDefaultSettings() Settings {
	return Settings{
		LoadingForeground: RGBColor{180, 190, 254},
		MessageForeground: RGBColor{230, 240, 255},
		Message:           "Loading",
		DontCleanup:       false,
		FPS:               5,
		LeftPadding:       3,
	}
}

func combineSettings(customSettings Settings) Settings {
	defaultSettings := getDefaultSettings()

	if customSettings.LoadingForeground != (RGBColor{}) {
		defaultSettings.LoadingForeground = customSettings.LoadingForeground
	}

	if customSettings.MessageForeground != (RGBColor{}) {
		defaultSettings.MessageForeground = customSettings.MessageForeground
	}

	if customSettings.Message != "" {
		defaultSettings.Message = customSettings.Message
	}

	if customSettings.DontCleanup {
		defaultSettings.DontCleanup = customSettings.DontCleanup
	}

	if customSettings.FPS != 0 {
		defaultSettings.FPS = customSettings.FPS
	}

	if customSettings.LeftPadding != 0 {
		defaultSettings.LeftPadding = customSettings.LeftPadding
	}

	return defaultSettings
}
