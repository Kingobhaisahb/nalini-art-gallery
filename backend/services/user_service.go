package services

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Kingobhaisahb/nalini-art-gallery/models"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
)

type UserService struct {
	UserRepo  repositories.UserRepository
	OrderRepo repositories.OrderRepository
}

func (s *UserService) GetAllUsers() ([]models.User, error) {

	return s.UserRepo.GetAllUsers()
}

func (s *UserService) GetUserByID(
	id uint,
) (*models.User, []models.Order, error) {

	user, err := s.UserRepo.GetUserByID(id)

	if err != nil {
		return nil, nil, err
	}

	orders, err := s.OrderRepo.GetOrdersByUserID(id)

	if err != nil {
		return nil, nil, err
	}

	return user, orders, nil
}

func (s *UserService) UpdateUserRole(
	adminID uint,
	userID uint,
	role string,
) error {

	if role != "ADMIN" && role != "CUSTOMER" {
		return errors.New("invalid user role")
	}

	// Prevent an admin from removing their own admin access.
	if adminID == userID {
		return errors.New("you cannot change your own role")
	}

	err := s.UserRepo.UpdateUserRole(userID, role)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}

		return err
	}

	return nil
}