package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tnp2004/audio-player-tui/ui"
)

type AudioCategory struct {
	title    string
	options  []AudioOption
	cursor   int
	selected int
}

type AudioOption struct {
	title string
	desc  string
	kind  AudioKind
}

type AudioKind int

const (
	Unselect AudioKind = iota
	AudioFile
	YouTube
)

const (
	audioCategoryTitle         = "Select audio type"
	defaultAudioCategoryCursor = 0
	unselectAudioCategory      = -1
)

var defaultAudioOptions = []AudioOption{
	{"Audio File", "Select an audio file from your local system", AudioFile},
}

func newAudioCategory() AudioCategory {
	return AudioCategory{
		title:    audioCategoryTitle,
		options:  defaultAudioOptions,
		cursor:   defaultAudioCategoryCursor,
		selected: unselectAudioCategory,
	}
}

func (m Model) handleAudioCategoryKeyEvent(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch key := msg.(type) {
	case tea.KeyMsg:
		switch key.Type {
		case tea.KeyUp:
			m = m.moveCursorUp()
		case tea.KeyDown:
			m = m.moveCursorDown()
		case tea.KeyEnter:
			m = m.selectCurrent()
			m.state = audioDestState
		}
	}

	return m, nil
}

func (m Model) renderAudioCategoryView() string {
	var sections []string
	sections = append(sections, m.audioCategory.title+"\n")

	for i, opt := range m.audioCategory.options {
		item := renderAudioOption(opt, i == m.audioCategory.cursor)
		sections = append(sections, item)
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (m Model) moveCursorUp() Model {
	if m.audioCategory.cursor > 0 {
		m.audioCategory.cursor--
	} else {
		m.audioCategory.cursor = len(m.audioCategory.options) - 1
	}
	return m
}

func (m Model) moveCursorDown() Model {
	if m.audioCategory.cursor < len(m.audioCategory.options)-1 {
		m.audioCategory.cursor++
	} else {
		m.audioCategory.cursor = 0
	}
	return m
}

func (m Model) selectCurrent() Model {
	m.audioCategory.selected = m.audioCategory.cursor
	m.audioCategory.cursor = defaultAudioCategoryCursor
	return m
}

func (m Model) getAudioKind() AudioKind {
	if m.audioCategory.selected < 0 || m.audioCategory.selected >= len(m.audioCategory.options) {
		return Unselect
	}
	return m.audioCategory.options[m.audioCategory.selected].kind
}

func renderAudioOption(opt AudioOption, isSelected bool) string {
	var titleStyle, descStyle lipgloss.Style
	if isSelected {
		titleStyle, descStyle = ui.SelectStyle, ui.DescSelectStyle
	} else {
		titleStyle, descStyle = ui.UnselectStyle, ui.UnselectStyle
	}

	title := titleStyle.Render(opt.title)
	desc := descStyle.Render(opt.desc)
	return lipgloss.JoinVertical(lipgloss.Left, title, desc)
}
