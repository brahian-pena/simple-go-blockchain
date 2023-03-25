package server

import "github.com/gofiber/fiber/v2"

func StartListening() {
	app := fiber.New()

	initRouter(app)

	app.Listen(":3000")
}
