package controllers

import (
	"context"
	"auth-service-backend/internal/interfaces"
	"auth-service-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoleController struct {
	s interfaces.RoleService
}

func NewRoleController(s interfaces.RoleService) *RoleController {
	return &RoleController{
		s: s,
	}
}

func (c *RoleController) CreateRole(ctx *fiber.Ctx) error {
	role := new(models.Role)

	if err := ctx.BodyParser(role); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	if err := validate.Struct(role); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": err.Error()})
	}

	if err := c.s.CreateRole(context.Background(), role); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return ctx.Status(fiber.StatusCreated).JSON("")
}

func (c *RoleController) GetRoles(ctx *fiber.Ctx) error {
	roles, err := c.s.GetRoles(context.Background())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(roles)
}

func (c *RoleController) GetRoleByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Uncorrect id"})
	}

	role, err := c.s.GetRoleByID(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(role)
}

func (c *RoleController) UpdateRole(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	role := new(models.Role)
	if err := ctx.BodyParser(role); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	if err := validate.Struct(role); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": err.Error()})
	}

	role.ID = id
	if err := c.s.UpdateRole(context.Background(), role); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return ctx.Status(fiber.StatusCreated).JSON(role)
}

func (c *RoleController) DeleteRole(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	err = c.s.DeleteRole(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Role deleted successfully"})
}

func (c *RoleController) CreateUserRole(ctx *fiber.Ctx) error {
	userRole := new(models.UserRole)

	if err := ctx.BodyParser(userRole); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	if err := validate.Struct(userRole); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": err.Error()})
	}

	if err := c.s.CreateUserRole(context.Background(), userRole); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return ctx.Status(fiber.StatusCreated).JSON(userRole)
}

func (c *RoleController) GetRoleByUserId(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	role, err := c.s.GetRoleByUserId(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(role)
}
