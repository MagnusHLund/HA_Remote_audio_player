package audio

import (
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

// Player handles audio playback.
type Player struct {
	sampleRate beep.SampleRate
}

// NewPlayer creates a player with a fixed sample rate.
func NewPlayer() (*Player, error) {
	sr := beep.SampleRate(44100)
	if err := speaker.Init(sr, sr.N(time.Second/10)); err != nil {
		return nil, err
	}

	return &Player{sampleRate: sr}, nil
}
