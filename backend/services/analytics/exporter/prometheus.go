package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	dauGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{Name: "game_dau_total", Help: "Daily active users"},
	)
)

func init() {
	prometheus.MustRegister(dauGauge)
}

func UpdateDAU(val uint64) { dauGauge.Set(float64(val)) }
