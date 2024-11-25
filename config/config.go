package config

import (
	"encoding/json"
	"os"

	"github.com/taz03/gophertype/config/caret"
	"github.com/taz03/gophertype/config/confidence_mode"
	"github.com/taz03/gophertype/config/difficulty"
	"github.com/taz03/gophertype/config/quick_restart"
	"github.com/taz03/gophertype/config/stop_on_error"
	"github.com/taz03/gophertype/config/timer"
	"github.com/taz03/gophertype/config/timer_color"
)

type Config struct {
	Theme                   string                        `json:"theme"`
	TransparentBackground   bool                          `json:"transparentBackground"`
	CustomTheme             bool                          `json:"customTheme"`
	CustomThemeColors       []string                      `json:"customThemeColors"`
	FavThemes               []string                      `json:"favThemes"`
	ShowKeyTips             bool                          `json:"showKeyTips"`
	QuickRestart            quick_restart.QuickRestart    `json:"quickRestart"`
	Punctuation             bool                          `json:"punctuation"`
	Numbers                 bool                          `json:"numbers"`
	Mode                    string                        `json:"mode"`
	QuoteLength             []int                         `json:"quoteLength"`
	Language                string                        `json:"language"`
	FreedomMode             bool                          `json:"freedomMode"`
	Difficulty              difficulty.Difficulty         `json:"difficulty"`
	BlindMode               bool                          `json:"blindMode"`
	QuickEnd                bool                          `json:"quickEnd"`
	CaretStyle              caret.Caret                   `json:"caretStyle"`
	FlipTestColors          bool                          `json:"flipTestColors"`
	ConfidenceMode          confidencemode.ConfidenceMode `json:"confidenceMode"`
	IndicateTypos           bool                          `json:"indicateTypos"`
	TimerStyle              timer.Timer                   `json:"timerStyle"`
	LiveSpeed               bool                          `json:"liveSpeedStyle"`
	LiveAcc                 bool                          `json:"liveAccStyle"`
	LiveBurst               bool                          `json:"liveBurstStyle"`
	ColorfulMode            bool                          `json:"colorfulMode"`
	RandomTheme             string                        `json:"randomTheme"`
	TimerColor              timercolor.TimerColor         `json:"timerColor"`
	StopOnError             stoponerror.StopOnError       `json:"stopOnError"`
	ShowAllLines            bool                          `json:"showAllLines"`
	AlwaysShowDecimalPlaces bool                          `json:"alwaysShowDecimalPlaces"`
	AlwaysShowWordsHistory  bool                          `json:"alwaysShowWordsHistory"`
	SingleListCommandLine   bool                          `json:"singleListCommandLine"`
	CapsLockWarning         bool                          `json:"capsLockWarning"`
	PlaySoundOnError        string                        `json:"playSoundOnError"`
	PlaySoundOnClick        string                        `json:"playSoundOnClick"`
	SoundVolume             float32                       `json:"soundVolume"`
	StartGraphsAtZero       bool                          `json:"startGraphsAtZero"`
	PaceCaret               string                        `json:"paceCaret"`
	PaceCaretCustomSpeed    int                           `json:"paceCaretCustomSpeed"`
	RepeatedPace            bool                          `json:"repeatedPace"`
	AccountChart            []string                      `json:"accountChart"`
	MinWpm                  string                        `json:"minWpm"`
	MinWpmCustomSpeed       int                           `json:"minWpmCustomSpeed"`
	HighlightMode           string                        `json:"highlightMode"`
	TypingSpeedUnit         string                        `json:"typingSpeedUnit"`
	HideExtraLetters        bool                          `json:"hideExtraLetters"`
	StrictSpace             bool                          `json:"strictSpace"`
	MinAcc                  string                        `json:"minAcc"`
	MinAccCustom            int                           `json:"minAccCustom"`
	RepeatQuotes            string                        `json:"repeatQuotes"`
	OppositeShiftMode       string                        `json:"oppositeShiftMode"`
	CustomBackground        string                        `json:"customBackground"`
	CustomLayoutFluid       string                        `json:"customLayoutFluid"`
	MinBurst                string                        `json:"minBurst"`
	MinBurstCustomSpeed     int                           `json:"minBurstCustomSpeed"`
	BurstHeatmap            bool                          `json:"burstHeatmap"`
	BritishEnglish          bool                          `json:"britishEnglish"`
	LazyMode                bool                          `json:"lazyMode"`
	ShowAverage             string                        `json:"showAverage"`
	TapeMode                string                        `json:"tapeMode"`
	MaxLineWidth            int                           `json:"maxLineWidth"`
}

func New(path string) Config {
	fileContent, _ := os.ReadFile(path)

	// Fill with default values
	config := Config{
		Theme:                   "sekira_dark",
        TransparentBackground:   false,
		CustomTheme:             false,
		CustomThemeColors:       []string{"#323437", "#e2b714", "#e2b714", "#646669", "#2c2e31", "#d1d0c5", "#ca4754", "#7e2a33", "#ca4754", "#7e2a33"},
		ShowKeyTips:             false,
		QuickRestart:            quick_restart.Tab,
		Punctuation:             false,
		Numbers:                 false,
		Mode:                    "time 30",
		QuoteLength:             []int{1},
		Language:                "english",
		FreedomMode:             false,
		Difficulty:              difficulty.Normal,
		BlindMode:               false,
		QuickEnd:                false,
		CaretStyle:              caret.Line,
		FlipTestColors:          false,
		ConfidenceMode:          confidencemode.Off,
		IndicateTypos:           false,
		TimerStyle:              timer.Text,
		LiveSpeed:               true,
		LiveAcc:                 true,
		LiveBurst:               false,
		ColorfulMode:            false,
		RandomTheme:             "off",
		TimerColor:              timercolor.Sub,
		StopOnError:             stoponerror.Off,
		ShowAllLines:            false,
		AlwaysShowDecimalPlaces: false,
		AlwaysShowWordsHistory:  false,
		SingleListCommandLine:   false,
		CapsLockWarning:         true,
	}
	json.Unmarshal(fileContent, &config)

	return config
}
