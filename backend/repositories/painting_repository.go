package repositories

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"

	"gorm.io/gorm"
	"errors"
)

type PaintingRepository struct{}

func (r *PaintingRepository) CreatePainting(painting *models.Painting) error {
	return database.DB.Create(painting).Error
}

func (r *PaintingRepository) GetAllPaintings(
	featured *bool,
	status *string,
) ([]models.Painting, error) {

	var paintings []models.Painting

	query := database.DB

	if featured != nil {
		query = query.Where("featured = ?", *featured)
	}

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	err := query.
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

func (r *PaintingRepository) SearchPaintings(query string) ([]models.Painting, error) {

	var paintings []models.Painting

	search := "%" + query + "%"

	err := database.DB.
		Where(
			"title LIKE ? OR description LIKE ? OR category LIKE ? OR medium LIKE ? OR tags LIKE ?",
			search,
			search,
			search,
			search,
			search,
		).
		Order("created_at DESC").
		Find(&paintings).Error

	if err != nil {
		return nil, err
	}

	return paintings, nil
}

func (r *PaintingRepository) GetPaintingsByPrice(
	order string,
) ([]models.Painting, error) {

	var paintings []models.Painting

	if order != "asc" && order != "desc" {
		return nil, errors.New("invalid price sort order")
	}

	err := database.DB.
		Order("price " + order).
		Find(&paintings).Error

	if err != nil {
		return nil, err
	}

	return paintings, nil
}

func (r *PaintingRepository) GetNewestPaintings() ([]models.Painting, error) {

	var paintings []models.Painting

	err := database.DB.
		Order("created_at DESC").
		Find(&paintings).Error

	if err != nil {
		return nil, err
	}

	return paintings, nil
}

func (r *PaintingRepository) IncrementViews(id uint) error {

	return database.DB.
		Model(&models.Painting{}).
		Where("id = ?", id).
		UpdateColumn("views", gorm.Expr("views + ?", 1)).
		Error
}

func (r *PaintingRepository) GetAllPaintingsForAdmin() ([]models.Painting, error) {

	var paintings []models.Painting

	err := database.DB.
		Unscoped().
		Order("created_at DESC").
		Find(&paintings).Error

	if err != nil {
		return nil, err
	}

	return paintings, nil
}