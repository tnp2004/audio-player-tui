package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AudioType struct {
	title    string
	types    []audioChoice
	cursor   int
	selected int
}

type audioChoice struct {
	title string
	desc  string
}

type audioType int

var (
	// Color code
	selectColor     = "#a983f7"
	descSelectColor = "#8d6dcf"

	// Style
	baseSelectStyle = lipgloss.NewStyle().
			Bold(true).
			BorderStyle(lipgloss.BlockBorder()).
			BorderForeground(lipgloss.Color(selectColor)).
			BorderLeft(true).
			PaddingLeft(1)
	selectStyle     = baseSelectStyle.Foreground(lipgloss.Color(selectColor))
	descSelectStyle = baseSelectStyle.Foreground(lipgloss.Color(descSelectColor))
	unselectStyle   = lipgloss.NewStyle().PaddingLeft(2)
)

const audioTypeTitle = "Select audio type"

const (
	unselect audioType = iota
	audioFile
	youtube
)

var audioTypes = []audioChoice{
	{"Audio File", "select an audio file from your local system"},
	{"YouTube", "enter a YouTube video URL"},
}

const defaultAudioTypeCursor int = 0
const unselectAudioType int = -1

func initAudioType() AudioType {
	return AudioType{
		title:    audioTypeTitle,
		types:    audioTypes,
		cursor:   defaultAudioTypeCursor,
		selected: unselectAudioType,
	}
}

func (m model) updateAudioTypeEvent(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

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
		m.audioType.cursor = defaultAudioTypeCursor
		m.state = audioDstInputState
	}

	return m, cmd
}

func (m model) selectAudioTypeView() string {
	var sections []string
	title := m.audioType.title + "\n"
	sections = append(sections, title)

	for i, c := range m.audioType.types {
		var title, desc string
		if i == m.audioType.cursor {
			title = selectStyle.Render(c.title)
			desc = descSelectStyle.Render(c.desc)
		} else {
			title = unselectStyle.Render(c.title)
			desc = unselectStyle.Render(c.desc)
		}

		item := lipgloss.JoinVertical(lipgloss.Left, title, desc)
		sections = append(sections, item)
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (m model) getAudioType() audioType {
	switch m.audioType.types[m.audioType.selected].title {
	case "Audio File":
		return audioFile
	case "YouTube":
		return youtube
	}

	return unselect
}
