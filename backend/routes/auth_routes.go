package routes

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/controllers"
	"github.com/Kingobhaisahb/nalini-art-gallery/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	authController := controllers.AuthController{}

	auth := router.Group("/api/auth")
	{
		auth.POST("/signup", authController.Signup)
		auth.POST("/login", authController.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.GET("/profile", authController.Profile)
	}
}