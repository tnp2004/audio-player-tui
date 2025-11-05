package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type appState int

const (
	audioCategoryState appState = iota
	audioDestState
)

type Model struct {
	state         appState
	audioCategory AudioCategory
	audioDest     AudioDest
	terminal      struct{ width, height int }
}

func NewModel() Model {
	return Model{
		state:         audioCategoryState,
		audioCategory: newAudioCategory(),
		audioDest:     newAudioDest(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminal.width = msg.Width
		m.terminal.height = msg.Height
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	switch m.state {
	case audioCategoryState:
		return m.handleAudioCategoryKeyEvent(msg)
	case audioDestState:
		return m.handleAudioDestKeyEvent(msg)
	}

	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case audioCategoryState:
		return m.renderAudioCategoryView()
	case audioDestState:
		return m.renderAudioDestView()
	}

	return "empty view"
}
