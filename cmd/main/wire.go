package main

import (
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/app"
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/audio"
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/config"
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/mqtt"
)

// InitializeApp creates and returns the application with all dependencies
func InitializeApp(configPath string) (*app.App, error) {
	logger := app.NewLogger()

	audioConfigs, err := config.NewAudioConfig(configPath)
	if err != nil {
		return nil, err
	}

	mqttCfg := config.NewMQTTConfig()
	mqttClient := mqtt.NewClient(mqttCfg)

	audioPlayer, err := audio.NewPlayer()
	if err != nil {
		return nil, err
	}

	controller := app.NewController(logger, mqttClient, audioPlayer, audioConfigs)

	return app.NewApp(controller), nil
}
