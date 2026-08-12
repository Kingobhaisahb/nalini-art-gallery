package services

import (
	"errors"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
)

type PaintingService struct {
	PaintingRepo repositories.PaintingRepository
}

func (s *PaintingService) CreatePainting(req dto.CreatePaintingRequest) (*models.Painting, error) {

	status := req.Status

	if status == "" {
		status = "AVAILABLE"
	}

	if status != "AVAILABLE" && status != "SOLD" {
		return nil, errors.New("invalid painting status")
	}

	painting := models.Painting{
		Title:           req.Title,
		Price:           req.Price,
		Description:     req.Description,
		Category:        req.Category,
		Medium:          req.Medium,
		Width:           req.Width,
		Height:          req.Height,
		Unit:            req.Unit,
		Featured:        req.Featured,
		Status:          status,
		Tags:            req.Tags,
		ProcessVideoURL: req.ProcessVideoURL,
		Views:           0,
	}

	err := s.PaintingRepo.CreatePainting(&painting)

	if err != nil {
		return nil, err
	}

	return &painting, nil
}
