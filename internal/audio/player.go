package audio

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
)

type Player struct{}

func NewPlayer() (*Player, error) {
	return &Player{}, nil
}

func (p *Player) PlayAudioFile(path string) error {
	ext, err := getFileExtension(path)
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

func openSource(path string) (io.ReadCloser, error) {
	// Remote file?
	if isHTTP(path) {
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("http error: %s", resp.Status)
		}
		return resp.Body, nil
	}

	// Local file
	return os.Open(path)
}

func initSpeaker(format beep.Format) error {
	return speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
}

func play(streamer beep.StreamSeekCloser, format beep.Format) error {
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

func playMp3AudioFile(path string) error {
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

func playWavAudioFile(path string) error {
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

func playOggAudioFile(path string) error {
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

func playFlacAudioFile(path string) error {
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

func (p *Player) DoesFileExist(path string) bool {
	if isHTTP(path) {
		resp, err := http.Head(path)
		return err == nil && resp.StatusCode == http.StatusOK
	}

	_, err := os.Stat(path)
	return err == nil
}

func isHTTP(path string) bool {
	return len(path) > 4 && (path[:7] == "http://" || path[:8] == "https://")
}

func (p *Player) IsSupportedAudioFile(path string) bool {
	_, err := getFileExtension(path)
	if err != nil {
		return false
	}

	return true
}

func getFileExtension(path string) (string, error) {
	validExtensions := [4]string{".mp3", ".wav", ".ogg", ".flac"}
	fileExtension := filepath.Ext(path)

	for _, validExt := range validExtensions {
		if fileExtension == validExt {
			return fileExtension, nil
		}
	}

	return "", fmt.Errorf("invalid audio file extension: %s", fileExtension)
}
