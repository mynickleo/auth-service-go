package models

import "github.com/google/uuid"

type Code struct {
	ID    uuid.UUID `json:"id"`
	Code  *int16    `json:"code"`
	Email string    `json:"email"`
}
