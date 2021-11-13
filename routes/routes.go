package routes

import "github.com/gofiber/fiber/v2"

func Setup(api fiber.Router) {
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
