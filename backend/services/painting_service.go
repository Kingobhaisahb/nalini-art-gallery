package services

import (
	"errors"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
)

type PaintingService struct {
	PaintingRepo repositories.PaintingRepository
	PaintingImageRepo repositories.PaintingImageRepository
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

func (s *PaintingService) GetAllPaintings(
	featured *bool,
	status *string,
) ([]models.Painting, error) {

	paintings, err := s.PaintingRepo.GetAllPaintings(
		featured,
		status,
	)

	if err != nil {
		return nil, err
	}

	return paintings, nil
}

func (s *PaintingService) GetPaintingByID(
	id uint,
) (*models.Painting, []models.PaintingImage, error) {

	painting, err := s.PaintingRepo.GetPaintingByID(id)

	if err != nil {
		return nil, nil, err
	}

	images, err := s.PaintingImageRepo.GetImagesByPaintingID(id)

	if err != nil {
		return nil, nil, err
	}

	return painting, images, nil
}

func (s *PaintingService) UpdatePainting(
	id uint,
	req dto.UpdatePaintingRequest,
) (*models.Painting, error) {

	painting, err := s.PaintingRepo.GetPaintingByID(id)

	if err != nil {
		return nil, err
	}

	status := req.Status

	if status == "" {
		status = painting.Status
	}

	if status != "AVAILABLE" && status != "SOLD" {
		return nil, errors.New("invalid painting status")
	}

	painting.Title = req.Title
	painting.Price = req.Price
	painting.Description = req.Description
	painting.Category = req.Category
	painting.Medium = req.Medium
	painting.Width = req.Width
	painting.Height = req.Height
	painting.Unit = req.Unit
	painting.Featured = req.Featured
	painting.Status = status
	painting.Tags = req.Tags
	painting.ProcessVideoURL = req.ProcessVideoURL

	err = s.PaintingRepo.UpdatePainting(painting)

	if err != nil {
		return nil, err
	}

	return painting, nil
}

func (s *PaintingService) DeletePainting(id uint) error {

	err := s.PaintingRepo.DeletePainting(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *PaintingService) UpdateFeatured(id uint, featured bool) error {

	return s.PaintingRepo.UpdateFeatured(id, featured)
}

func (s *PaintingService) UpdateStatus(id uint, status string) error {

	if status != "AVAILABLE" && status != "SOLD" {
		return errors.New("invalid painting status")
	}

	return s.PaintingRepo.UpdateStatus(id, status)
}

func (s *PaintingService) SearchPaintings(query string) ([]models.Painting, error) {

	if query == "" {
		return nil, errors.New("search query is required")
	}

	return s.PaintingRepo.SearchPaintings(query)
}
