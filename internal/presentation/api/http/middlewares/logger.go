package middlewares

import (
	"github.com/London57/profiles/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger(log logger.Logger) gin.HandlerFunc {
	const requestIDHeader = "X-Request-ID"

	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(requestIDHeader)

		l := log.With(
			zap.String("request_id", requestID),
			zap.String("url", c.Request.URL.String()),
		)

		c.Set("log", l)
		c.Next()
	}
}