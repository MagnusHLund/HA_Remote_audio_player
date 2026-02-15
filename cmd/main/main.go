package main

import (
	"log"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := log.New(os.Stderr, "", log.LstdFlags|log.LUTC)

	application, err := InitializeApp()
	if err != nil {
		logger.Printf("init failed: %v", err)
		return 1
	}

	if err := application.Run(); err != nil {
		logger.Printf("run failed: %v", err)
		return 1
	}

	return 0
}
