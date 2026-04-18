package middlewares

import (
	"time"

	"github.com/London57/profiles/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.FromContext(c)
		
		before := time.Now()
		log.Debug(
			"incoming HTTP request",
			zap.Time("time", before.UTC()),
		)
		c.Next()
		log.Debug(
			"done request",
			zap.Int("status_code", c.Writer.Status()),
			zap.Duration("latency", time.Since(before)),
		)
	}
}