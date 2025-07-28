package processor

import "math"

var mu sync.Mutex
var lastCap float64

func DetectFlashCrash(marketCap float64) (alert bool) {
	mu.Lock()
	defer mu.Unlock()

	if lastCap == 0 {
		lastCap = marketCap
		return false
	}
	drop := (lastCap - marketCap) / lastCap
	lastCap = marketCap
	return drop > 0.3 // аларм, если просадка >30%
}
