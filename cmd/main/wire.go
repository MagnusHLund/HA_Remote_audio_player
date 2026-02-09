//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/MagnusHLund/HA_Remote_audio_player/internal/app"
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/audio"
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/config"
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/mqtt"
)

// InitializeApp creates and returns the application with all dependencies
func InitializeApp(configPath string) (*app.App, error) {
	wire.Build(
		config.NewConfig,
		audio.NewPlayer,
		mqtt.NewClient,
		app.NewApp,
	)
	return nil, nil
}
