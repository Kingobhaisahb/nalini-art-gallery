package services

import (
	"errors"
	"gorm.io/gorm"

	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
)

type OrderService struct {
	OrderRepo repositories.OrderRepository
	CartRepo  repositories.CartRepository
}

func (s *OrderService) Checkout(
	userID uint,
	addressID uint,
) (*models.Order, error) {

	if addressID == 0 {
		return nil, errors.New("address is required")
	}

	var cartItems []models.Cart

	err := database.DB.
		Preload("Painting").
		Where("user_id = ?", userID).
		Find(&cartItems).Error

	if err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	var totalPrice float64

	for _, item := range cartItems {

		if item.Painting.ID == 0 {
			return nil, errors.New("painting not found")
		}

		if item.Painting.Status != "AVAILABLE" {
			return nil, errors.New(
				"painting is no longer available",
			)
		}

		totalPrice += item.Painting.Price
	}

	var orderID uint

	err = database.DB.Transaction(func(tx *gorm.DB) error {

		newOrder := models.Order{
			UserID:     userID,
			AddressID:  addressID,
			TotalPrice: totalPrice,
			Status:     "PENDING",
		}

		if err := tx.Create(&newOrder).Error; err != nil {
			return err
		}

		for _, item := range cartItems {

			orderItem := models.OrderItem{
				OrderID:    newOrder.ID,
				PaintingID: item.PaintingID,
				Price:      item.Painting.Price,
			}

			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}
		}

		if err := tx.
			Where("user_id = ?", userID).
			Delete(&models.Cart{}).Error; err != nil {
			return err
		}

		orderID = newOrder.ID

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Reload order with OrderItems and Painting
	order, err := s.OrderRepo.GetOrderByID(orderID, userID)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetUserOrders(
	userID uint,
) ([]models.Order, error) {

	return s.OrderRepo.GetOrdersByUserID(userID)
}

func (s *OrderService) GetUserOrder(
	orderID uint,
	userID uint,
) (*models.Order, error) {

	return s.OrderRepo.GetOrderByID(orderID, userID)
}

func (s *OrderService) UpdateOrderStatus(
	orderID uint,
	status string,
) error {

	validStatuses := map[string]bool{
		"PENDING":   true,
		"CONFIRMED": true,
		"SHIPPED":   true,
		"DELIVERED": true,
		"CANCELLED": true,
	}

	if !validStatuses[status] {
		return errors.New("invalid order status")
	}

	return s.OrderRepo.UpdateOrderStatus(orderID, status)
}