package main

import (
	"embed"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/taz03/gophertype/config"
	"github.com/taz03/gophertype/gophertype"
)

//go:embed resources/*
var resourcesFs embed.FS

func main() {
    cfg := config.New("~/.config/gophertype/config.json")
    g := gophertype.New(cfg, resourcesFs)

    p := tea.NewProgram(g, tea.WithColorProfile(colorprofile.TrueColor), tea.WithAltScreen())
    p.Run()
}
