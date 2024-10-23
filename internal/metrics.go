package internal

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Definição de outras métricas se necessário
var (
	viewersGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_viewers",
		Help: "Number of current viewers",
	})
)

func init() {
	prometheus.MustRegister(viewersGauge)
}
