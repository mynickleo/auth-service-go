package module

import (
	"auth-service-backend/internal/controllers"
	"auth-service-backend/internal/database/redis"
	"auth-service-backend/internal/interfaces"
	"auth-service-backend/internal/repository"
	"auth-service-backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthModule struct {
	app        *fiber.App
	redis      redis.RedisCache
	mailModule interfaces.MailModule
	userRepo   interfaces.UserRepository
	roleRepo   interfaces.RoleRepository
	jwtModule  interfaces.JWTModule
}

func NewAuthModule(
	app *fiber.App,
	redis redis.RedisCache,
	mailModule interfaces.MailModule,
	userRepo interfaces.UserRepository,
	roleRepo interfaces.RoleRepository,
	jwtModule interfaces.JWTModule,
) *AuthModule {
	return &AuthModule{
		app:        app,
		redis:      redis,
		mailModule: mailModule,
		userRepo:   userRepo,
		roleRepo:   roleRepo,
		jwtModule:  jwtModule,
	}
}

func (m *AuthModule) Initialization() error {
	repo := repository.NewCodeRepository(m.redis)
	service := services.NewAuthService(repo, m.userRepo, m.roleRepo, m.mailModule, m.jwtModule)
	controller := controllers.NewAuthController(service)

	m.app.Post("/api/auth/send-mail", controller.SendMail)
	m.app.Post("/api/auth/register", controller.Register)
	m.app.Post("/api/auth/login", controller.Login)

	return nil
}
