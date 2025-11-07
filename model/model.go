package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type appState int

const (
	audioCategoryState appState = iota
	audioDestState
	audioPlayerState
)

type Model struct {
	state         appState
	audioCategory AudioCategory
	audioDest     AudioDest
	audioPlayer   AudioPlayer
	terminal      struct{ width, height int }
}

func NewModel() Model {
	return Model{
		state:         audioCategoryState,
		audioCategory: newAudioCategory(),
		audioDest:     newAudioDest(),
		audioPlayer:   newAudioPlayer(),
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
	case audioPlayerState:
		return m.handleAudioPlayerKeyEvent(msg)
	}

	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case audioCategoryState:
		return m.renderAudioCategoryView()
	case audioDestState:
		return m.renderAudioDestView()
	case audioPlayerState:
		return m.renderAudioPlayerView()
	}

	return "empty view"
}
