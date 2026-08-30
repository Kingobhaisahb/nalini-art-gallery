package services

import (
    "errors"

    "gorm.io/gorm"
    "gorm.io/gorm/clause"

    "github.com/Kingobhaisahb/nalini-art-gallery/database"
    "github.com/Kingobhaisahb/nalini-art-gallery/models"
    "github.com/Kingobhaisahb/nalini-art-gallery/repositories"
)

type OrderService struct {
	OrderRepo     repositories.OrderRepository
	CartRepo      repositories.CartRepository
	PaymentService *PaymentService
}

func (s *OrderService) Checkout(
	userID uint,
	addressID uint,
) (*models.Order, *models.Payment, error) {

	if addressID == 0 {
		return nil, nil, errors.New("address is required")
	}

	var cartItems []models.Cart

	err := database.DB.
		Preload("Painting").
		Where("user_id = ?", userID).
		Find(&cartItems).Error

	if err != nil {
		return nil, nil, err
	}

	if len(cartItems) == 0 {
		return nil, nil, errors.New("cart is empty")
	}

	var totalPrice float64

	// Check that all paintings are available.
	for _, item := range cartItems {

		if item.Painting.ID == 0 {
			return nil, nil, errors.New("painting not found")
		}

		if item.Painting.Status != "AVAILABLE" {
			return nil, nil, errors.New(
				"painting is no longer available",
			)
		}

		totalPrice += item.Painting.Price
	}

	var orderID uint

	err = database.DB.Transaction(func(tx *gorm.DB) error {

		// Re-check and lock each painting before creating the order.
		for _, item := range cartItems {

			var painting models.Painting

			err := tx.
				Clauses(
					clause.Locking{
						Strength: "UPDATE",
					},
				).
				First(&painting, item.PaintingID).Error

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return errors.New("painting not found")
				}

				return err
			}

			if painting.Status != "AVAILABLE" {
				return errors.New(
					"painting is no longer available",
				)
			}
		}

		// Create order.
		newOrder := models.Order{
			UserID:     userID,
			AddressID:  addressID,
			TotalPrice: totalPrice,
			Status:     "PENDING",
		}

		if err := tx.Create(&newOrder).Error; err != nil {
			return err
		}

		// Create order items.
		// DO NOT mark paintings SOLD yet.
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

		// Clear user's cart.
		if err := tx.
			Where("user_id = ?", userID).
			Delete(&models.Cart{}).Error; err != nil {
			return err
		}

		orderID = newOrder.ID

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	// Reload order with OrderItems and Painting.
	order, err := s.OrderRepo.GetOrderByID(orderID, userID)

	if err != nil {
		return nil, nil, err
	}

	// Create Razorpay payment order and Payment DB record.
	payment, err := s.PaymentService.CreatePayment(order)

	if err != nil {
		return nil, nil, err
	}

	return order, payment, nil
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

func (s *OrderService) GetAllOrders() ([]models.Order, error) {

	return s.OrderRepo.GetAllOrders()
}

func (s *OrderService) GetAdminOrder(
	orderID uint,
) (*models.Order, error) {

	return s.OrderRepo.GetOrderByIDAdmin(orderID)
}

func (s *OrderService) CancelUserOrder(
	orderID uint,
	userID uint,
) error {

	return database.DB.Transaction(func(tx *gorm.DB) error {

		var order models.Order

		err := tx.
			Where("id = ? AND user_id = ?", orderID, userID).
			First(&order).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("order not found")
			}

			return err
		}

		// Customer can cancel only PENDING orders.
		if order.Status != "PENDING" {
			return errors.New(
				"order cannot be cancelled after confirmation",
			)
		}

		var items []models.OrderItem

		if err := tx.
			Where("order_id = ?", order.ID).
			Find(&items).Error; err != nil {
			return err
		}

		// Restore paintings to AVAILABLE.
		for _, item := range items {

			if err := tx.
				Model(&models.Painting{}).
				Where("id = ?", item.PaintingID).
				Update("status", "AVAILABLE").Error; err != nil {
				return err
			}
		}

		// Cancel order.
		if err := tx.
			Model(&order).
			Update("status", "CANCELLED").Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *OrderService) CancelAdminOrder(
	orderID uint,
) error {

	return database.DB.Transaction(func(tx *gorm.DB) error {

		var order models.Order

		err := tx.
			Where("id = ?", orderID).
			First(&order).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("order not found")
			}

			return err
		}

		// Already cancelled.
		if order.Status == "CANCELLED" {
			return errors.New("order is already cancelled")
		}

		var items []models.OrderItem

		if err := tx.
			Where("order_id = ?", order.ID).
			Find(&items).Error; err != nil {
			return err
		}

		// Restore all paintings belonging to this order.
		for _, item := range items {

			if err := tx.
				Model(&models.Painting{}).
				Where("id = ?", item.PaintingID).
				Update("status", "AVAILABLE").Error; err != nil {
				return err
			}
		}

		// Cancel order.
		if err := tx.
			Model(&order).
			Update("status", "CANCELLED").Error; err != nil {
			return err
		}

		return nil
	})
}