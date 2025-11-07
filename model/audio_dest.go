package model

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tnp2004/audio-player-tui/ui"
)

type AudioDest struct {
	title string
	input textinput.Model
}

func newAudioDest() AudioDest {
	ti := textinput.New()
	ti.Placeholder = "Enter destination..."
	ti.Width = 50
	ti.CharLimit = 100
	ti.Focus()

	return AudioDest{
		title: "Audio Destination",
		input: ti,
	}
}

func (m Model) handleAudioDestKeyEvent(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch key := msg.(type) {
	case tea.KeyMsg:
		switch key.Type {
		case tea.KeyEnter:
			if err := m.setupAudioPlayer(m.audioDest.input.Value()); err != nil {
				log.Fatal(err)
			}
			m.state = audioPlayerState
			return m, nil
		}
	}

	m.audioDest.input, cmd = m.audioDest.input.Update(msg)
	return m, cmd
}

func (m Model) renderAudioDestView() string {
	title := renderAudioTitle(m.getAudioKind())
	inputView := m.audioDest.input.View()
	inputBox := ui.TextBoxStyle.Render(inputView)

	content := fmt.Sprintf("%s\n%s", title, inputBox)

	return lipgloss.PlaceHorizontal(
		m.terminal.width,
		lipgloss.Center,
		content,
	)
}

func renderAudioTitle(kind AudioKind) string {
	switch kind {
	case AudioFile:
		return "Audio File"
	case YouTube:
		return "YouTube"
	}

	return "Unselect audio kind"
}
