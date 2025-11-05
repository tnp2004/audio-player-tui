package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tnp2004/audio-player-tui/model"
)

func main() {
	m := model.NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running app: %v", err)
	}
}
