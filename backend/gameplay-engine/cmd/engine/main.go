package main

import (
    "context"
    "log/slog"
    "os"
    "time"

    "altverse/gameplay-engine/internal/game"
    "altverse/gameplay-engine/internal/network"
    "altverse/gameplay-engine/internal/store"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    db := store.MustPostgres(os.Getenv("PG_DSN"))
    cache := store.MustRedis(os.Getenv("REDIS_URL"))

    engine := game.NewEngine(db, cache, logger)

    srv := network.NewGRPCServer(engine, logger)
    go srv.Start(os.Getenv("GRPC_ADDR"))

    ticker := time.NewTicker(240 * time.Millisecond)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            if err := engine.Tick(ctx); err != nil {
                logger.Error("tick failed", "err", err)
            }
        case <-ctx.Done():
            logger.Warn("shutting down")
            srv.Stop()
            return
        }
    }
}
