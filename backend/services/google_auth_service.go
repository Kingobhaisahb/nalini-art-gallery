package services

import (
	"context"
	"errors"
	"os"

	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
	"github.com/Kingobhaisahb/nalini-art-gallery/utils"

	"google.golang.org/api/idtoken"
)

type GoogleAuthService struct {
	UserRepo repositories.UserRepository
}

func (s *GoogleAuthService) GoogleLogin(req dto.GoogleLoginRequest) (*dto.AuthResponse, error) {

	payload, err := idtoken.Validate(
		context.Background(),
		req.IDToken,
		os.Getenv("GOOGLE_CLIENT_ID"),
	)

	if err != nil {
		return nil, errors.New("invalid google token")
	}

	email, ok := payload.Claims["email"].(string)

	if !ok {
		return nil, errors.New("email not found")
	}

	name, _ := payload.Claims["name"].(string)
	googleID, _ := payload.Claims["sub"].(string)

	user, err := s.UserRepo.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	// -----------------------------
	// CASE 1
	// New User
	// -----------------------------
	if user == nil {

		newUser := models.User{
			Name:         name,
			Email:        email,
			Role:         "CUSTOMER",
			GoogleID:     &googleID,
			AuthProvider: "GOOGLE",
		}

		err = s.UserRepo.CreateUser(&newUser)

		if err != nil {
			return nil, err
		}

		user = &newUser

	} else {

		// -----------------------------
		// CASE 2
		// Existing LOCAL account
		// -----------------------------

		if user.GoogleID == nil {

			user.GoogleID = &googleID

			err = s.UserRepo.UpdateUser(user)

			if err != nil {
				return nil, err
			}
		}

		// CASE 3
		// Google already linked
		// Nothing to do.
	}

	token, err := utils.GenerateJWT(
		user.ID,
		user.Email,
		user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{

		Success: true,
		Message: "Google login successful",
		Token:   token,

		User: dto.UserResponse{
			ID:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			Phone:        user.Phone,
			Role:         user.Role,
			AuthProvider: user.AuthProvider,
		},
	}, nil
}