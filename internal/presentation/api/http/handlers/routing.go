package handlers

import (
	"github.com/London57/jwt-auth/middleware"
	create "github.com/London57/profiles/internal/presentation/api/http/handlers/profileCreate"
	update "github.com/London57/profiles/internal/presentation/api/http/handlers/profileUpdate"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine, createHand create.ProfileCreateHandler, update update.ProfileUpdateHandler, secret string) {
	profiles := app.Group("profiles")

	auth_profiles := app.Group("auth-profiles")

	jwt_middleware := middleware.JwtAuthMiddleware(secret)

	auth_profiles.Use(jwt_middleware)
	
	profiles.POST("/create", createHand.CreateProfile)
	auth_profiles.PATCH("/update", update.UpdateProfile)
}