package repositories

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
)

type OrderRepository struct{}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return database.DB.Create(order).Error
}

func (r *OrderRepository) CreateOrderItem(item *models.OrderItem) error {
	return database.DB.Create(item).Error
}

func (r *OrderRepository) GetOrdersByUserID(
	userID uint,
) ([]models.Order, error) {

	var orders []models.Order

	err := database.DB.
		Preload("OrderItems").
		Preload("OrderItems.Painting").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error

	return orders, err
}

func (r *OrderRepository) GetOrderByID(
	orderID uint,
	userID uint,
) (*models.Order, error) {

	var order models.Order

	err := database.DB.
		Preload("OrderItems").
		Preload("OrderItems.Painting").
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) UpdateOrderStatus(
	orderID uint,
	status string,
) error {

	return database.DB.
		Model(&models.Order{}).
		Where("id = ?", orderID).
		Update("status", status).Error
}