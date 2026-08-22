package routes

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/controllers"
	"github.com/Kingobhaisahb/nalini-art-gallery/middleware"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {

	orderRepo := repositories.OrderRepository{}
	cartRepo := repositories.CartRepository{}

	orderService := services.OrderService{
		OrderRepo: orderRepo,
		CartRepo:  cartRepo,
	}

	orderController := controllers.OrderController{
		OrderService: orderService,
	}

	orderRoutes := router.Group("/api/orders")

	orderRoutes.Use(
		middleware.JWTMiddleware(),
	)

	{
		// Checkout / create order
		orderRoutes.POST(
			"/checkout",
			orderController.Checkout,
		)

		// Customer order history
		orderRoutes.GET(
			"",
			orderController.GetMyOrders,
		)

		// Customer single order
		orderRoutes.GET(
			"/:id",
			orderController.GetMyOrder,
		)
	}
}