package module

import (
	"auth-service-backend/internal/controllers"
	"auth-service-backend/internal/interfaces"
	"auth-service-backend/internal/repository"
	"auth-service-backend/internal/services"
	"auth-service-backend/pkg/sqlcqueries"

	"github.com/gofiber/fiber/v2"
)

type UserModule struct {
	q    *sqlcqueries.Queries
	app  *fiber.App
	repo interfaces.UserRepository
}

func NewUserModule(q *sqlcqueries.Queries, app *fiber.App) *UserModule {
	return &UserModule{
		q:   q,
		app: app,
	}
}

func (m *UserModule) Initialization() error {
	m.repo = repository.NewUserRepository(m.q)
	service := services.NewUserService(m.repo)
	controller := controllers.NewUserController(service)

	m.app.Post("/api/users", controller.CreateUser)
	m.app.Get("/api/users", controller.GetUsers)
	m.app.Get("/api/users/:id", controller.GetUserByID)
	m.app.Put("/api/users/:id", controller.UpdateUser)
	m.app.Post("/api/users/avatar", controller.UpdateAvatar)
	m.app.Delete("/api/users/:id", controller.DeleteUser)

	return nil
}

func (m *UserModule) GetRepository() interfaces.UserRepository {
	return m.repo
}
