package middleware

import (
	"github.com/Stefan923/go-estate-market/metrics"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func NewPrometheusMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		path := context.FullPath()
		method := context.Request.Method

		context.Next()
		status := context.Writer.Status()
		metrics.HttpDurationHistogram.WithLabelValues(path, method, strconv.Itoa(status)).
			Observe(float64(time.Since(start) / time.Millisecond))
	}
}
