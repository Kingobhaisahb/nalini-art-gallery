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