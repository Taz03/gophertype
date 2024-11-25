package test

import (
	"github.com/charmbracelet/bubbletea/v2"
	"github.com/taz03/gophertype/config/caret"
)

func (t *Test) cursorCmd(blink bool) (cmd tea.Cmd) {
	switch t.cfg.CaretStyle {
	case caret.Off:
		cmd = tea.Batch(tea.HideCursor)
	case caret.Line:
		cmd = tea.Batch(tea.ShowCursor, tea.SetCursorStyle(tea.CursorBar, blink), tea.SetCursorColor(t.theme.CaretColor))
	case caret.Block:
		cmd = tea.Batch(tea.ShowCursor, tea.SetCursorStyle(tea.CursorBlock, blink), tea.SetCursorColor(t.theme.CaretColor))
	case caret.Underline:
		cmd = tea.Batch(tea.ShowCursor, tea.SetCursorStyle(tea.CursorUnderline, blink), tea.SetCursorColor(t.theme.CaretColor))
	}

    return
}
