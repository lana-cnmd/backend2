package service

import (
	"github.com/google/uuid"
	"github.com/lana-cnmd/backend2/pkg/repository"
)

type ImageService struct {
	repo repository.IImageRepo
}

func NewImageService(repo repository.IImageRepo) *ImageService {
	return &ImageService{
		repo: repo}
}

func (s *ImageService) AddImage(fileBytes []byte) (uuid.UUID, error) {
	return s.repo.AddImage(fileBytes)
}

func (s *ImageService) GetImageByImageUUID(imageUUID uuid.UUID) ([]byte, error) {
	return s.repo.GetImageByImageUUID(imageUUID)
}

func (s *ImageService) DeleteImageByImageUUID(imageUUID uuid.UUID) error {
	return s.repo.DeleteImageByImageUUID(imageUUID)
}

func (s *ImageService) UpdateImage(imageId uuid.UUID, fileBytes []byte) error {
	return s.repo.UpdateImage(imageId, fileBytes)
}

func (s *ImageService) GetImageByProductId(productId int) ([]byte, error) {
	return s.repo.GetImageByProductId(productId)
}
