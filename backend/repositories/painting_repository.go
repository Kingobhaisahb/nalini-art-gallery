package repositories

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
)

type PaintingRepository struct{}

func (r *PaintingRepository) CreatePainting(painting *models.Painting) error {
	return database.DB.Create(painting).Error
}