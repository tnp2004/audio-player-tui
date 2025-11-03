package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type AudioType struct {
	types    []string
	cursor   int
	selected int
}

var audioTypes = []string{"audio file", "youtube"}

const defaultAudioTypeCursor int = 0
const unselectAudioType int = -1

func (m model) updateAudioTypeEvent(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.audioType.cursor > 0 {
				m.audioType.cursor--
			} else {
				m.audioType.cursor = len(m.audioType.types) - 1
			}
		case tea.KeyDown:
			if m.audioType.cursor < len(m.audioType.types)-1 {
				m.audioType.cursor++
			} else {
				m.audioType.cursor = 0
			}
		case tea.KeyEnter:
			m.audioType.selected = m.audioType.cursor
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	return m, cmd
}

func (m model) selectAudioTypeView() string {
	s := "Select audio type"
	for i, v := range m.audioType.types {
		prefix := "\n "
		if i == m.audioType.cursor {
			prefix = "\n>"
		}
		s += fmt.Sprintf("%s %s", prefix, v)
	}

	if m.audioType.selected != unselectAudioType {
		s += fmt.Sprintf("\nSelected: %s", m.audioType.types[m.audioType.selected])
	}

	return s
}
