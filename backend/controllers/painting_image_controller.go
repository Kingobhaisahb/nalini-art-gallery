package controllers

import (
	"net/http"
	"strconv"

	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

type PaintingImageController struct {
	PaintingImageService services.PaintingImageService
}

func (p *PaintingImageController) UploadImage(c *gin.Context) {

	// Get painting ID
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid painting ID",
		})
		return
	}

	// Get uploaded file
	fileHeader, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Image file is required",
		})
		return
	}

	// Open file
	file, err := fileHeader.Open()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to upload image",
			"error":   err.Error(),
		})
		return
	}

	defer file.Close()

	// Upload image
	image, err := p.PaintingImageService.UploadImage(
		uint(id),
		file,
	)

	if err != nil {

		if err.Error() == "painting not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Painting not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to upload image",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Image uploaded successfully",
		"image": image,
	})
}