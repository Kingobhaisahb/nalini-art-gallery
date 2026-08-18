package routes

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/controllers"
	"github.com/Kingobhaisahb/nalini-art-gallery/middleware"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

func PaintingVideoRoutes(router *gin.Engine) {

	// Repository
	paintingRepo := repositories.PaintingRepository{}

	// Service
	paintingVideoService := services.PaintingVideoService{
		PaintingRepo: paintingRepo,
	}

	// Controller
	paintingVideoController := controllers.PaintingVideoController{
		PaintingVideoService: paintingVideoService,
	}

	// Admin Video Routes
	adminPaintingVideos := router.Group("/api/paintings")

	adminPaintingVideos.Use(
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
	)

	{
		adminPaintingVideos.POST(
			"/:id/video",
			paintingVideoController.UploadVideo,
		)
	}
}