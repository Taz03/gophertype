package gophertype

import (
	"embed"

	"github.com/charmbracelet/bubbletea/v2"
	"github.com/taz03/gophertype/config"
	"github.com/taz03/gophertype/gophertype/test"
	"github.com/taz03/gophertype/gophertype/theme"
)

type GopherType struct {
    Config config.Config
    Theme  theme.Theme
    
    Test test.Test
}

func New(cfg config.Config, resourcesFs embed.FS) GopherType {
    theme := theme.New(resourcesFs, cfg)
    test := test.New(resourcesFs, &cfg, &theme)

    return GopherType{
        Config: cfg,
        Theme:  theme,

        Test: test,
    }
}

func (g GopherType) Init() (tea.Model, tea.Cmd) {
    _, cmd := g.Test.Init()

    return g, cmd
}

func (g GopherType) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return g, nil
}

func (g GopherType) View() string {
    return "Hello, Gophers!"
}
