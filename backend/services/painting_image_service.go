package services

import (
	"context"
	"errors"

	"github.com/Kingobhaisahb/nalini-art-gallery/config"
	"github.com/Kingobhaisahb/nalini-art-gallery/models"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type PaintingImageService struct {
	PaintingRepo      repositories.PaintingRepository
	PaintingImageRepo repositories.PaintingImageRepository
}

func (s *PaintingImageService) UploadImage(
	paintingID uint,
	file interface{},
) (*models.PaintingImage, error) {

	// First check that the painting exists
	_, err := s.PaintingRepo.GetPaintingByID(paintingID)

	if err != nil {
		return nil, errors.New("painting not found")
	}

	// Get Cloudinary client
	cld, err := config.GetCloudinary()

	if err != nil {
		return nil, err
	}

	// Upload image to Cloudinary
	result, err := cld.Upload.Upload(
		context.Background(),
		file,
		uploader.UploadParams{
			Folder: "nalini-art-gallery/paintings",
		},
	)

	if err != nil {
		return nil, err
	}

	// Create database record
	image := models.PaintingImage{
		PaintingID: paintingID,
		ImageURL:  result.SecureURL,
	}

	err = s.PaintingImageRepo.CreateImage(&image)

	if err != nil {
		return nil, err
	}

	return &image, nil
}