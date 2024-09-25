package kafka

import (
	"fmt"
	"time"

	"github.com/op/go-logging"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
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

	mechanism, err := scram.Mechanism(scram.SHA256, config.username, config.password)
	if err != nil {
		return nil, fmt.Errorf("failed to create scram mechanism: %w", err)
	}

	dialer := &kafka.Dialer{
		Timeout:       dialerTimeout,
		DualStack:     true,
		SASLMechanism: mechanism,
	}

	srv := &service{
		config: config,
		dialer: dialer,
	}

	return srv, nil
}

func (s service) Producer(topic string) producer.Producer {
	return producer.New(s.dialer, s.config.brokers, topic)
}

func (s service) Consumer(topic string, groupID string) consumer.Consumer {
	return consumer.New(s.dialer, s.config.brokers, topic, groupID)
}
