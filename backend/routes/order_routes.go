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
	paymentRepo := repositories.PaymentRepository{}

	razorpayService := services.NewRazorpayService()

	paymentService := services.PaymentService{
		PaymentRepo:     paymentRepo,
		RazorpayService: razorpayService,
	}

	orderService := services.OrderService{
		OrderRepo:      orderRepo,
		CartRepo:       cartRepo,
		PaymentService: &paymentService,
	}

	orderController := controllers.OrderController{
		OrderService: orderService,
	}

	// =========================
	// CUSTOMER ORDER ROUTES
	// =========================

	orderRoutes := router.Group("/api/orders")

	orderRoutes.Use(
		middleware.JWTMiddleware(),
	)

	{
		orderRoutes.POST(
			"/checkout",
			orderController.Checkout,
		)

		orderRoutes.GET(
			"",
			orderController.GetMyOrders,
		)

		orderRoutes.PATCH(
			"/:id/cancel",
			orderController.CancelMyOrder,
		)

		orderRoutes.GET(
			"/:id",
			orderController.GetMyOrder,
		)
	}

	// =========================
	// ADMIN ORDER ROUTES
	// =========================

	adminOrderRoutes := router.Group("/api/admin/orders")

	adminOrderRoutes.Use(
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
	)

	{
		// View all orders
		adminOrderRoutes.GET(
			"",
			orderController.GetAllOrders,
		)

		// Cancel order
		adminOrderRoutes.PATCH(
			"/:id/cancel",
			orderController.CancelAdminOrder,
		)

		// View any order
		adminOrderRoutes.GET(
			"/:id",
			orderController.GetAdminOrder,
		)

		// Update order status
		adminOrderRoutes.PATCH(
			"/:id/status",
			orderController.UpdateOrderStatus,
		)
	}
}