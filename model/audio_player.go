package model

import (
	"fmt"
	"math"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tnp2004/audio-player-tui/audio"
	"github.com/tnp2004/audio-player-tui/ui"
)

type AudioPlayer struct {
	initialized   bool
	dest          string
	player        *audio.AudioPanel
	totalDuration time.Duration
	elapsedTime   time.Duration
}

type elapsedTimeMsg time.Time

func newAudioPlayer() AudioPlayer {
	return AudioPlayer{initialized: false}
}

func elapsedTimeTicker(elapsedTime time.Duration) tea.Cmd {
	tickDuration := getElapsedDurationTick(elapsedTime)
	return tea.Tick(tickDuration, func(t time.Time) tea.Msg {
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

	if !m.audioPlayer.initialized {
		m.audioPlayer.initialized = true
		return m, elapsedTimeTicker(m.audioPlayer.elapsedTime)
	}

	switch key := msg.(type) {
	case tea.KeyMsg:
		switch key.Type {
		case tea.KeySpace:
			fallthrough
		case tea.KeyRunes:
			switch key.String() {
			case " ":
				if m.audioPlayer.player.IsPlaying() {
					m.audioPlayer.player.Pause()
					m = m.updateElapsedTime()
					return m, nil
				}

				m.audioPlayer.player.Play()
				return m, elapsedTimeTicker(m.audioPlayer.elapsedTime)
			case "r":
				m.audioPlayer = newAudioPlayer()
				m, m.audioDest.err = m.setupAudioPlayer(m.audioDest.input.Value())
				return m, nil
			}
		}
	case elapsedTimeMsg:
		m = m.updateElapsedTime()
		if m.audioPlayer.player.IsPlaying() {
			return m, elapsedTimeTicker(m.audioPlayer.elapsedTime)
		}
		return m, nil
	}

	return m, cmd
}

func (m Model) updateElapsedTime() Model {
	m.audioPlayer.elapsedTime = m.audioPlayer.player.GetElapsedTime()
	return m
}

func (m Model) renderAudioPlayerView() string {
	filename := m.getDestFilename()
	title := fmt.Sprintf("Audio Player: %s", filename)

	indicator := fmt.Sprintf("%s %s / %s",
		m.getAudioPlayerStatus(),
		m.audioPlayer.elapsedTime.Round(time.Second).String(),
		m.audioPlayer.totalDuration.Round(time.Second).String())
	progressBar := drawProgressBar(m.terminal.width, m.audioPlayer.elapsedTime, m.audioPlayer.totalDuration)
	content := lipgloss.JoinVertical(lipgloss.Center,
		title,
		progressBar,
		indicator)

	return lipgloss.PlaceHorizontal(m.terminal.width, lipgloss.Center, content)
}

func (m Model) getAudioPlayerStatus() string {
	if m.audioPlayer.elapsedTime == m.audioPlayer.totalDuration {
		return "Endded"
	}

	if m.audioPlayer.player.IsPlaying() {
		return "Playing"
	} else {
		return "Paused"
	}
}

func (m Model) getDestFilename() string {
	dest := m.audioPlayer.dest
	for i := len(dest) - 1; i >= 0; i-- {
		v := dest[i]
		if v == '/' || v == '\\' {
			return dest[i+1:]
		}
	}

	return ""
}

func drawProgressBar(terminalWidth int, elapsedTime, totalDuration time.Duration) string {
	padding := (terminalWidth * 20) / 100
	width := terminalWidth - padding
	ratio := elapsedTime.Seconds() / totalDuration.Seconds()
	if ratio > 1 {
		ratio = 1
	}

	indicatorSymbol := "o"
	progressSymbol := "-"
	pos := int(math.Round(ratio * float64(width-1)))
	left := ui.ProgressBarStyle.Render(strings.Repeat(progressSymbol, pos))
	right := ui.ProgressBarStyle.Render(strings.Repeat(progressSymbol, width-pos-1))

	return fmt.Sprintf("[%s%s%s]", left, indicatorSymbol, right)
}

func getElapsedDurationTick(elapsedTime time.Duration) time.Duration {
	ms := elapsedTime % time.Second
	if ms == 0 {
		return time.Second
	}
	return time.Second - ms
}
