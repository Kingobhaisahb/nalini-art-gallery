package routes

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/controllers"
	"github.com/Kingobhaisahb/nalini-art-gallery/middleware"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

func DashboardRoutes(router *gin.Engine) {

	dashboardRepo := repositories.DashboardRepository{}

	dashboardService := services.DashboardService{
		DashboardRepo: dashboardRepo,
	}

	dashboardController := controllers.DashboardController{
		DashboardService: dashboardService,
	}

	adminDashboard := router.Group("/api/admin/dashboard")

	adminDashboard.Use(
		middleware.JWTMiddleware(),
		middleware.AdminMiddleware(),
	)

	{
		adminDashboard.GET(
			"",
			dashboardController.GetAdminDashboard,
		)
	}
}