package consumer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Handler func(msg kafka.Message)

func Run(ctx context.Context, brokers []string, group, topic string, handle Handler) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: group,
		Topic:   topic,
	})
	defer r.Close()

	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			return err
		}
		handle(m)
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Printf("commit error: %v", err)
		}
	}
}
