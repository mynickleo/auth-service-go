package interfaces

import (
	"context"
	"auth-service-backend/internal/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.GetUserDto, error)
	GetUsers(ctx context.Context) ([]*models.GetUserDto, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.GetUserDto, error)
	GetByEmail(ctx context.Context, email string) (*models.UserWithRole, error)
	Update(ctx context.Context, user *models.User) error
	UpdateAvatar(ctx context.Context, url string, id uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type CodeRepository interface {
	Create(ctx context.Context, model *models.Code) error
	GetByCode(ctx context.Context, code int16) (string, error)
	Delete(ctx context.Context, code int16) error
}

type RoleRepository interface {
	CreateRole(ctx context.Context, role *models.Role) error
	GetRoles(ctx context.Context) ([]*models.Role, error)
	GetRoleByID(ctx context.Context, id uuid.UUID) (*models.Role, error)
	GetRoleByName(ctx context.Context, name string) (*models.Role, error)
	Update(ctx context.Context, role *models.Role) error
	Delete(ctx context.Context, id uuid.UUID) error
	CreateUserRole(ctx context.Context, userRole *models.UserRole) error
	GetRoleByUserId(ctx context.Context, userId uuid.UUID) (*models.Role, error)
}
