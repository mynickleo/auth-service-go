package repository

import (
	"context"
	"errors"
	"fmt"
	"auth-service-backend/internal/database/redis"
	"auth-service-backend/internal/models"
)

type CodeRepository struct {
	redis redis.RedisCache
}

func NewCodeRepository(redis redis.RedisCache) *CodeRepository {
	return &CodeRepository{
		redis: redis,
	}
}

func (r *CodeRepository) Create(ctx context.Context, model *models.Code) error {
	err := r.redis.Set(fmt.Sprint(*model.Code), model.Email)

	if err != nil {
		return err
	}

	return nil
}

func (r *CodeRepository) GetByCode(ctx context.Context, code int16) (string, error) {
	email, err := r.redis.Get(fmt.Sprint(code))
	if err != nil {
		return "", err
	}

	if email == "" {
		return "", errors.New("this code doesn't exists")
	}

	return email, nil
}

func (r *CodeRepository) Delete(ctx context.Context, code int16) error {
	err := r.redis.Delete(fmt.Sprint(code))

	return err
}
