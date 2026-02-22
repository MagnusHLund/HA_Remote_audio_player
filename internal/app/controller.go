package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/MagnusHLund/HA_Remote_audio_player/internal/audio"
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/config"
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/mqtt"
	paho "github.com/eclipse/paho.mqtt.golang"
)

// Controller contains the core application workflow.
type Controller struct {
	logger       *log.Logger
	mqttClient   *mqtt.Client
	audioPlayer  *audio.Playback
	audioConfigs []config.AudioConfig
	mqttConfig   *config.MQTTConfig
}

// NewController builds the controller with required dependencies.
func NewController(
	logger *log.Logger,
	mqttClient *mqtt.Client,
	audioPlayer *audio.Playback,
	audioConfigs []config.AudioConfig,
	mqttConfig *config.MQTTConfig,
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

	c.pairToHomeAssistant()

	err := c.mqttClient.Subscribe(mqtt.CommandTopic(), func(client paho.Client, msg paho.Message) {
		c.handleMqttMessage(msg.Payload())
	})
	if err != nil {
		return err
	}

	defer c.mqttClient.Disconnect()

	// TODO: Gracefully shutdown on interrupt signal
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
	c.logger.Println("Shutting down...")

	return nil
}

func (c *Controller) handleMqttMessage(payload []byte) {
	audioPath, err := c.mapMqttPayloadToAudioPath(string(payload))
	if err != nil {
		c.logger.Printf("Error mapping MQTT payload to audio path: %v", err)
		return
	}

	c.logger.Printf("Mapped MQTT payload to audio path: %s", audioPath)

	if !c.audioPlayer.DoesFileExist(audioPath) {
		c.logger.Printf("Audio file does not exist: %s", audioPath)
		return
	}

	if !c.audioPlayer.IsSupportedAudioFile(audioPath) {
		c.logger.Printf("Unsupported audio file type: %s", audioPath)
		return
	}

	c.mqttClient.Publish(mqtt.StateTopic(), "playing")

	if err := c.audioPlayer.PlayAudioFile(audioPath); err != nil {
		c.logger.Printf("Error playing audio file: %v", err)
	}

	c.mqttClient.Publish(mqtt.StateTopic(), "idle")
}

func (c *Controller) mapMqttPayloadToAudioPath(payload string) (string, error) {
	for _, audioConfig := range c.audioConfigs {
		if audioConfig.Name == payload {
			return audioConfig.Path, nil
		}
	}

	return "", fmt.Errorf("Audio file not found: %s", payload)
}

func (c *Controller) pairToHomeAssistant() {
	for _, entry := range c.audioConfigs {
		objectID := entry.Name
		discoveryTopic := fmt.Sprintf("homeassistant/button/%s/%s/config", c.mqttConfig.DeviceID, objectID)

		payload := map[string]any{
			"name":          entry.Name,
			"unique_id":     fmt.Sprintf("%s-%s", c.mqttConfig.DeviceID, objectID),
			"command_topic": mqtt.CommandTopic(),
			"state_topic":   mqtt.StateTopic(),
			"payload_press": entry.Name,
			"device": map[string]any{
				"identifiers":  []string{c.mqttConfig.DeviceID},
				"name":         c.mqttConfig.DeviceName,
				"manufacturer": "Magnus H Lund",
				"model":        "remote-audio-player",
			},
		}

		jsonPayload, _ := json.Marshal(payload)
		_ = c.mqttClient.Publish(discoveryTopic, string(jsonPayload))
	}
}
