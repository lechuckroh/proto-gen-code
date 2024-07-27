package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lechuckroh/protogencode/internal/app/protogen"
	"log/slog"
)

var logger *slog.Logger

func startTUI() {
	p := tea.NewProgram(protogen.InitModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		logger.Error("failed to initialize model", err)
	}
}

func main() {
	startTUI()
}
