package main

import (
	"github.com/den19980107/go-fiber-gorm-starter/config"
	"github.com/den19980107/go-fiber-gorm-starter/db"
	"github.com/den19980107/go-fiber-gorm-starter/log"
	"github.com/den19980107/go-fiber-gorm-starter/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadConfigFromEnv()

	log.SetupLogger()

	db.Connect()
	db.Migrate()

	app := fiber.New()

	routers.RegisterRoutes(app)

	app.Listen(config.App.Server.Host + ":" + config.App.Server.Port)
}
