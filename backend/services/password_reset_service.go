package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"

	"golang.org/x/crypto/bcrypt"
)

type PasswordResetService struct {
	UserRepo          repositories.UserRepository
	PasswordResetRepo repositories.PasswordResetRepository
}

// Generates a secure random token
func generateResetToken() (string, error) {

	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

// Forgot Password
func (s *PasswordResetService) ForgotPassword(req dto.ForgotPasswordRequest) (string, error) {

	user, err := s.UserRepo.GetUserByEmail(req.Email)

	if err != nil || user == nil {
		return "", errors.New("user not found")
	}

	token, err := generateResetToken()

	if err != nil {
		return "", err
	}

	resetToken := models.PasswordResetToken{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}

	err = s.PasswordResetRepo.CreateToken(&resetToken)

	if err != nil {
		return "", err
	}

	return token, nil
}

// Reset Password
func (s *PasswordResetService) ResetPassword(req dto.ResetPasswordRequest) error {

	if req.NewPassword != req.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	resetToken, err := s.PasswordResetRepo.GetToken(req.Token)

	if err != nil {
		return err
	}

	if resetToken == nil {
		return errors.New("invalid reset token")
	}

	if time.Now().After(resetToken.ExpiresAt) {
		return errors.New("reset token has expired")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user, err := s.UserRepo.GetUserByID(resetToken.UserID)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = s.UserRepo.UpdateUser(user)

	if err != nil {
		return err
	}

	err = s.PasswordResetRepo.DeleteToken(req.Token)

	if err != nil {
		return err
	}

	return nil
}