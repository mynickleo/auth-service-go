package repository

import (
	"context"
	"auth-service-backend/internal/models"
	"auth-service-backend/pkg/sqlcqueries"

	"github.com/google/uuid"
)

type RoleRepository struct {
	q *sqlcqueries.Queries
}

func NewRoleRepository(q *sqlcqueries.Queries) *RoleRepository {
	return &RoleRepository{
		q: q,
	}
}

func (r *RoleRepository) CreateRole(ctx context.Context, role *models.Role) error {
	err := r.q.CreateRole(ctx, sqlcqueries.CreateRoleParams{
		Name: role.Name,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *RoleRepository) GetRoles(ctx context.Context) ([]*models.Role, error) {
	rolesRow, err := r.q.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	roles := make([]*models.Role, len(rolesRow))
	for i, value := range rolesRow {
		roles[i] = &models.Role{
			ID:   value.ID,
			Name: value.Name,
		}
	}

	return roles, nil
}

func (r *RoleRepository) GetRoleByName(ctx context.Context, name string) (*models.Role, error) {
	roleRow, err := r.q.GetRoleByName(ctx, sqlcqueries.GetRoleByNameParams{Name: name})
	if err != nil {
		return nil, err
	}

	return &models.Role{
		ID:   roleRow.ID,
		Name: roleRow.Name,
	}, nil
}

func (r *RoleRepository) GetRoleByID(ctx context.Context, id uuid.UUID) (*models.Role, error) {
	roleRow, err := r.q.GetRoleByID(ctx, sqlcqueries.GetRoleByIDParams{
		ID: id,
	})

	if err != nil {
		return nil, err
	}

	return &models.Role{
		ID:   roleRow.ID,
		Name: roleRow.Name,
	}, nil
}

func (r *RoleRepository) Update(ctx context.Context, role *models.Role) error {
	err := r.q.UpdateRole(ctx, sqlcqueries.UpdateRoleParams{
		Name: role.Name,
		ID:   role.ID,
	})

	return err
}

func (r *RoleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.q.DeleteRole(ctx, sqlcqueries.DeleteRoleParams{
		ID: id,
	})

	return err
}

func (r *RoleRepository) CreateUserRole(ctx context.Context, userRole *models.UserRole) error {
	err := r.q.CreateUserRole(ctx, sqlcqueries.CreateUserRoleParams{
		UserID: &userRole.User_ID,
		RoleID: &userRole.Role_ID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *RoleRepository) GetRoleByUserId(ctx context.Context, userId uuid.UUID) (*models.Role, error) {
	userRoleRow, err := r.q.GetUserRoleByUserID(ctx, sqlcqueries.GetUserRoleByUserIDParams{
		UserID: &userId,
	})

	if err != nil {
		return nil, err
	}
	return &models.Role{
		ID:   *userRoleRow.RoleID,
		Name: *userRoleRow.RoleName,
	}, nil
}
