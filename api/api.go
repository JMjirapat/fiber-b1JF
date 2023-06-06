package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitAPI(app *fiber.App) {
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	app.Use(recover.New())

	path := app.Group("/api")
	bindDeviceRouter(path)
	bindLineBotRouter(path)
	bindAccountRouter(path)
}
