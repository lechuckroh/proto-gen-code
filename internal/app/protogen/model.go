package protogen

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lechuckroh/protogencode/internal/app/protogen/option"
	"github.com/lechuckroh/protogencode/internal/app/protogen/protoinput"
)

type State int

const (
	InputSource State = iota + 1
	ProtoLoading
	SelectLanguage
	SelectTypes
	InputTarget
	Generating
	Finished
)

var docStyle = lipgloss.NewStyle().Margin(0, 0)

type Model struct {
	protoInput  protoinput.Model
	option      option.Model
	currentStep State
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		_, _ = docStyle.GetFrameSize()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	case protoinput.SelectFileMsg:
		m.currentStep = ProtoLoading
		break
	}

	switch m.currentStep {
	case InputSource:
		newModel, cmd := m.protoInput.Update(msg)
		m.protoInput = newModel.(protoinput.Model)
		return m, cmd
	case ProtoLoading:
		return m, nil
	case SelectLanguage:
		newModel, cmd := m.option.Update(msg)
		m.option = newModel.(option.Model)
		return m, cmd
	case SelectTypes:
		return m, nil
	case InputTarget:
		return m, nil
	case Generating:
		return m, nil
	case Finished:
		return m, tea.Quit
	}

	return m, nil
}

func (m Model) View() string {
	return m.protoInput.View()
}

func InitModel() Model {
	return Model{
		protoInput:  protoinput.NewModel(),
		currentStep: InputSource,
	}
}
