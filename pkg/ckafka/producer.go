package ckafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)


func Producer(borkers []string, topic string, dialer *kafka.Dialer, key []byte, value []byte) (err error) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  borkers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		Dialer:   dialer,
	})
	if err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   key,
			Value: value,
		}); err != nil {
		return err
	}
	defer w.Close()
	return err
}