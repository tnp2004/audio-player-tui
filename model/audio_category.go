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
	{"YouTube", "Enter a YouTube video URL", YouTube},
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
			m.audioCategory.moveCursorUp()
		case tea.KeyDown:
			m.audioCategory.moveCursorDown()
		case tea.KeyEnter:
			m.audioCategory.selectCurrent()
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

func (a *AudioCategory) moveCursorUp() {
	if a.cursor > 0 {
		a.cursor--
	} else {
		a.cursor = len(a.options) - 1
	}
}

func (a *AudioCategory) moveCursorDown() {
	if a.cursor < len(a.options)-1 {
		a.cursor++
	} else {
		a.cursor = 0
	}
}

func (a *AudioCategory) selectCurrent() {
	a.selected = a.cursor
	a.cursor = defaultAudioCategoryCursor
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
