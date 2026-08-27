package repositories

import (

	"errors"
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
    "github.com/Kingobhaisahb/nalini-art-gallery/models"
	"gorm.io/gorm"

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

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

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

func (r *UserRepository) GetAllUsers() ([]models.User, error) {

	var users []models.User

	err := database.DB.
		Order("created_at DESC").
		Find(&users).Error

	return users, err
}

func (r *UserRepository) UpdateUserRole(
	id uint,
	role string,
) error {

	result := database.DB.
		Model(&models.User{}).
		Where("id = ?", id).
		Update("role", role)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

