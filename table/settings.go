package table

import "errors"

type RGBColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Settings struct {
	LabelForeground  RGBColor
	TextForeground   RGBColor
	BorderForeground RGBColor
	Separator        bool
	Style            string
	Align            string
	LabelDecoration  string
	MaxLengthsCol    uint8
	LeftPadding      uint8
	Elements         [][]string
}

func getDefaultSettings() Settings {
	return Settings{
		LabelForeground:  RGBColor{180, 190, 254},
		TextForeground:   RGBColor{230, 240, 255},
		BorderForeground: RGBColor{69, 71, 90},
		Separator:        false,
		Style:            "-",
		Align:            "left",
		LabelDecoration:  "default",
		MaxLengthsCol:    40,
		LeftPadding:      3,
		Elements:         nil,
	}
}

func combineSettings(customSettings Settings) (Settings, error) {
	defaultSettings := getDefaultSettings()

	if customSettings.LabelForeground != (RGBColor{}) {
		defaultSettings.LabelForeground = customSettings.LabelForeground
	}

	if customSettings.TextForeground != (RGBColor{}) {
		defaultSettings.TextForeground = customSettings.TextForeground
	}

	if customSettings.BorderForeground != (RGBColor{}) {
		defaultSettings.BorderForeground = customSettings.BorderForeground
	}

	if customSettings.Separator {
		defaultSettings.Separator = customSettings.Separator
	}

	if customSettings.Style != "" {
		defaultSettings.Style = customSettings.Style
	}

	if customSettings.Align != "" {
		defaultSettings.Align = customSettings.Align
	}

	if customSettings.LabelDecoration != "" {
		defaultSettings.LabelDecoration = customSettings.LabelDecoration
	}

	if customSettings.LeftPadding != 0 {
		defaultSettings.LeftPadding = customSettings.LeftPadding
	}

	if customSettings.Elements != nil {
		defaultSettings.Elements = customSettings.Elements
	} else {
		return Settings{}, errors.New("No Elements")
	}

	return defaultSettings, nil
}
