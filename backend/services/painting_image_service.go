package services

import (
	"context"
	"errors"
	"io"
	"fmt"

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
	file io.Reader,
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

	fmt.Printf("Cloudinary result: %+v\n", result)
	fmt.Printf("Cloudinary error: %v\n", err)

	if err != nil {
		return nil, err
	}

	// Check Cloudinary response
	if result.SecureURL == "" {
		return nil, errors.New("cloudinary upload returned an empty URL")
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

func (s *PaintingImageService) GetImagesByPaintingID(
	paintingID uint,
) ([]models.PaintingImage, error) {

	// First check that the painting exists
	_, err := s.PaintingRepo.GetPaintingByID(paintingID)

	if err != nil {
		return nil, errors.New("painting not found")
	}

	images, err := s.PaintingImageRepo.GetImagesByPaintingID(paintingID)

	if err != nil {
		return nil, err
	}

	return images, nil
}