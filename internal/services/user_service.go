package services

import (
	"context"
	"fmt"
	"auth-service-backend/internal/interfaces"
	"auth-service-backend/internal/models"
	"auth-service-backend/internal/utils"
	"auth-service-backend/pkg/minio"

	"github.com/google/uuid"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	_, err = s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]*models.GetUserDto, error) {
	return s.repo.GetUsers(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*models.GetUserDto, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.repo.Update(ctx, user)
}

func (s *UserService) UpdateAvatar(ctx context.Context, userId uuid.UUID, file []byte, fileName string) error {
	_, err := minio.PutData(ctx, file, "test", fileName)
	if err != nil {
		return fmt.Errorf("failed to upload file to MinIO: %w", err)
	}

	url, err := minio.GetDataUrl(ctx, fmt.Sprintf("%s/%s", "test", fileName))
	if err != nil {
		return fmt.Errorf("failed to retrieve file URL from MinIO: %w", err)
	}

	if err := s.repo.UpdateAvatar(ctx, url, userId); err != nil {
		return fmt.Errorf("failed to update avatar URL in database: %w", err)
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
