package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"simpleCrud/routes"
	"simpleCrud/utilties"
)

func main() {
	utilties.Connect()

	app := fiber.New(fiber.Config{
		ErrorHandler: utilties.CustomErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	routes.Setup(app)

	app.All("*", func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)

}
