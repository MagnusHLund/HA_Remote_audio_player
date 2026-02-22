package audio

import (
	"fmt"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
)

type Playback struct {
}

func NewPlayback() (*Playback, error) {
	return &Playback{}, nil
}

func (p *Playback) PlayAudioFile(path string) error {
	ext, err := GetFileExtension(path)
	if err != nil {
		return err
	}

	switch ext {
	case ".mp3":
		return playMp3AudioFile(path)
	case ".wav":
		return playWavAudioFile(path)
	case ".ogg":
		return playOggAudioFile(path)
	case ".flac":
		return playFlacAudioFile(path)
	default:
		return fmt.Errorf("unsupported audio format: %s", ext)
	}
}

func (p *Playback) initSpeaker(format beep.Format) error {
	return speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
}

func (p *Playback) play(streamer beep.StreamSeekCloser, format beep.Format) error {
	defer streamer.Close()

	if err := initSpeaker(format); err != nil {
		return err
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	return nil
}

func (p *Playback) playMp3AudioFile(path string) error {
	r, err := openSource(path)
	if err != nil {
		return err
	}
	defer r.Close()

	streamer, format, err := mp3.Decode(r)
	if err != nil {
		return err
	}

	return play(streamer, format)
}

func (p *Playback) playWavAudioFile(path string) error {
	r, err := openSource(path)
	if err != nil {
		return err
	}
	defer r.Close()

	streamer, format, err := wav.Decode(r)
	if err != nil {
		return err
	}

	return play(streamer, format)
}

func (p *Playback) playOggAudioFile(path string) error {
	r, err := openSource(path)
	if err != nil {
		return err
	}
	defer r.Close()

	streamer, format, err := vorbis.Decode(r)
	if err != nil {
		return err
	}

	return play(streamer, format)
}

func (p *Playback) playFlacAudioFile(path string) error {
	r, err := openSource(path)
	if err != nil {
		return err
	}
	defer r.Close()

	streamer, format, err := flac.Decode(r)
	if err != nil {
		return err
	}

	return play(streamer, format)
}
