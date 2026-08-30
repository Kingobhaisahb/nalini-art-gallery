package repositories

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
)

type PaymentRepository struct{}

func (r *PaymentRepository) CreatePayment(
	payment *models.Payment,
) error {
	return database.DB.Create(payment).Error
}

func (r *PaymentRepository) GetPaymentByOrderID(
	orderID uint,
) (*models.Payment, error) {

	var payment models.Payment

	err := database.DB.
		Where("order_id = ?", orderID).
		First(&payment).Error

	if err != nil {
		return nil, err
	}

	return &payment, nil
}