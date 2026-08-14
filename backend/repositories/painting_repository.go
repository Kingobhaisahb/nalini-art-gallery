package repositories

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"

	"gorm.io/gorm"
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

func (r *PaintingRepository) DeletePainting(id uint) error {
	return database.DB.Delete(&models.Painting{}, id).Error
}

func (r *PaintingRepository) UpdateFeatured(id uint, featured bool) error {

	result := database.DB.
		Model(&models.Painting{}).
		Where("id = ?", id).
		Update("featured", featured)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PaintingRepository) UpdateStatus(id uint, status string) error {

	result := database.DB.
		Model(&models.Painting{}).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}