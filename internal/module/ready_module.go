package module

import "github.com/gofiber/fiber/v2"

type ReadyModule struct {
	app *fiber.App
}

func NewReadyModule(app *fiber.App) *ReadyModule {
	return &ReadyModule{app: app}
}

func (m *ReadyModule) Initialization() error {
	m.app.Get("/api/ready", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Server is running"})
	})

	return nil
}
