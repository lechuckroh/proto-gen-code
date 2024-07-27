package protoinput

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	input    textinput.Model
	selected bool
	keyMap   KeyMap
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.selected {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.SelectFile):
			m.selected = true
			return m, SelectFile(m.input.Value())
		}
	}

	newInput, cmd := m.input.Update(msg)
	m.input = newInput
	return m, cmd
}

func (m Model) View() string {
	if m.selected {
		return fmt.Sprintf("Source proto file: %s", m.input.Value())
	}
	return m.input.View()
}

func NewModel() Model {
	ti := textinput.New()
	ti.Prompt = "Enter proto file to load: "
	ti.Focus()

	return Model{
		input:    ti,
		selected: false,
		keyMap:   defaultKeyMap,
	}
}
