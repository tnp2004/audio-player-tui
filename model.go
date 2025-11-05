package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type State uint

const (
	audioTypeState State = iota
	audioDstInputState
)

type Terminal struct {
	width  int
	height int
}

type model struct {
	state         State
	audioType     AudioType
	audioDstInput AudioDstInput
	terminal      Terminal
}

func initModel() model {
	return model{
		state:         audioTypeState,
		audioType:     initAudioType(),
		audioDstInput: initAudioDstInput(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminal.width = msg.Width
		m.terminal.height = msg.Height
	case tea.KeyMsg:
		// Default key
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

		// State key
		switch m.state {
		case audioTypeState:
			return m.updateAudioTypeEvent(msg)
		case audioDstInputState:
			return m.updateAudioDstInputEvent(msg)
		}
	}

	return m, cmd
}

func (m model) View() string {
	switch m.state {
	case audioTypeState:
		return m.selectAudioTypeView()
	case audioDstInputState:
		return m.inputAudioDstView()
	}

	return "empty view"
}
