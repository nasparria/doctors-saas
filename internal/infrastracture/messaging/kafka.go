// internal/infrastructure/messaging/kafka.go
package messaging

import (
	"context"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	Writer *kafka.Writer
	Reader *kafka.Reader
}

func NewKafkaClient(brokers []string, topic, groupID string) (*KafkaClient, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return &KafkaClient{
		Writer: writer,
		Reader: reader,
	}, nil
}

func (k *KafkaClient) ProduceMessage(ctx context.Context, key, value []byte) error {
	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	return k.Writer.WriteMessages(ctx, message)
}

func (k *KafkaClient) ConsumeMessages(ctx context.Context, handler func([]byte) error) error {
	for {
		m, err := k.Reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		if err := handler(m.Value); err != nil {
			return err
		}
	}
}

func (k *KafkaClient) Close() error {
	if err := k.Writer.Close(); err != nil {
		return err
	}
	return k.Reader.Close()
}
