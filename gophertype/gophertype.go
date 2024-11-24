package gophertype

import "github.com/charmbracelet/bubbletea/v2"

type gopherType struct {}

func New() gopherType {
    return gopherType{}
}

func (g gopherType) Init() (tea.Model, tea.Cmd) {
    return g, nil
}

func (g gopherType) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return g, nil
}

func (g gopherType) View() string {
    return "Hello, Gophers!"
}
