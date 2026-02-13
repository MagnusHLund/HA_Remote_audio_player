package app

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/MagnusHLund/HA_Remote_audio_player/internal/audio"
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/config"
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/mqtt"
)

// Controller contains the core application workflow.
type Controller struct {
	logger       *log.Logger
	mqttClient   *mqtt.Client
	audioPlayer  *audio.Player
	audioConfigs []config.AudioConfig
}

// NewController builds the controller with required dependencies.
func NewController(
	logger *log.Logger,
	mqttClient *mqtt.Client,
	audioPlayer *audio.Player,
	audioConfigs []config.AudioConfig,
) *Controller {
	return &Controller{
		logger:       logger,
		mqttClient:   mqttClient,
		audioPlayer:  audioPlayer,
		audioConfigs: audioConfigs,
	}
}

// Run starts the controller lifecycle.
func (c *Controller) Run() error {
	if err := c.mqttClient.Connect(); err != nil {
		return err
	}
	defer c.mqttClient.Disconnect()

	c.logger.Printf("loaded %d audio entries", len(c.audioConfigs))
	c.logger.Printf("audio player ready with sample rate %d", c.audioPlayer.SampleRate())

	// TODO: Gracefully shutdown on interrupt signal
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
	c.logger.Println("Shutting down...")

	return nil
}
