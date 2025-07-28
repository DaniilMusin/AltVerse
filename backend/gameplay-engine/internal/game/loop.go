package game

import "context"

func (e *Engine) Tick(ctx context.Context) error {
    // 1. Идём в Kafka / Redis за новыми командами игроков
    cmds, err := e.commandRepo.Fetch(ctx)
    if err != nil { return err }

    // 2. Проводим бизнес-логику
    for _, c := range cmds {
        if err := e.applyCommand(ctx, c); err != nil {
            e.logger.Warn("command failed", "id", c.ID, "err", err)
        }
    }

    // 3. Рассылаем события и сохраняем снапшот
    return e.eventBus.Flush(ctx)
}
