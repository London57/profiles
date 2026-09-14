package middlewares

import "github.com/gin-gonic/gin"

func UserActivity() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Get("userID")

		// отправка сообщения с user_id в брокер
		c.Next()
		
	}
}