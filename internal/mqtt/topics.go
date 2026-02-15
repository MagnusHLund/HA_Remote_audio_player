package mqtt

import (
	"fmt"
	"sync"

	"github.com/MagnusHLund/HA_Remote_audio_player/internal/config"
)

type Topics struct {
	Discovery string
	State     string
	Command   string
}

var (
	topicsOnce   sync.Once
	cachedTopics Topics
)

func NewTopics(cfg *config.MQTTConfig) Topics {
	base := fmt.Sprintf("homeassistant/media_player/%s", cfg.DeviceID)

	return Topics{
		State:   fmt.Sprintf("%s/state", base),
		Command: fmt.Sprintf("%s/set", base),
	}
}

func StateTopic() string {
	return loadTopics().State
}

func CommandTopic() string {
	return loadTopics().Command
}

func loadTopics() Topics {
	topicsOnce.Do(func() {
		cachedTopics = NewTopics(config.NewMQTTConfig())
	})

	return cachedTopics
}
