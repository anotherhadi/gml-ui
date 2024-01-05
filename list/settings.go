package list

import "errors"

type RGBColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Options struct {
	Title       string
	Description string
}

type Settings struct {
	SelectedTitleForeground       RGBColor
	SelectedDescriptionForeground RGBColor
	UnselectedForeground          RGBColor
	DontCleanup                   bool
	LeftPadding                   uint8
	Options                       []Options
	UnknownKeysErr                bool
	MaxRows                       int
}

func getDefaultSettings() Settings {
	return Settings{
		SelectedTitleForeground:       RGBColor{180, 190, 254},
		SelectedDescriptionForeground: RGBColor{99, 101, 120},
		UnselectedForeground:          RGBColor{69, 71, 90},
		DontCleanup:                   false,
		LeftPadding:                   3,
		Options:                       []Options{},
		UnknownKeysErr:                false,
		MaxRows:                       30,
	}
}

func combineSettings(customSettings Settings) (Settings, error) {
	defaultSettings := getDefaultSettings()

	if len(customSettings.Options) > 0 {
		defaultSettings.Options = customSettings.Options
	} else {
		return Settings{}, errors.New("No Options provided")
	}

	if customSettings.SelectedTitleForeground != (RGBColor{}) {
		defaultSettings.SelectedTitleForeground = customSettings.SelectedTitleForeground
	}

	if customSettings.SelectedDescriptionForeground != (RGBColor{}) {
		defaultSettings.SelectedDescriptionForeground = customSettings.SelectedDescriptionForeground
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

	if customSettings.UnknownKeysErr {
		defaultSettings.UnknownKeysErr = customSettings.UnknownKeysErr
	}

	if customSettings.MaxRows != 0 {
		defaultSettings.MaxRows = customSettings.MaxRows
	}

	return defaultSettings, nil
}
