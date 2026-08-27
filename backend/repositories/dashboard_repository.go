package repositories

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
)

type DashboardRepository struct{}

func (r *DashboardRepository) GetTotalPaintings() (int64, error) {

	var count int64

	err := database.DB.
		Model(&models.Painting{}).
		Count(&count).Error

	return count, err
}

func (r *DashboardRepository) GetAvailablePaintings() (int64, error) {

	var count int64

	err := database.DB.
		Model(&models.Painting{}).
		Where("status = ?", "AVAILABLE").
		Count(&count).Error

	return count, err
}

func (r *DashboardRepository) GetSoldPaintings() (int64, error) {

	var count int64

	err := database.DB.
		Model(&models.Painting{}).
		Where("status = ?", "SOLD").
		Count(&count).Error

	return count, err
}

func (r *DashboardRepository) GetFeaturedPaintings() (int64, error) {

	var count int64

	err := database.DB.
		Model(&models.Painting{}).
		Where("featured = ?", true).
		Count(&count).Error

	return count, err
}

func (r *DashboardRepository) GetTotalOrders() (int64, error) {

	var count int64

	err := database.DB.
		Model(&models.Order{}).
		Count(&count).Error

	return count, err
}

func (r *DashboardRepository) GetOrdersByStatus(status string) (int64, error) {

	var count int64

	err := database.DB.
		Model(&models.Order{}).
		Where("status = ?", status).
		Count(&count).Error

	return count, err
}

func (r *DashboardRepository) GetTotalCustomers() (int64, error) {

	var count int64

	err := database.DB.
		Model(&models.User{}).
		Where("role = ?", "CUSTOMER").
		Count(&count).Error

	return count, err
}

func (r *DashboardRepository) GetTotalRevenue() (float64, error) {

	var revenue float64

	err := database.DB.
		Model(&models.Order{}).
		Where("status != ?", "CANCELLED").
		Select("COALESCE(SUM(total_price), 0)").
		Scan(&revenue).Error

	return revenue, err
}