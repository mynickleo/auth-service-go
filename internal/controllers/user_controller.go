package controllers

import (
	"context"
	"auth-service-backend/internal/interfaces"
	"auth-service-backend/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type UserController struct {
	service interfaces.UserService
}

var validate = validator.New()

func NewUserController(s interfaces.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	if err := validate.Struct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": err.Error()})
	}

	if err := c.service.Create(context.Background(), user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	users, err := c.service.GetUsers(context.Background())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}

func (c *UserController) GetUserByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Uncorrect id"})
	}

	user, err := c.service.GetByID(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	if err := validate.Struct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": err.Error()})
	}

	user.ID = id
	if err := c.service.Update(context.Background(), user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	user.Password = ""
	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func (c *UserController) UpdateAvatar(ctx *fiber.Ctx) error {
	userToken := ctx.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	id, ok := claims["user.id"].(string)
	if !ok || id == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token payload"})
	}

	userId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	fileHeader, err := ctx.FormFile("avatar")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to retrieve avatar file"})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to open avatar file"})
	}
	defer file.Close()

	fileData := make([]byte, fileHeader.Size)
	_, err = file.Read(fileData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to read avatar file"})
	}

	err = c.service.UpdateAvatar(context.Background(), userId, fileData, fileHeader.Filename)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update avatar", "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Avatar updated successfully"})
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	if err := c.service.Delete(context.Background(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "User deleted successfully"})
}
