package protoinput

import tea "github.com/charmbracelet/bubbletea"

type SelectFileMsg struct {
	Filename string
}

func SelectFile(filename string) tea.Cmd {
	return func() tea.Msg {
		return SelectFileMsg{
			Filename: filename,
		}
	}
}
