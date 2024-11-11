package app

import (
	"auth-service-backend/internal/database/redis"
	"auth-service-backend/internal/interfaces"
	"auth-service-backend/internal/module"
	"auth-service-backend/pkg/sqlcqueries"

	"github.com/gofiber/fiber/v2"
)

type DIContainer struct {
	queries     *sqlcqueries.Queries
	app         *fiber.App
	jwtModule   interfaces.JWTModule
	mailModule  interfaces.MailModule
	userModule  interfaces.UserModule
	roleModule  interfaces.RoleModule
	readyModule interfaces.ReadyModule
	authModule  interfaces.AuthModule
	redis       redis.RedisCache
}

func NewDIContainer(q *sqlcqueries.Queries, app *fiber.App, redis redis.RedisCache) *DIContainer {
	return &DIContainer{
		queries:    q,
		app:        app,
		jwtModule:  module.NewJWTModule(),
		mailModule: module.NewEmailModule(),
		redis:      redis,
	}
}

func (di *DIContainer) InitializationModules() error {
	err := di.InitializationReadyModule()
	if err != nil {
		return err
	}

	err = di.InitializationUserModule()
	if err != nil {
		return err
	}

	err = di.InitializationRoleModule()
	if err != nil {
		return err
	}

	err = di.InitializationAuthModule()
	if err != nil {
		return err
	}

	return nil
}

func (di *DIContainer) InitializationReadyModule() error {
	di.readyModule = module.NewReadyModule(di.app)
	err := di.readyModule.Initialization()

	if err != nil {
		return err
	}

	return nil
}

func (di *DIContainer) InitializationUserModule() error {
	di.app.Group("/api/users", di.jwtModule.JWTGuard(), di.jwtModule.CheckUserGuard)

	di.userModule = module.NewUserModule(di.queries, di.app)
	err := di.userModule.Initialization()

	if err != nil {
		return err
	}

	return nil
}

func (di *DIContainer) InitializationAuthModule() error {
	di.authModule = module.NewAuthModule(di.app, di.redis, di.mailModule, di.userModule.GetRepository(), di.roleModule.GetRepository(), di.jwtModule)
	err := di.authModule.Initialization()

	if err != nil {
		return err
	}

	return nil
}

func (di *DIContainer) InitializationRoleModule() error {
	di.app.Group("/api/roles", di.jwtModule.JWTGuard(), di.jwtModule.CheckUserGuard)
	di.app.Group("/api/user-roles", di.jwtModule.JWTGuard(), di.jwtModule.CheckUserGuard)

	di.roleModule = module.NewRoleModule(di.queries, di.app)
	err := di.roleModule.Initialization()

	if err != nil {
		return err
	}

	return nil
}
