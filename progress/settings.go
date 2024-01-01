package progress

type RGBColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Settings struct {
	ProgressForeground       RGBColor
	ProgressSecondForeground RGBColor
	EmptyForeground          RGBColor
	PercentageForeground     RGBColor
	DontCleanup              bool
	LeftPadding              uint8
	Width                    uint8
}

func getDefaultSettings() Settings {
	return Settings{
		ProgressForeground:       RGBColor{180, 190, 254},
		ProgressSecondForeground: RGBColor{220, 230, 255},
		EmptyForeground:          RGBColor{69, 71, 90},
		PercentageForeground:     RGBColor{230, 240, 255},
		DontCleanup:              false,
		LeftPadding:              3,
		Width:                    40,
	}
}

func combineSettings(customSettings Settings) Settings {
	defaultSettings := getDefaultSettings()

	if customSettings.ProgressForeground != (RGBColor{}) {
		defaultSettings.ProgressForeground = customSettings.ProgressForeground
	}

	if customSettings.ProgressSecondForeground != (RGBColor{}) {
		defaultSettings.ProgressSecondForeground = customSettings.ProgressSecondForeground
	} else if customSettings.ProgressForeground != (RGBColor{}) {
		defaultSettings.ProgressSecondForeground = customSettings.ProgressForeground
	}

	if customSettings.EmptyForeground != (RGBColor{}) {
		defaultSettings.EmptyForeground = customSettings.EmptyForeground
	}

	if customSettings.PercentageForeground != (RGBColor{}) {
		defaultSettings.PercentageForeground = customSettings.PercentageForeground
	}

	if customSettings.DontCleanup {
		defaultSettings.DontCleanup = customSettings.DontCleanup
	}

	if customSettings.LeftPadding != 0 {
		defaultSettings.LeftPadding = customSettings.LeftPadding
	}

	if customSettings.Width != 0 {
		defaultSettings.Width = customSettings.Width
	}

	return defaultSettings
}
