package module

import (
	"auth-service-backend/internal/controllers"
	"auth-service-backend/internal/interfaces"
	"auth-service-backend/internal/repository"
	"auth-service-backend/internal/services"
	"auth-service-backend/pkg/sqlcqueries"

	"github.com/gofiber/fiber/v2"
)

type RoleModule struct {
	q    *sqlcqueries.Queries
	app  *fiber.App
	repo interfaces.RoleRepository
}

func NewRoleModule(q *sqlcqueries.Queries, app *fiber.App) *RoleModule {
	return &RoleModule{
		q:   q,
		app: app,
	}
}

func (m *RoleModule) Initialization() error {
	m.repo = repository.NewRoleRepository(m.q)
	service := services.NewRoleService(m.repo)
	controller := controllers.NewRoleController(service)

	m.app.Post("/api/roles", controller.CreateRole)
	m.app.Post("/api/user-roles", controller.CreateUserRole)
	m.app.Get("/api/roles", controller.GetRoles)
	m.app.Get("/api/roles/:id", controller.GetRoleByID)
	m.app.Get("/api/user-role/:id", controller.GetRoleByUserId)
	m.app.Put("/api/roles", controller.UpdateRole)
	m.app.Delete("/api/roles", controller.DeleteRole)

	return nil
}

func (m *RoleModule) GetRepository() interfaces.RoleRepository {
	return m.repo
}
