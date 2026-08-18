package services

import (
	"context"
	"errors"
	"io"

	"github.com/Kingobhaisahb/nalini-art-gallery/config"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type PaintingVideoService struct {
	PaintingRepo repositories.PaintingRepository
}

func (s *PaintingVideoService) UploadVideo(
	paintingID uint,
	file io.Reader,
) (string, error) {

	// Check that the painting exists
	painting, err := s.PaintingRepo.GetPaintingByID(paintingID)

	if err != nil {
		return "", errors.New("painting not found")
	}

	// Get Cloudinary client
	cld, err := config.GetCloudinary()

	if err != nil {
		return "", err
	}

	// Upload video to Cloudinary
	result, err := cld.Upload.Upload(
		context.Background(),
		file,
		uploader.UploadParams{
			Folder:       "nalini-art-gallery/paintings/videos",
			ResourceType: "video",
		},
	)

	if err != nil {
		return "", err
	}

	// Cloudinary can return an error inside the result
	if result.Error.Message != "" {
		return "", errors.New(result.Error.Message)
	}

	// Make sure Cloudinary returned a URL
	if result.SecureURL == "" {
		return "", errors.New("cloudinary video upload returned an empty URL")
	}

	// Save video URL to painting
	painting.ProcessVideoURL = &result.SecureURL

	err = s.PaintingRepo.UpdatePainting(painting)

	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}