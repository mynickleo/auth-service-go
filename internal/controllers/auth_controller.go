package controllers

import (
	"context"
	"auth-service-backend/internal/interfaces"
	"auth-service-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service interfaces.AuthService
}

func NewAuthController(s interfaces.AuthService) *AuthController {
	return &AuthController{
		service: s,
	}
}

func (c *AuthController) SendMail(ctx *fiber.Ctx) error {
	dto := new(models.CreateMailDto)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Please enter a all fields"})
	}

	if err := validate.Struct(dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": err.Error()})
	}

	if err := c.service.SendMail(context.Background(), dto); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Mail was sent to email. Please check"})
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	dto := new(models.CreateUserDto)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Please enter a all fields"})
	}

	if err := validate.Struct(dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": err.Error()})
	}

	jwtToken, err := c.service.Register(context.Background(), dto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"x-access-token": jwtToken})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	dto := new(models.LoginDto)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Please enter a all fields"})
	}

	if err := validate.Struct(dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": err.Error()})
	}

	jwtToken, err := c.service.Login(context.Background(), dto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"x-access-token": jwtToken})
}
