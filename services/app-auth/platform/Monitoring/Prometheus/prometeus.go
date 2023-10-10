package Prometheus

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

type Metrics struct {
	HitsTotal prometheus.Counter
	Hits      *prometheus.CounterVec
	Times     *prometheus.HistogramVec
	cpuTemp   prometheus.Gauge
}

func (m Metrics) ObserveResponseTime(statusCode int, method, path string, duration float64) {
	m.HitsTotal.Inc()
	m.Hits.WithLabelValues(strconv.Itoa(statusCode), method, path).Inc()
	m.Times.WithLabelValues(strconv.Itoa(statusCode), method, path).Observe(duration)
}

func (m Metrics) IncHits(statusCode int, method, path string) {
	m.Hits.WithLabelValues(strconv.Itoa(statusCode), method, path).Inc()

}

func NewMetrics(HitsTotal prometheus.Counter, Hits *prometheus.CounterVec, Times *prometheus.HistogramVec, cpuTemp prometheus.Gauge) metrics.MetricsHandler {
	return &Metrics{
		HitsTotal: HitsTotal,
		Hits:      Hits,
		Times:     Times,
		cpuTemp:   cpuTemp,
	}

}
