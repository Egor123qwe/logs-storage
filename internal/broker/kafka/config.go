package kafka

import "github.com/spf13/viper"

type config struct {
	brokers []string

	username string
	password string
}

func newConfig() config {
	return config{
		brokers: viper.GetStringSlice("broker.URLs"),

		username: viper.GetString("broker.username"),
		password: viper.GetString("broker.password"),
	}
}
