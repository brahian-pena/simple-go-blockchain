package server

import (
	"github.com/gofiber/fiber/v2"

	"github.com/brahian-pena/simple-go-blockchain/controllers"
)

func initRouter(app *fiber.App) {
	app.Get("/_health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	v1 := app.Group("/api/v1")

	v1.Post("/blockchain/mine", controllers.MineBlock)

	v1.Get("/blockchain", controllers.GetBlockChain)

	v1.Get("/blockchain/valid", controllers.ValidateBlockChain)
}
