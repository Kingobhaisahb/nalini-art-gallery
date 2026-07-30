package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	"github.com/Kingobhaisahb/nalini-art-gallery/utils"
)

type AuthService struct {
	UserRepo repositories.UserRepository
}

func (s *AuthService) Signup(req dto.SignupRequest) (*models.User, error) {

	existingUser, _ := s.UserRepo.GetUserByEmail(req.Email)

	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		Password:     string(hashedPassword),
		Phone:        req.Phone,
		Role:         "CUSTOMER",
		AuthProvider: "LOCAL",
	}

	err = s.UserRepo.CreateUser(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {

	user, err := s.UserRepo.GetUserByEmail(req.Email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(
		user.ID,
		user.Email,
		user.Role,
	)

	if err != nil {
		return nil, err
	}

	response := dto.AuthResponse{
		Success: true,
		Message: "Login successful",
		Token:   token,
		User: dto.UserResponse{
			ID:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			Phone:        user.Phone,
			Role:         user.Role,
			AuthProvider: user.AuthProvider,
		},
	}

	return &response, nil
}

func (s *AuthService) GetProfile(userID uint) (*dto.UserResponse, error) {

	user, err := s.UserRepo.GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		Role:         user.Role,
		AuthProvider: user.AuthProvider,
	}

	return &response, nil
}