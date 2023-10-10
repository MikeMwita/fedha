package middleware

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/pkg/metrics"
	"github.com/gin-gonic/gin"
	"time"
)

func (m middlewareManager) RequestMetrics(metrics metrics.MetricsHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		duration := time.Since(start).Seconds()

		status := c.Writer.Status()
		method := c.Request.Method
		path := c.FullPath()

		// Record response time and hits
		metrics.ObserveResponseTime(status, method, path, duration)
		metrics.IncHits(status, method, path)

	}
}
