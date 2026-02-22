package audio

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func OpenSource(path string) (io.ReadCloser, error) {
	if IsHTTP(path) {
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

func GetFileExtension(path string) (string, error) {
	validExtensions := [4]string{".mp3", ".wav", ".ogg", ".flac"}
	fileExtension := filepath.Ext(path)

	for _, validExt := range validExtensions {
		if fileExtension == validExt {
			return fileExtension, nil
		}
	}

	return "", fmt.Errorf("invalid audio file extension: %s", fileExtension)
}

func IsHTTP(path string) bool {
	return len(path) > 4 && (path[:7] == "http://" || path[:8] == "https://")
}
