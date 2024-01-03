package number_picker

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
	Default          float64
	Increment        float64
	Round            uint8
	Maximum          float64
	Minimum          float64
	Decimal          bool
	UnknownKeysErr   bool
}

func getDefaultSettings() Settings {
	return Settings{
		PromptForeground: RGBColor{230, 240, 255},
		BorderForeground: RGBColor{69, 71, 90},
		InputForeground:  RGBColor{180, 190, 254},
		DontCleanup:      false,
		MaxCols:          40,
		LeftPadding:      3,
		Prompt:           "Choose a number:",
		Default:          0,
		Increment:        1,
		Round:            2,
		Maximum:          100,
		Minimum:          0,
		Decimal:          false,
		UnknownKeysErr:   false,
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

	if customSettings.Decimal {
		defaultSettings.Decimal = customSettings.Decimal
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

	if customSettings.Default != 0.0 {
		defaultSettings.Default = customSettings.Default
	}

	if customSettings.Increment != 0.0 {
		defaultSettings.Increment = customSettings.Increment
	}

	if customSettings.Round != 0 {
		defaultSettings.Round = customSettings.Round
	}

	if customSettings.Maximum != 0.0 {
		defaultSettings.Maximum = customSettings.Maximum
	}

	if customSettings.Minimum != 0.0 {
		defaultSettings.Minimum = customSettings.Minimum
	}

	if customSettings.UnknownKeysErr {
		defaultSettings.UnknownKeysErr = customSettings.UnknownKeysErr
	}

	return defaultSettings
}
