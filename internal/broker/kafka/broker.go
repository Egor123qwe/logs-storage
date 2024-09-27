package kafka

import (
	"time"

	"github.com/op/go-logging"
	"github.com/segmentio/kafka-go"

	"github.com/Egor123qwe/logs-storage/internal/broker/kafka/consumer"
	"github.com/Egor123qwe/logs-storage/internal/broker/kafka/producer"
)

const (
	dialerTimeout = 10 * time.Second
)

var log = logging.MustGetLogger("kafka")

type Service interface {
	Producer(topic string) producer.Producer
	Consumer(topic string, groupID string) consumer.Consumer
}

type service struct {
	config config
	dialer *kafka.Dialer
}

func New() (Service, error) {
	config := newConfig()

	srv := &service{
		config: config,
	}

	return srv, nil
}

func (s service) Producer(topic string) producer.Producer {
	return producer.New(s.dialer, s.config.brokers, topic)
}

func (s service) Consumer(topic string, groupID string) consumer.Consumer {
	return consumer.New(s.dialer, s.config.brokers, topic, groupID)
}
