package paragraph

type RGBColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Settings struct {
	Foreground    RGBColor
	Background    RGBColor
	LeftPadding   uint8
	TopPadding    uint8
	BottomPadding uint8
	MaxCols       uint8
}

func getDefaultSettings() Settings {
	return Settings{
		Foreground:    RGBColor{230, 240, 255},
		Background:    RGBColor{},
		LeftPadding:   3,
		TopPadding:    0,
		BottomPadding: 0,
		MaxCols:       80,
	}
}

func combineSettings(customSettings Settings) Settings {
	var defaultSettings Settings = getDefaultSettings()

	if customSettings.Foreground != (RGBColor{}) {
		defaultSettings.Foreground = customSettings.Foreground
	}

	if customSettings.Background != (RGBColor{}) {
		defaultSettings.Background = customSettings.Background
	}

	if customSettings.LeftPadding != 0 {
		defaultSettings.LeftPadding = customSettings.LeftPadding
	}

	if customSettings.TopPadding != 0 {
		defaultSettings.TopPadding = customSettings.TopPadding
	}

	if customSettings.BottomPadding != 0 {
		defaultSettings.BottomPadding = customSettings.BottomPadding
	}

	if customSettings.MaxCols != 0 {
		defaultSettings.MaxCols = customSettings.MaxCols
	}

	return defaultSettings
}
