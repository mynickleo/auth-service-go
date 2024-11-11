package repository

import (
	"context"
	"auth-service-backend/internal/models"
	"auth-service-backend/pkg/sqlcqueries"
	"time"

	"github.com/google/uuid"
)

type UserRepository struct {
	q *sqlcqueries.Queries
}

func NewUserRepository(q *sqlcqueries.Queries) *UserRepository {
	return &UserRepository{
		q: q,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.GetUserDto, error) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := r.q.CreateUser(ctx, sqlcqueries.CreateUserParams{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		FullName:  &user.FullName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})

	if err != nil {
		return nil, err
	}

	return &models.GetUserDto{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]*models.GetUserDto, error) {
	userRow, err := r.q.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*models.GetUserDto, len(userRow))
	for i, value := range userRow {
		user := &models.GetUserDto{
			ID:        value.ID,
			Email:     value.Email,
			FullName:  *value.FullName,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			AvatarImg: *value.AvatarImg,
		}
		users[i] = user
	}

	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.GetUserDto, error) {
	userRow, err := r.q.GetUserByID(ctx, sqlcqueries.GetUserByIDParams{ID: id})
	if err != nil {
		return nil, err
	}

	return &models.GetUserDto{
		ID:        userRow.ID,
		Email:     userRow.Email,
		FullName:  *userRow.FullName,
		CreatedAt: userRow.CreatedAt,
		UpdatedAt: userRow.UpdatedAt,
		AvatarImg: *userRow.AvatarImg,
	}, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.UserWithRole, error) {
	userRow, err := r.q.GetUserByEmail(ctx, sqlcqueries.GetUserByEmailParams{Email: email})
	if err != nil {
		return nil, err
	}

	return &models.UserWithRole{
		ID:        userRow.UserID,
		Email:     userRow.Email,
		Password:  userRow.Password,
		FullName:  *userRow.FullName,
		CreatedAt: userRow.CreatedAt,
		UpdatedAt: userRow.UpdatedAt,
		RoleName:  *userRow.RoleName,
		AvatarImg: *userRow.AvatarImg,
	}, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	user.UpdatedAt = time.Now()

	err := r.q.UpdateUser(ctx, sqlcqueries.UpdateUserParams{
		FullName: &user.FullName,
		Email:    user.Email,
		Password: user.Password,
		ID:       user.ID,
	})
	return err
}

func (r *UserRepository) UpdateAvatar(ctx context.Context, url string, id uuid.UUID) error {
	return r.q.UpdateUserAvatar(ctx, sqlcqueries.UpdateUserAvatarParams{AvatarImg: &url, ID: id})
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.q.DeleteUser(ctx, sqlcqueries.DeleteUserParams{ID: id})
	return err
}
