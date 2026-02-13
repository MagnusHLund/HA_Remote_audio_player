package config

import (
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
)

type MQTTConfig struct {
	Broker      string
	Username    string
	Password    string
	ClientId    string
	DeviceID    string
	DeviceName  string
}

type AudioConfig struct {
	Name string
	Path string
}

func NewMQTTConfig() *MQTTConfig {
	return LoadMQTTConfig()
}

func LoadMQTTConfig() *MQTTConfig {
	_ = godotenv.Load()

	return &MQTTConfig{
		Broker:     getEnv("MQTT_BROKER", "tcp://localhost:1883"),
		Username:   getEnv("MQTT_USERNAME", ""),
		Password:   getEnv("MQTT_PASSWORD", ""),
		ClientId:   getEnv("MQTT_CLIENT_ID", "ha-remote-audio-player"),
		DeviceID:   getEnv("MQTT_DEVICE_ID", "ha-remote-audio-player"),
		DeviceName: getEnv("MQTT_DEVICE_NAME", "HA Remote Audio Player"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewAudioConfig(path string) ([]AudioConfig, error) {
	return LoadAudioConfig(path)
}

func LoadAudioConfig(path string) ([]AudioConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var audioConfigs []AudioConfig
	if err := json.Unmarshal(data, &audioConfigs); err != nil {
		return nil, err
	}

	return audioConfigs, nil
}
