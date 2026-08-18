package controllers

import (
	"net/http"
	"strconv"

	"github.com/Kingobhaisahb/nalini-art-gallery/services"
	"github.com/gin-gonic/gin"
)

type PaintingVideoController struct {
	PaintingVideoService services.PaintingVideoService
}

func (p *PaintingVideoController) UploadVideo(c *gin.Context) {

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

	// Get uploaded video
	fileHeader, err := c.FormFile("video")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Video file is required",
		})
		return
	}

	// Open video
	file, err := fileHeader.Open()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to open video",
			"error":   err.Error(),
		})
		return
	}

	defer file.Close()

	// Upload video
	videoURL, err := p.PaintingVideoService.UploadVideo(
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
			"message": "Failed to upload video",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Video uploaded successfully",
		"video_url": videoURL,
	})
}