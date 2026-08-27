package controllers

import (
	"net/http"

	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	DashboardService services.DashboardService
}

func (d *DashboardController) GetAdminDashboard(c *gin.Context) {

	dashboard, err := d.DashboardService.GetAdminDashboard()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch dashboard data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"dashboard": dashboard,
	})
}