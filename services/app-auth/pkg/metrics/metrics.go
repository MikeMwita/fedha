package metrics

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/platform/Monitoring/Prometheus"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
)

type MetricsHandler interface {
	ObserveResponseTime(statusCode int, method, path string, duration float64)
	IncHits(statusCode int, method, path string)
}

// CreateMetrics creates and registers Prometheus metrics

func CreateMetrics(address string, name string) (MetricsHandler, error) {
	var metr Prometheus.Metrics
	metr.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: name + "_hits_total",
	})

	if err := prometheus.Register(metr.HitsTotal); err != nil {
		return nil, err
	}

	metr.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name + "_hits",
		},
		[]string{"status", "method", "path"},
	)

	if err := prometheus.Register(metr.Hits); err != nil {
		return nil, err
	}

	metr.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: name + "_times",
		},
		[]string{"status", "method", "path"},
	)

	if err := prometheus.Register(metr.Times); err != nil {
		return nil, err
	}

	if err := prometheus.Register(prometheus.NewBuildInfoCollector()); err != nil {
		return nil, err
	}

	go func() {
		router := gin.New()
		router.GET("/metrics", gin.WrapH(promhttp.Handler()))
		slog.Info("Metrics server is running on port: %s", address)
		if err := router.Run(address); err != nil {
			slog.Any("Metrics server error", err)
		}
	}()

	return &metr, nil
}
