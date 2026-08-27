package routes

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/controllers"
	"github.com/Kingobhaisahb/nalini-art-gallery/middleware"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	userRepo := repositories.UserRepository{}
	orderRepo := repositories.OrderRepository{}

	userService := services.UserService{
		UserRepo:  userRepo,
		OrderRepo: orderRepo,
	}

	userController := controllers.UserController{
		UserService: userService,
	}

	adminUsers := router.Group("/api/admin/users")

	adminUsers.Use(
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
	)

	{
		// View all users
		adminUsers.GET(
			"",
			userController.GetAllUsers,
		)

		// View specific user + orders
		adminUsers.GET(
			"/:id",
			userController.GetUserByID,
		)

		// Change user role
		adminUsers.PATCH(
			"/:id/role",
			userController.UpdateUserRole,
		)
	}
}