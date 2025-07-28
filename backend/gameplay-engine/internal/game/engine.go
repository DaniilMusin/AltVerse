package game

import (
    "context"
    "log/slog"

    "altverse/gameplay-engine/internal/store"
)

type Command struct {
    ID      string
    Payload string
}

type CommandRepo interface {
    Fetch(ctx context.Context) ([]Command, error)
}

type EventBus interface {
    Flush(ctx context.Context) error
}

type Engine struct {
    commandRepo CommandRepo
    eventBus    EventBus
    logger      *slog.Logger
}

func NewEngine(db *store.Postgres, cache *store.Redis, logger *slog.Logger) *Engine {
    return &Engine{logger: logger}
}

func (e *Engine) applyCommand(ctx context.Context, c Command) error {
    return nil
}
