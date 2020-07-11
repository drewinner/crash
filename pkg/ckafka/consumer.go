package ckafka

import (
	"context"
	"fmt"
)

func consumer(borkers []string, topic string, groupId string, dialer *kafka.Dialer) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  borkers,
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 10e1, //10KB
		MaxBytes: 10e6, //10MB
		Dialer:   dialer,
	})
	defer func() {
		_ = r.Close()
	}()
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("err %+v", err)
			continue
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
