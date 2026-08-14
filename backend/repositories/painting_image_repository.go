package repositories

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
)

type PaintingImageRepository struct{}

func (r *PaintingImageRepository) CreateImage(
	image *models.PaintingImage,
) error {

	return database.DB.Create(image).Error
}