package models

import "github.com/google/uuid"

type Role struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name" validate:"required"`
}

type UserRole struct {
	ID      uuid.UUID `json:"id"`
	User_ID uuid.UUID `json:"user_id" validate:"required"`
	Role_ID uuid.UUID `json:"role_id" validate:"required"`
}
