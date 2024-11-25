package theme

import (
	"embed"
	"encoding/json"
	"fmt"
	"image/color"

	"github.com/taz03/gophertype/config"
)

type Theme struct {
	Name string

	BgColor                 color.Color
	MainColor               color.Color
	CaretColor              color.Color
	SubColor                color.Color
	SubAltColor             color.Color
	TextColor               color.Color
	ErrorColor              color.Color
	ErrorExtraColor         color.Color
	ColorfulErrorColor      color.Color
	ColorfulErrorExtraColor color.Color
}

type jsonTheme struct {
	BgColor                 string `json:"bg_color"`
	MainColor               string `json:"main_color"`
	CaretColor              string `json:"caret_color"`
	SubColor                string `json:"sub_color"`
	SubAltColor             string `json:"sub_alt_color"`
	TextColor               string `json:"text_color"`
	ErrorColor              string `json:"error_color"`
	ErrorExtraColor         string `json:"error_extra_color"`
	ColorfulErrorColor      string `json:"colorful_error_color"`
	ColorfulErrorExtraColor string `json:"colorful_error_extra_color"`
}

func New(resourcesFs embed.FS, cfg config.Config) Theme {
	var jsonTheme jsonTheme
	bytes, _ := resourcesFs.ReadFile("themes/" + cfg.Theme + ".json")
	json.Unmarshal(bytes, &jsonTheme)

	bgColor, _ := parseHexColor(jsonTheme.BgColor)
    if (cfg.TransparentBackground) {
        bgColor = color.Transparent
    }

	mainColor, _ := parseHexColor(jsonTheme.MainColor)
	caretColor, _ := parseHexColor(jsonTheme.CaretColor)
	subColor, _ := parseHexColor(jsonTheme.SubColor)
	subAltColor, _ := parseHexColor(jsonTheme.SubAltColor)
	textColor, _ := parseHexColor(jsonTheme.TextColor)
	errorColor, _ := parseHexColor(jsonTheme.ErrorColor)
	errorExtraColor, _ := parseHexColor(jsonTheme.ErrorExtraColor)
	colorfulErrorColor, _ := parseHexColor(jsonTheme.ColorfulErrorColor)
	colorfulErrorExtraColor, _ := parseHexColor(jsonTheme.ColorfulErrorExtraColor)

	return Theme{
		Name: cfg.Theme,

		BgColor:                 bgColor,
		MainColor:               mainColor,
		CaretColor:              caretColor,
		SubColor:                subColor,
		SubAltColor:             subAltColor,
		TextColor:               textColor,
		ErrorColor:              errorColor,
		ErrorExtraColor:         errorExtraColor,
		ColorfulErrorColor:      colorfulErrorColor,
		ColorfulErrorExtraColor: colorfulErrorExtraColor,
	}
}

func parseHexColor(s string) (color.Color, error) {
    c := color.RGBA{}
    var err error

	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")
	}

    return c, err
}
