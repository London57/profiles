package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const requestIDHeader = "X-Request-ID"

func RequestID() gin.HandlerFunc {

	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(requestIDHeader)
		if requestID == "" {
			requestID = uuid.NewString()
		}
		c.Request.Header.Set(requestIDHeader, requestID)
		c.Writer.Header().Set(requestIDHeader, requestID)

		c.Next()
	}
}