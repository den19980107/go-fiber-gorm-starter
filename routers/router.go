package routers

import (
	"github.com/den19980107/go-fiber-gorm-starter/routers/api"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	apiGroup := app.Group("api")

	apiGroup.Post("/register", api.ValidRigisterRequest, api.Register)
	apiGroup.Post("/login", api.ValidLoginRequest, api.Login)
}
