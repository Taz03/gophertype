package test

import (
	"github.com/charmbracelet/bubbletea/v2"
	"github.com/taz03/gophertype/gophertype/test/test_state"
)

func (t Test) Update(msg tea.Msg) (m tea.Model, cmd tea.Cmd) {
	switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        t.w = msg.Width
        t.h = msg.Height

	case tea.KeyMsg:
        if t.TestState == teststate.INITIALIZED {
            t.TestState = teststate.STARTED
            cmd = tea.Batch(cmd, t.cursorCmd(true))
        }

		switch msg {
		}
	}

	return t, cmd
}
