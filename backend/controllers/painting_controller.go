package controllers

import (
	"net/http"
	"strconv"
	"errors"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/services"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaintingController struct {
	PaintingService services.PaintingService
	PaintingImageRepo   repositories.PaintingImageRepository
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

	paintings, err := p.PaintingService.GetAllPaintings(nil)

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

func (p *PaintingController) GetPaintingByID(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid painting ID",
		})
		return
	}

	painting, images, err := p.PaintingService.GetPaintingByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Painting not found",
		})
		return
	}

	imageResponses := make([]dto.PaintingImageResponse, 0, len(images))

	for _, image := range images {
		imageResponses = append(imageResponses, dto.PaintingImageResponse{
			ID:        image.ID,
			ImageURL:  image.ImageURL,
			CreatedAt: image.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"painting": dto.PaintingResponse{
			ID:              painting.ID,
			Title:           painting.Title,
			Price:           painting.Price,
			Description:     painting.Description,
			Category:        painting.Category,
			Medium:          painting.Medium,
			Width:           painting.Width,
			Height:          painting.Height,
			Unit:            painting.Unit,
			Featured:        painting.Featured,
			Status:          painting.Status,
			Tags:            painting.Tags,
			ProcessVideoURL:  painting.ProcessVideoURL,
			Views:           painting.Views,
			Images: imageResponses,
			CreatedAt:       painting.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:       painting.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func (p *PaintingController) UpdatePainting(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid painting ID",
		})
		return
	}

	var req dto.UpdatePaintingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	painting, err := p.PaintingService.UpdatePainting(
		uint(id),
		req,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Painting not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Painting updated successfully",
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
			Status:          painting.Status,
			Tags:             painting.Tags,
			ProcessVideoURL: painting.ProcessVideoURL,
			Views:            painting.Views,
			CreatedAt:        painting.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        painting.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func (p *PaintingController) DeletePainting(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid painting ID",
		})
		return
	}

	err = p.PaintingService.DeletePainting(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete painting",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Painting deleted successfully",
	})
}

func (p *PaintingController) UpdateFeatured(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid painting ID",
		})
		return
	}

	var req dto.UpdateFeaturedRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	err = p.PaintingService.UpdateFeatured(
		uint(id),
		req.Featured,
	)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Painting not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update featured status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "Featured status updated successfully",
		"featured": req.Featured,
	})
}

func (p *PaintingController) UpdateStatus(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid painting ID",
		})
		return
	}

	var req dto.UpdatePaintingStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	err = p.PaintingService.UpdateStatus(
		uint(id),
		req.Status,
	)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Painting not found",
			})
			return
		}

		if err.Error() == "invalid painting status" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update painting status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Painting status updated successfully",
		"status":  req.Status,
	})
}

func (p *PaintingController) GetFeaturedPaintings(c *gin.Context) {

	featured := true

	paintings, err := p.PaintingService.GetAllPaintings(&featured)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch featured paintings",
		})
		return
	}

	responses := make([]dto.PaintingResponse, 0, len(paintings))

	for _, painting := range paintings {

		responses = append(responses, dto.PaintingResponse{
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

