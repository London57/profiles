package middlewares

import (
	"fmt"
	"net/http"

	"github.com/London57/profiles/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Panic() gin.HandlerFunc {
	const requestIDHeader = "X-Request-ID"

	return func(c *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				logger := logger.FromContext(c)
				err := fmt.Errorf("unexpected panic: %v", p)
				logger.Error("HTTP request got unexpected panic", zap.Error(err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": err,
				})
			}
		}()
		c.Next()
	}
}