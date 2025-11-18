package game

import (
    "context"
    "log/slog"
    "time"

    "altverse/gameplay-engine/internal/store"
)

type Command struct {
    ID      string
    Payload string
}

type CommandRepo interface {
    Fetch(ctx context.Context) ([]Command, error)
}

type Clock interface {
    Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }

type Engine struct {
    commandRepo CommandRepo
    eventBus    EventBus
    logger      *slog.Logger
    clock       Clock
}

func NewEngine(db *store.PG, cache *store.RedisCache, logger *slog.Logger) *Engine {
    return &Engine{
        logger: logger,
        clock:  RealClock{},
    }
}

func (e *Engine) applyCommand(ctx context.Context, c Command) error {
    return nil
}
