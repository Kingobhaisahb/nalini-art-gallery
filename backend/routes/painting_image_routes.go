package routes

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/controllers"
	"github.com/Kingobhaisahb/nalini-art-gallery/middleware"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

func PaintingImageRoutes(router *gin.Engine) {

	// Repository
	paintingRepo := repositories.PaintingRepository{}
	paintingImageRepo := repositories.PaintingImageRepository{}

	// Service
	paintingImageService := services.PaintingImageService{
		PaintingRepo:      paintingRepo,
		PaintingImageRepo: paintingImageRepo,
	}

	// Controller
	paintingImageController := controllers.PaintingImageController{
		PaintingImageService: paintingImageService,
	}

	// Admin Image Routes
	adminPaintingImages := router.Group("/api/paintings")

	adminPaintingImages.Use(
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
	)

	{
		adminPaintingImages.POST(
			"/:id/images",
			paintingImageController.UploadImage,
		)
	}
}