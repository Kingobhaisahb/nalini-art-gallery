package repositories

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
    "github.com/Kingobhaisahb/nalini-art-gallery/models"
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {

	var user models.User

	err := database.DB.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {

	var user models.User

	err := database.DB.
		First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}