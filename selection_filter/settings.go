package selection_filter

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
	MaxRows              uint8
	CaseSensitive        bool
	Prompt               string
	Options              []string
}

func getDefaultSettings() Settings {
	return Settings{
		PromptForeground:     RGBColor{230, 240, 255},
		SelectedForeground:   RGBColor{180, 190, 254},
		UnselectedForeground: RGBColor{69, 71, 90},
		DontCleanup:          false,
		CaseSensitive:        false,
		LeftPadding:          3,
		MaxCols:              40,
		MaxRows:              10,
		Prompt:               "Select an option:",
		Options:              []string{},
	}
}

func combineSettings(customSettings Settings) Settings {
	defaultSettings := getDefaultSettings()

	if len(customSettings.Options) > 0 {
		defaultSettings.Options = customSettings.Options
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

	if customSettings.MaxRows != 0 {
		defaultSettings.MaxRows = customSettings.MaxRows
	}

	if customSettings.CaseSensitive {
		defaultSettings.CaseSensitive = customSettings.CaseSensitive
	}

	if customSettings.Prompt != "" {
		defaultSettings.Prompt = customSettings.Prompt
	}

	return defaultSettings
}
