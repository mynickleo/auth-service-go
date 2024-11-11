package module

import (
	"auth-service-backend/config"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

type JWTModule struct {
}

func NewJWTModule() *JWTModule {
	return &JWTModule{}
}

func (jwtm *JWTModule) JWTGuard() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.AppConfig.SecretKey),
		ContextKey:   "user",
		ErrorHandler: jwtm.jwtErrorHandler,
		TokenLookup:  "header:x-access-token",
	})
}

func (jwtm *JWTModule) jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	return nil
}

func (jwtm *JWTModule) CheckUserGuard(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, idOk := claims["user.id"].(string)
	role, roleOk := claims["user.role"].(string)

	if !idOk || !roleOk || id == "" || role == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token payload"})
	}

	return c.Next()
}

func (jwtm *JWTModule) GenerateToken(userId string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user.id":   userId,
		"user.role": role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.AppConfig.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
