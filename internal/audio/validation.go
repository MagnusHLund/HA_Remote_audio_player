package audio

import (
	"net/http"
	"os"
)

func DoesFileExist(path string) bool {
	if IsHTTP(path) {
		resp, err := http.Head(path)
		return err == nil && resp.StatusCode == http.StatusOK
	}

	_, err := os.Stat(path)
	return err == nil
}

func IsSupportedAudioFile(path string) bool {
	_, err := GetFileExtension(path)
	if err != nil {
		return false
	}

	return true
}
