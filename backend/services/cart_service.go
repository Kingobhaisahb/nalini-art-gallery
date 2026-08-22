package services

import (
	"errors"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"

	"gorm.io/gorm"
)

type CartService struct {
	CartRepo     repositories.CartRepository
	PaintingRepo repositories.PaintingRepository
}

func (s *CartService) AddToCart(
	userID uint,
	req dto.AddToCartRequest,
) error {

	// Check painting exists
	painting, err := s.PaintingRepo.GetPaintingByID(req.PaintingID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("painting not found")
		}

		return err
	}

	// Painting must be available
	if painting.Status != "AVAILABLE" {
		return errors.New("painting is no longer available")
	}

	// Prevent duplicate
	_, err = s.CartRepo.GetCartItem(
		userID,
		req.PaintingID,
	)

	if err == nil {
		return errors.New("painting is already in your cart")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	cart := models.Cart{
		UserID:     userID,
		PaintingID: req.PaintingID,
	}

	return s.CartRepo.AddToCart(&cart)
}

func (s *CartService) GetCart(
	userID uint,
) ([]models.Cart, float64, error) {

	cart, err := s.CartRepo.GetCartByUserID(userID)

	if err != nil {
		return nil, 0, err
	}

	total := 0.0

	for _, item := range cart {
		total += item.Painting.Price
	}

	return cart, total, nil
}

func (s *CartService) RemoveFromCart(
	userID uint,
	paintingID uint,
) error {

	return s.CartRepo.RemoveFromCart(
		userID,
		paintingID,
	)
}

func (s *CartService) ClearCart(
	userID uint,
) error {

	return s.CartRepo.ClearCart(userID)
}