package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AudioDstInput struct {
	title string
	input textinput.Model
}

var textinputStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())

func initAudioDstInput() AudioDstInput {
	ti := textinput.New()
	ti.Placeholder = "enter destination..."
	ti.Width = 50
	ti.CharLimit = 100
	ti.Focus()

	return AudioDstInput{
		input: ti,
	}
}

func (m model) updateAudioDstInputEvent(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, nil
		}
	}

	m.audioDstInput.input, cmd = m.audioDstInput.input.Update(msg)
	return m, cmd
}

func (m model) inputAudioDstView() string {
	var title string
	switch m.getAudioType() {
	case audioFile:
		title = "Audio File"
	case youtube:
		title = "YouTube"
	}
	inputBox := textinputStyle.Render(m.audioDstInput.input.View())
	str := fmt.Sprintf("%s\n%s", title, inputBox)
	return lipgloss.PlaceHorizontal(
		m.terminal.width,
		lipgloss.Center,
		str,
	)
}
