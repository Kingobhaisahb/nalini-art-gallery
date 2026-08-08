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

	// Verify the Google ID token
	payload, err := idtoken.Validate(
		context.Background(),
		req.IDToken,
		os.Getenv("GOOGLE_CLIENT_ID"),
	)

	if err != nil {
		return nil, errors.New("invalid google authentication token")
	}

	// Extract email
	email, ok := payload.Claims["email"].(string)

	if !ok || email == "" {
		return nil, errors.New("email not found in google account")
	}

	// Make sure Google has verified the email
	emailVerified, ok := payload.Claims["email_verified"].(bool)

	if !ok || !emailVerified {
		return nil, errors.New("google email is not verified")
	}

	// Extract Google user information
	name, _ := payload.Claims["name"].(string)
	googleID, ok := payload.Claims["sub"].(string)

	if !ok || googleID == "" {
		return nil, errors.New("google account ID not found")
	}

	// Find existing user using email
	user, err := s.UserRepo.GetUserByEmail(email)

	if err != nil {
		return nil, errors.New("failed to check existing user")
	}

	// CASE 1:
	// User does not exist -> create a new Google user
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
			return nil, errors.New("failed to create user")
		}

		user = &newUser

	} else {

		// CASE 2:
		// User exists but Google is not linked yet
		if user.GoogleID == nil {

			user.GoogleID = &googleID

			err = s.UserRepo.UpdateUser(user)

			if err != nil {
				return nil, errors.New("failed to link google account")
			}

		} else {

			// CASE 3:
			// Google is already linked.
			// Make sure it is the SAME Google account.
			if *user.GoogleID != googleID {
				return nil, errors.New("this email is already linked to another google account")
			}
		}
	}

	// Generate our application's JWT
	token, err := utils.GenerateJWT(
		user.ID,
		user.Email,
		user.Role,
	)

	if err != nil {
		return nil, errors.New("failed to generate authentication token")
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