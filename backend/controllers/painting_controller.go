package controllers

import (
	"net/http"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"

	"github.com/gin-gonic/gin"
)

type PaintingController struct {
	PaintingService services.PaintingService
}

func (p *PaintingController) CreatePainting(c *gin.Context) {

	var req dto.CreatePaintingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	painting, err := p.PaintingService.CreatePainting(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Painting created successfully",
		"painting": dto.PaintingResponse{
			ID:              painting.ID,
			Title:           painting.Title,
			Price:           painting.Price,
			Description:     painting.Description,
			Category:        painting.Category,
			Medium:          painting.Medium,
			Width:           painting.Width,
			Height:          painting.Height,
			Unit:             painting.Unit,
			Featured:        painting.Featured,
			Status:           painting.Status,
			Tags:             painting.Tags,
			ProcessVideoURL:  painting.ProcessVideoURL,
			Views:            painting.Views,
			CreatedAt:        painting.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        painting.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func (p *PaintingController) GetAllPaintings(c *gin.Context) {

	paintings, err := p.PaintingService.GetAllPaintings()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch paintings",
		})
		return
	}

	responses := make([]dto.PaintingResponse, 0, len(paintings))

	for _, painting := range paintings {

		responses = append(responses, dto.PaintingResponse{
			ID:             painting.ID,
			Title:          painting.Title,
			Price:          painting.Price,
			Description:    painting.Description,
			Category:       painting.Category,
			Medium:         painting.Medium,
			Width:          painting.Width,
			Height:         painting.Height,
			Unit:            painting.Unit,
			Featured:       painting.Featured,
			Status:         painting.Status,
			Tags:            painting.Tags,
			ProcessVideoURL: painting.ProcessVideoURL,
			Views:           painting.Views,
			CreatedAt:       painting.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:       painting.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"paintings": responses,
	})
}