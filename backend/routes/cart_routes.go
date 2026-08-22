package routes

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/controllers"
	"github.com/Kingobhaisahb/nalini-art-gallery/middleware"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.Engine) {

	cartRepo := repositories.CartRepository{}
	paintingRepo := repositories.PaintingRepository{}

	cartService := services.CartService{
		CartRepo:     cartRepo,
		PaintingRepo: paintingRepo,
	}

	cartController := controllers.CartController{
		CartService: cartService,
	}

	cart := router.Group("/api/cart")

	cart.Use(
		middleware.JWTMiddleware(),
	)

	{
		cart.POST(
			"",
			cartController.AddToCart,
		)

		cart.GET(
			"",
			cartController.GetCart,
		)

		cart.DELETE(
			"/:painting_id",
			cartController.RemoveFromCart,
		)

		cart.DELETE(
			"",
			cartController.ClearCart,
		)
	}
}