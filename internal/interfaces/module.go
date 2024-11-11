package interfaces

import "github.com/gofiber/fiber/v2"

type MailModule interface {
	SendEmail(to, subject, body string) error
}

type JWTModule interface {
	JWTGuard() fiber.Handler
	CheckUserGuard(c *fiber.Ctx) error
	GenerateToken(userId string, role string) (string, error)
}

type UserModule interface {
	Initialization() error
	GetRepository() UserRepository
}

type RoleModule interface {
	Initialization() error
	GetRepository() RoleRepository
}

type ReadyModule interface {
	Initialization() error
}

type AuthModule interface {
	Initialization() error
}
