package settings

import (
	"reflect"
)

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Settings struct {
	DefaultInt     int
	DefaultFloat   float64
	DefaultString  string
	DefaultColor   Color
	DefaultChecked []bool
	DefaultBool    bool

	Affirmative   string
	Negative      string
	Increment     float64
	Decimal       int
	Maximum       float64
	Minimum       float64
	Filter        bool
	CaseSensitive bool

	TopPadding    int
	BottomPadding int
	LeftPadding   int
	MaxCols       int
	MaxRows       int
	Style         string

	DontLoop         bool
	ExitOnUnknownKey bool

	AccentColor              Color
	AccentBackgroundColor    Color
	TextColor                Color
	TextBackgroundColor      Color
	SecondaryColor           Color
	SecondaryBackgroundColor Color
}

func getDefaultSettings() Settings {
	return Settings{
		DefaultInt:     0,
		DefaultFloat:   0.0,
		DefaultString:  "default",
		DefaultColor:   Color{180, 190, 254},
		DefaultChecked: nil,
		DefaultBool:    false,

		Affirmative:   "Yes",
		Negative:      "No",
		Increment:     1,
		Decimal:       0,
		Maximum:       1000,
		Minimum:       0,
		Filter:        false,
		CaseSensitive: false,

		TopPadding:    0,
		BottomPadding: 0,
		LeftPadding:   0,
		MaxCols:       80,
		MaxRows:       40,
		Style:         "-",

		DontLoop:         false,
		ExitOnUnknownKey: false,

		AccentColor:              Color{180, 190, 254},
		AccentBackgroundColor:    Color{17, 17, 17},
		TextColor:                Color{230, 240, 255},
		TextBackgroundColor:      Color{17, 17, 17},
		SecondaryColor:           Color{69, 71, 90},
		SecondaryBackgroundColor: Color{230, 240, 255},
	}
}

func GetSettings(customSettings []Settings) Settings {
	defaultSettings := getDefaultSettings()

	for _, custom := range customSettings {
		customValue := reflect.ValueOf(custom)
		defaultValue := reflect.ValueOf(&defaultSettings).Elem()
		for i := 0; i < customValue.NumField(); i++ {
			field := customValue.Type().Field(i)
			customFieldValue := customValue.Field(i)
			if !reflect.Value.IsZero(customFieldValue) {
				defaultValue.FieldByName(field.Name).Set(customFieldValue)
			}
		}
	}

	return defaultSettings
}
