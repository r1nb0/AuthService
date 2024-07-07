package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/r1nb0/UserService/pkg/metrics"
	"strconv"
	"time"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.FullPath()
		method := ctx.Request.Method
		ctx.Next()
		status := ctx.Writer.Status()
		metrics.HttpDuration.WithLabelValues(
			path, method, strconv.Itoa(status),
		).Observe(float64(time.Since(start)) / float64(time.Millisecond))
	}
}
