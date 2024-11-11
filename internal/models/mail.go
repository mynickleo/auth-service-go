package models

type CreateMailDto struct {
	Email string `json:"email" validate:"required"`
}
