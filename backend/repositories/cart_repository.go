package repositories

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
)

type CartRepository struct{}

func (r *CartRepository) AddToCart(cart *models.Cart) error {
	return database.DB.Create(cart).Error
}

func (r *CartRepository) GetCartByUserID(userID uint) ([]models.Cart, error) {

	var cart []models.Cart

	err := database.DB.
		Preload("Painting").
		Where("user_id = ?", userID).
		Order("created_at ASC").
		Find(&cart).Error

	return cart, err
}

func (r *CartRepository) GetCartItem(
	userID uint,
	paintingID uint,
) (*models.Cart, error) {

	var cart models.Cart

	err := database.DB.
		Preload("Painting").
		Where("user_id = ? AND painting_id = ?", userID, paintingID).
		First(&cart).Error

	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *CartRepository) RemoveFromCart(
	userID uint,
	paintingID uint,
) error {

	return database.DB.
	Where("user_id = ? AND painting_id = ?", userID, paintingID).
	Delete(&models.Cart{}).Error
}

func (r *CartRepository) ClearCart(userID uint) error {

	return database.DB.
	Where("user_id = ?", userID).
	Delete(&models.Cart{}).Error
}