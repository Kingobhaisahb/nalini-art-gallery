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

func (r *PaintingImageRepository) GetImagesByPaintingID(
	paintingID uint,
) ([]models.PaintingImage, error) {

	var images []models.PaintingImage

	err := database.DB.
		Where("painting_id = ?", paintingID).
		Order("created_at ASC").
		Find(&images).Error

	if err != nil {
		return nil, err
	}

	return images, nil
}