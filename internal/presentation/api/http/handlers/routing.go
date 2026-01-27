package handlers

import (
	create "github.com/London57/profiles/internal/presentation/api/http/handlers/profileCreaate"
	update "github.com/London57/profiles/internal/presentation/api/http/handlers/profileUpdate"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine, createHand create.ProfileCreateHandler, update update.ProfileUpdateHandler) {
	profiles := app.Group("profiles")

	profiles.POST("/registration", createHand.CreateProfile)
	profiles.PATCH("/patch", update.UpdateProfile)
}