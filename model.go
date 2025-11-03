package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type State uint

const (
	audioTypeState State = iota
)

type model struct {
	state     State
	audioType AudioType
}

func initModel() model {
	return model{
		state: audioTypeState,
		audioType: AudioType{
			types:    audioTypes,
			cursor:   defaultAudioTypeCursor,
			selected: unselectAudioType,
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.state {
	case audioTypeState:
		return m.updateAudioTypeEvent(msg)
	}

	return m, cmd
}

func (m model) View() string {
	switch m.state {
	case audioTypeState:
		return m.selectAudioTypeView()
	}

	return "empty view"
}
