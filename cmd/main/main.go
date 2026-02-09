package main

import (
	"flag"
	"log"
)

func main() {
	configPath := flag.String("config", "configs/audio.json", "Path to sounds JSON file")
	flag.Parse()

	application, err := InitializeApp(*configPath)
	if err != nil {
		log.Fatalf("init failed: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("run failed: %v", err)
	}
}
