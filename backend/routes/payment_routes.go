package routes

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/controllers"
	"github.com/Kingobhaisahb/nalini-art-gallery/middleware"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(router *gin.Engine) {

	paymentRepo := repositories.PaymentRepository{}

	razorpayService := services.NewRazorpayService()

	paymentService := services.PaymentService{
		PaymentRepo:     paymentRepo,
		RazorpayService: razorpayService,
	}

	paymentController := controllers.PaymentController{
		PaymentService: &paymentService,
	}

	paymentRoutes := router.Group("/api/payments")

	paymentRoutes.Use(
		middleware.JWTMiddleware(),
	)

	{
		paymentRoutes.POST(
			"/verify",
			paymentController.VerifyPayment,
		)
	}
}