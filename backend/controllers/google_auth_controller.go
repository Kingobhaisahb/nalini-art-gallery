package controllers

import (
	"net/http"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

type GoogleAuthController struct {
	GoogleAuthService services.GoogleAuthService
}

func (g *GoogleAuthController) GoogleLogin(c *gin.Context) {

	var req dto.GoogleLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	response, err := g.GoogleAuthService.GoogleLogin(req)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response)
}