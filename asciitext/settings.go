package asciitext

type Settings struct {
	LeftPadding uint8
}

func getDefaultSettings() Settings {
	return Settings{
		LeftPadding: 3,
	}
}

func combineSettings(customSettings Settings) Settings {
	var defaultSettings Settings = getDefaultSettings()

	if customSettings.LeftPadding != 0 {
		defaultSettings.LeftPadding = customSettings.LeftPadding
	}

	return defaultSettings
}
