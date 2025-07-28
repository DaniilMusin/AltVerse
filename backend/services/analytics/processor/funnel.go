package processor

import "sync/atomic"

var dau uint64

func TrackSession(playerID string) {
	// Упрощённый счётчик DAU ― продакшен-версия хранит uniq-id в Redis SET
	atomic.AddUint64(&dau, 1)
}

func DAU() uint64 { return atomic.LoadUint64(&dau) }
