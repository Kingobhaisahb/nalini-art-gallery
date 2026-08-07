package routes

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/controllers"
	"github.com/Kingobhaisahb/nalini-art-gallery/middleware"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	// Repositories
	userRepo := repositories.UserRepository{}
	passwordResetRepo := repositories.PasswordResetRepository{}

	// Services
	authService := services.AuthService{
		UserRepo: userRepo,
	}

	passwordResetService := services.PasswordResetService{
		UserRepo:          userRepo,
		PasswordResetRepo: passwordResetRepo,
	}

	googleAuthService := services.GoogleAuthService{
		UserRepo: userRepo,
	}

	// Controllers
	authController := controllers.AuthController{
		AuthService: authService,
	}

	passwordResetController := controllers.PasswordResetController{
		PasswordResetService: passwordResetService,
	}

	googleAuthController := controllers.GoogleAuthController{
		GoogleAuthService: googleAuthService,
	}

	// Public Routes
	auth := router.Group("/api/auth")
	{
		auth.POST("/signup", authController.Signup)
		auth.POST("/login", authController.Login)

		auth.POST("/forgot-password", passwordResetController.ForgotPassword)
		auth.POST("/reset-password", passwordResetController.ResetPassword)

		auth.POST("/google", googleAuthController.GoogleLogin)
	}

	// Protected Routes
	protected := router.Group("/api")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.GET("/profile", authController.Profile)
	}
}