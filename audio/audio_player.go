package audio

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

type Extension int

const (
	Mp3 Extension = iota
	Wav
)

var SupportFileExtensions = map[string]Extension{"mp3": Mp3, "wav": Wav}

type AudioPlayer struct {
	FilePath string
	Ctrl     *beep.Ctrl
}

func NewAudioPlayer(path string) (*AudioPlayer, error) {
	fileExt := findFileExtension(path)
	ext, ok := SupportFileExtensions[fileExt]
	if !ok {
		return nil, fmt.Errorf("unsupported audio format: .%s", fileExt)
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
	ctrl := &beep.Ctrl{Streamer: streamer, Paused: true}
	speaker.Play(beep.Seq(ctrl, beep.Callback(func() {
		file.Close()
		streamer.Close()
	})))

	return &AudioPlayer{FilePath: path, Ctrl: ctrl}, nil
}

func (p *AudioPlayer) Play() {
	p.Ctrl.Paused = false
}

func (p *AudioPlayer) Pause() {
	p.Ctrl.Paused = true
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

func findFileExtension(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			return s[i+1:]
		}
	}
	return ""
}
