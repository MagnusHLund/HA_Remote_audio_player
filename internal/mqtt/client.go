package mqtt

import (
	"time"

	"github.com/MagnusHLund/HA_Remote_audio_player/internal/config"
	"github.com/MagnusHLund/HA_Remote_audio_player/internal/mqtt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/wire"
)

type Client struct {
	client mqtt.Client
	cfg    *config.MQTTConfig
}

var ProviderSet = wire.NewSet(config.NewMQTTConfig,
	NewClient,
)

func NewClient(cfg *config.MQTTConfig) *client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.Broker)
	opts.SetClientID(cfg.ClientId)

	if cfg.Username != "" {
		opts.SetUsername(cfg.Username)
		opts.SetPassword(cfg.Password)
	}

	opts.AutoReconnect = true
	opts.ConnectRetry = true
	opts.ConnectRetryInterval = 10 * time.Second

	return &Client{
		client: mqtt.NewClient(opts),
		cfg:    cfg,
	}
}

func (c *Client) Connect() error {
	token := c.client.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *Client) Subscribe(topic string, handler mqtt.MessageHandler) error {
	token := c.client.Subscribe(topic, 1, handler)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *Client) Publish(topic string, payload string) error {
	token := c.client.Publish(topic, 1, false, payload)
	token.Wait()
	return token.Error()
}

func (c *Client) Disconnect() {
	c.client.Disconnect()
}
