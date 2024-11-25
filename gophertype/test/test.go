package test

import (
	"embed"

	"github.com/charmbracelet/bubbletea/v2"
	"github.com/taz03/gophertype/config"
	"github.com/taz03/gophertype/gophertype/test/test_state"
	"github.com/taz03/gophertype/gophertype/theme"
	"golang.org/x/term"
)

type Test struct {
	cfg   *config.Config
	theme *theme.Theme

    width int

    testWidth, testHeight int

	words      *TestWords
	typedWords []string
	pos        [2]int

	TestState teststate.TestState
}

func New(resourcesFs embed.FS, cfg *config.Config, theme *theme.Theme) Test {
    w, _, _ := term.GetSize(0)

	return Test{
		cfg:   cfg,
		theme: theme,

        width: w,

		words: NewTestWords(*cfg, resourcesFs),
		pos:   [2]int{0, 0},

		TestState: teststate.INITIALIZED,
	}
}

func (t Test) Init() (tea.Model, tea.Cmd) {
	return t, t.cursorCmd(true)
}
