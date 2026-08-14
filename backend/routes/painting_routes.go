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

	// Public Painting Routes
	router.GET(
		"/api/paintings",
		paintingController.GetAllPaintings,
	)

	router.GET(
		"/api/paintings/:id",
		paintingController.GetPaintingByID,
	)

	// Admin Painting Routes
	adminPaintings := router.Group("/api/paintings")

	adminPaintings.Use(
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
	)

	{
		adminPaintings.POST(
			"",
			paintingController.CreatePainting,
		)

		adminPaintings.PUT(
			"/:id",
			paintingController.UpdatePainting,
		)

		adminPaintings.DELETE(
			"/:id",
			paintingController.DeletePainting,
		)

		adminPaintings.PATCH(
			"/:id/featured",
			paintingController.UpdateFeatured,
		)

		adminPaintings.PATCH(
			"/:id/status",
			paintingController.UpdateStatus,
		)
	}
}