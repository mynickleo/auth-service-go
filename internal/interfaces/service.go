package interfaces

import (
	"context"
	"auth-service-backend/internal/models"

	"github.com/google/uuid"
)

type UserService interface {
	Create(ctx context.Context, user *models.User) error
	GetUsers(ctx context.Context) ([]*models.GetUserDto, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.GetUserDto, error)
	Update(ctx context.Context, user *models.User) error
	UpdateAvatar(ctx context.Context, userId uuid.UUID, file []byte, fileName string) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type AuthService interface {
	SendMail(ctx context.Context, dto *models.CreateMailDto) error
	Register(ctx context.Context, dto *models.CreateUserDto) (string, error)
	Login(ctx context.Context, dto *models.LoginDto) (string, error)
}

type RoleService interface {
	CreateUserRole(ctx context.Context, userRole *models.UserRole) error
	GetRoleByUserId(ctx context.Context, userId uuid.UUID) (*models.Role, error)
	CreateRole(ctx context.Context, role *models.Role) error
	GetRoles(ctx context.Context) ([]*models.Role, error)
	GetRoleByID(ctx context.Context, id uuid.UUID) (*models.Role, error)
	UpdateRole(ctx context.Context, role *models.Role) error
	DeleteRole(ctx context.Context, id uuid.UUID) error
}
