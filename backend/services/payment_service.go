package services

import (
	"errors"
	"fmt"
	"os"

	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	razorpayutils "github.com/razorpay/razorpay-go/utils"
	"gorm.io/gorm"
)

type PaymentService struct {
	PaymentRepo     repositories.PaymentRepository
	RazorpayService *RazorpayService
}

func (s *PaymentService) CreatePayment(
	order *models.Order,
) (*models.Payment, error) {

	if order == nil {
		return nil, errors.New("order is required")
	}

	razorpayOrderID, err := s.RazorpayService.CreateOrder(
		order.TotalPrice,
		"INR",
		fmt.Sprintf("order_%d", order.ID),
	)

	if err != nil {
		return nil, err
	}

	payment := &models.Payment{
		OrderID:         order.ID,
		RazorpayOrderID: razorpayOrderID,
		Amount:          order.TotalPrice,
		Currency:        "INR",
		Status:          "CREATED",
	}

	if err := s.PaymentRepo.CreatePayment(payment); err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *PaymentService) VerifyPayment(
	userID uint,
	orderID uint,
	razorpayPaymentID string,
	razorpaySignature string,
) error {

	if razorpayPaymentID == "" {
		return errors.New("razorpay payment ID is required")
	}

	if razorpaySignature == "" {
		return errors.New("razorpay signature is required")
	}

	// Make sure this order belongs to the logged-in user.
	var order models.Order

	err := database.DB.
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("order not found")
		}

		return err
	}

	// Find our payment record.
	payment, err := s.PaymentRepo.GetPaymentByOrderID(orderID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("payment record not found")
		}

		return err
	}

	// If payment was already verified, make sure the paintings
	// are also marked as SOLD.
	if payment.Status == "SUCCESS" {
		return database.DB.Transaction(func(tx *gorm.DB) error {

			var orderItems []models.OrderItem

			if err := tx.
				Where("order_id = ?", orderID).
				Find(&orderItems).Error; err != nil {
				return err
			}

			for _, item := range orderItems {

				if err := tx.
					Model(&models.Painting{}).
					Where("id = ?", item.PaintingID).
					Update("status", "SOLD").Error; err != nil {
					return err
				}
			}

			return nil
		})
	}

	// Make sure the Razorpay order ID exists.
	if payment.RazorpayOrderID == "" {
		return errors.New("razorpay order ID missing")
	}

	// Razorpay secret from environment.
	secret := os.Getenv("RAZORPAY_KEY_SECRET")

	if secret == "" {
		return errors.New("razorpay secret is not configured")
	}

	// Verify Razorpay signature.
	params := map[string]interface{}{
		"razorpay_order_id":   payment.RazorpayOrderID,
		"razorpay_payment_id": razorpayPaymentID,
	}

	isValid := razorpayutils.VerifyPaymentSignature(
		params,
		razorpaySignature,
		secret,
	)

	if !isValid {
		return errors.New("invalid razorpay signature")
	}

	// Update payment, order, and paintings together.
	return database.DB.Transaction(func(tx *gorm.DB) error {

		// 1. Mark payment successful.
		if err := tx.
			Model(&models.Payment{}).
			Where("id = ?", payment.ID).
			Updates(map[string]interface{}{
				"razorpay_payment_id": razorpayPaymentID,
				"razorpay_signature":  razorpaySignature,
				"status":              "SUCCESS",
			}).Error; err != nil {
			return err
		}

		// 2. Payment successful → confirm order.
		if err := tx.
			Model(&models.Order{}).
			Where("id = ?", orderID).
			Update("status", "CONFIRMED").Error; err != nil {
			return err
		}

		// 3. Get all paintings belonging to this order.
		var orderItems []models.OrderItem

		if err := tx.
			Where("order_id = ?", orderID).
			Find(&orderItems).Error; err != nil {
			return err
		}

		// 4. Mark every purchased painting as SOLD.
		for _, item := range orderItems {

			if err := tx.
				Model(&models.Painting{}).
				Where("id = ?", item.PaintingID).
				Update("status", "SOLD").Error; err != nil {
				return err
			}
		}

		return nil
	})
}
