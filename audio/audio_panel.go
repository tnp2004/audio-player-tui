package audio

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

type Extension int

const (
	volumeStep            = 1.0
	playOnStart           = true
	Mp3         Extension = iota
	Wav
)

var SupportFileExtensions = map[string]Extension{"mp3": Mp3, "wav": Wav}

type AudioPanel struct {
	ctrl     *beep.Ctrl
	streamer beep.StreamSeekCloser
	format   beep.Format
	volume   *effects.Volume
}

func NewAudioPanel(path string) (*AudioPanel, error) {
	path = strings.Trim(path, `"`)
	fileExt, err := findFileExtension(path)
	if err != nil {
		return nil, err
	}
	ext, ok := SupportFileExtensions[fileExt]
	if !ok {
		return nil, fmt.Errorf("unsupported audio extension .%s", fileExt)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	streamer, format, err := DecodeAudioFile(file, ext)
	if err != nil {
		return nil, err
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	volume := &effects.Volume{
		Streamer: streamer,
		Base:     2.0,
		Volume:   0,
		Silent:   false,
	}

	ctrl := &beep.Ctrl{Streamer: volume, Paused: !playOnStart}
	speaker.Play(beep.Seq(ctrl, beep.Callback(func() {
		file.Close()
		streamer.Close()
	})))

	return &AudioPanel{
		ctrl:     ctrl,
		streamer: streamer,
		format:   format,
		volume:   volume,
	}, nil
}

func (p *AudioPanel) Play() {
	p.ctrl.Paused = false
}

func (p *AudioPanel) Pause() {
	p.ctrl.Paused = true
}

func (p AudioPanel) IsPlaying() bool {
	return !p.ctrl.Paused
}

func (p *AudioPanel) IncreaseVolume() {
	speaker.Lock()
	p.volume.Volume += volumeStep
	speaker.Unlock()
}

func (p *AudioPanel) DecreaseVolume() {
	speaker.Lock()
	p.volume.Volume -= volumeStep
	speaker.Unlock()
}

func (p AudioPanel) GetElapsedTime() time.Duration {
	pos := p.streamer.Position()
	return p.format.SampleRate.D(pos)
}

func (p AudioPanel) GetTotalDuration() time.Duration {
	pos := p.streamer.Len()
	return time.Duration(p.format.SampleRate.D(pos))
}

func DecodeAudioFile(file *os.File, ext Extension) (beep.StreamSeekCloser, beep.Format, error) {
	switch ext {
	case Mp3:
		return mp3.Decode(file)
	case Wav:
		return wav.Decode(file)
	}

	return nil, beep.Format{}, nil
}

func findFileExtension(s string) (string, error) {
	if !isFile(s) {
		return "", fmt.Errorf("invalid destination")
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			return s[i+1:], nil
		}
	}

	return "", nil
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
