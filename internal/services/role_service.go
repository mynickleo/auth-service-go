package services

import (
	"context"
	"auth-service-backend/internal/interfaces"
	"auth-service-backend/internal/models"

	"github.com/google/uuid"
)

type RoleService struct {
	repo interfaces.RoleRepository
}

func NewRoleService(repo interfaces.RoleRepository) *RoleService {
	return &RoleService{
		repo: repo,
	}
}

func (s *RoleService) CreateUserRole(ctx context.Context, userRole *models.UserRole) error {
	return s.repo.CreateUserRole(ctx, userRole)
}

func (s *RoleService) GetRoleByUserId(ctx context.Context, userId uuid.UUID) (*models.Role, error) {
	return s.repo.GetRoleByUserId(ctx, userId)
}

func (s *RoleService) CreateRole(ctx context.Context, role *models.Role) error {
	return s.repo.CreateRole(ctx, role)
}

func (s *RoleService) GetRoles(ctx context.Context) ([]*models.Role, error) {
	return s.repo.GetRoles(ctx)
}

func (s *RoleService) GetRoleByID(ctx context.Context, id uuid.UUID) (*models.Role, error) {
	return s.repo.GetRoleByID(ctx, id)
}

func (s *RoleService) UpdateRole(ctx context.Context, role *models.Role) error {
	return s.repo.Update(ctx, role)
}

func (s *RoleService) DeleteRole(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
