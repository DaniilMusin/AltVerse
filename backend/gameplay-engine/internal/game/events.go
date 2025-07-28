package game

import (
	"context"

	"github.com/google/uuid"
)

type Topic string

const (
	TopicPlayerUpdated   Topic = "player.updated"
	TopicBlockchainSpawn Topic = "blockchain.spawn"
)

type GameEvent struct {
	ID        string `json:"id"`
	Topic     Topic  `json:"topic"`
	Payload   any    `json:"payload"`
	Timestamp int64  `json:"ts"`
}

type EventBus interface {
	Publish(ctx context.Context, ev GameEvent) error
	Flush(ctx context.Context) error
}

// ---------------------------------------------------------------------------

func (e *Engine) emit(ctx context.Context, topic Topic, payload any) {
	ev := GameEvent{
		ID:        uuid.NewString(),
		Topic:     topic,
		Payload:   payload,
		Timestamp: e.clock.Now().UnixMilli(),
	}
	_ = e.eventBus.Publish(ctx, ev) // в engine.Tick() будет e.eventBus.Flush
}
