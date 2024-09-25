package kafka

import (
	"github.com/spf13/viper"
)

type consumer struct {
	topic   string
	groupID string
}

type config struct {
	consumers []consumer
}

func newConfig() config {
	config := config{}

	logs := consumer{
		topic:   viper.GetString("broker.consumer.logs.topic"),
		groupID: viper.GetString("broker.consumer.logs.group_id"),
	}

	config.consumers = append(config.consumers, logs)

	return config
}
