package app

import (
	"log"
	"os"
)

// NewLogger provides a standard application logger.
func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.LUTC)
}
