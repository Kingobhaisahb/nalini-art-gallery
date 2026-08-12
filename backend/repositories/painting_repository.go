package repositories

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
)

type PaintingRepository struct{}

func (r *PaintingRepository) CreatePainting(painting *models.Painting) error {
	return database.DB.Create(painting).Error
}

func (r *PaintingRepository) GetAllPaintings() ([]models.Painting, error) {

	var paintings []models.Painting

	err := database.DB.
		Order("created_at DESC").
		Find(&paintings).Error

	if err != nil {
		return nil, err
	}

	return paintings, nil
}

func (r *PaintingRepository) GetPaintingByID(id uint) (*models.Painting, error) {

	var painting models.Painting

	err := database.DB.
		First(&painting, id).Error

	if err != nil {
		return nil, err
	}

	return &painting, nil
}

func (r *PaintingRepository) UpdatePainting(painting *models.Painting) error {

	return database.DB.Save(painting).Error
}