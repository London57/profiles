package handlers

import "github.com/gin-gonic/gin"

func InitRouter(app *gin.Engine, profilesHand ProfilesHandler) {
	profiles := app.Group("profiles")

	profiles.POST("/registration", profilesHand.CreateProfile)
}