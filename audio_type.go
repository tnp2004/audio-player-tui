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

var (
	// Color code
	selectColor     = "#a983f7"
	descSelectColor = "#8d6dcf"

	// Style
	selectStyle = lipgloss.NewStyle().
			Bold(true).
			BorderStyle(lipgloss.BlockBorder()).
			BorderForeground(lipgloss.Color(selectColor)).
			BorderLeft(true).
			Foreground(lipgloss.Color(selectColor)).
			PaddingLeft(1)
	descSelectStyle = lipgloss.NewStyle().
			Bold(true).
			BorderStyle(lipgloss.BlockBorder()).
			BorderForeground(lipgloss.Color(selectColor)).
			BorderLeft(true).
			Foreground(lipgloss.Color(descSelectColor)).
			PaddingLeft(1)
	unselectStyle = lipgloss.NewStyle().
			PaddingLeft(2)
)

const audioTypeTitle = "Select audio type"

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
			m.audioType.cursor = defaultAudioTypeCursor
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	return m, cmd
}
func (m model) selectAudioTypeView() string {
	var sections []string

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

	list := lipgloss.JoinVertical(lipgloss.Left, sections...)

	return lipgloss.JoinVertical(lipgloss.Left,
		m.audioType.title+"\n",
		list,
	)
}
