package repositories

import (
	"time"

	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"

	"gorm.io/gorm"
)

type PasswordResetRepository struct{}

func (r *PasswordResetRepository) CreateToken(token *models.PasswordResetToken) error {

	return database.DB.Create(token).Error
}

func (r *PasswordResetRepository) GetToken(token string) (*models.PasswordResetToken, error) {

	var resetToken models.PasswordResetToken

	err := database.DB.Where("token = ?", token).First(&resetToken).Error

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &resetToken, nil
}

func (r *PasswordResetRepository) DeleteToken(token string) error {

	return database.DB.Where("token = ?", token).Delete(&models.PasswordResetToken{}).Error
}

func (r *PasswordResetRepository) DeleteExpiredTokens() error {

	return database.DB.
		Where("expires_at < ?", time.Now()).
		Delete(&models.PasswordResetToken{}).Error
}

func (r *PasswordResetRepository) DeleteTokensByUserID(userID uint) error {

	return database.DB.
		Where("user_id = ?", userID).
		Delete(&models.PasswordResetToken{}).Error
}