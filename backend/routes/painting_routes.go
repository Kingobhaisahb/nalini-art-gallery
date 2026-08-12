package routes

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/controllers"
	"github.com/Kingobhaisahb/nalini-art-gallery/middleware"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

func PaintingRoutes(router *gin.Engine) {

	// Repository
	paintingRepo := repositories.PaintingRepository{}

	// Service
	paintingService := services.PaintingService{
		PaintingRepo: paintingRepo,
	}

	// Controller
	paintingController := controllers.PaintingController{
		PaintingService: paintingService,
	}

	// Protected Painting Routes
	paintings := router.Group("/api/paintings")

	paintings.Use(
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
	)

	{
		paintings.POST("", paintingController.CreatePainting)
	}
}