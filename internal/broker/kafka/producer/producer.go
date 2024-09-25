package producer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Producer interface {
	Produce(ctx context.Context, m []byte) error
	Close() error
}

type producer struct {
	writer *kafka.Writer
}

func New(dialer *kafka.Dialer, brokers []string, topic string) Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Dialer:   dialer,
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return producer{writer: writer}
}

func (p producer) Produce(ctx context.Context, m []byte) error {
	err := p.writer.WriteMessages(ctx,
		kafka.Message{Value: m},
	)

	if err != nil {
		err = fmt.Errorf("failed to produce message: %w", err)
	}

	return err
}

func (p producer) Close() error {
	return p.writer.Close()
}
