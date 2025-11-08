package model

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tnp2004/audio-player-tui/audio"
)

type AudioPlayer struct {
	dest          string
	player        *audio.AudioPanel
	totalDuration time.Duration
	elapsedTime   time.Duration
}

type elapsedTimeMsg time.Time

func newAudioPlayer() AudioPlayer {
	return AudioPlayer{}
}

func elapsedTimeTicker() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return elapsedTimeMsg(t)
	})
}

func (m Model) setupAudioPlayer(dest string) (Model, error) {
	player, err := audio.NewAudioPanel(dest)
	if err != nil {
		return m, err
	}
	m.audioPlayer.dest = dest
	m.audioPlayer.player = player
	m.audioPlayer.totalDuration = m.audioPlayer.player.GetTotalDuration()

	return m, nil
}

func (m Model) handleAudioPlayerKeyEvent(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch key := msg.(type) {
	case tea.KeyMsg:
		switch key.Type {
		case tea.KeySpace:
			fallthrough
		case tea.KeyRunes:
			if key.String() == " " {
				if m.audioPlayer.player.IsPlaying() {
					m.audioPlayer.player.Pause()
					m.updateElapsedTime()
					return m, nil
				}

				m.audioPlayer.player.Play()
				return m, elapsedTimeTicker()
			}
		}
	case elapsedTimeMsg:
		m.updateElapsedTime()
		if m.audioPlayer.player.IsPlaying() {
			return m, elapsedTimeTicker()
		}
		return m, nil
	}

	return m, cmd
}

func (m *Model) updateElapsedTime() {
	m.audioPlayer.elapsedTime = m.audioPlayer.player.GetElapsedTime()
}

func (m Model) renderAudioPlayerView() string {
	return fmt.Sprintf("Audio Player View\n%s / %s", m.audioPlayer.elapsedTime, m.audioPlayer.totalDuration)
}
