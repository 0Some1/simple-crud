package routes

import (
	"github.com/gofiber/fiber/v2"
	"simpleCrud/controller"
)

func Setup(api fiber.Router) {
	api.Get("/comment", controller.GetComments)
	api.Post("/comment", controller.CreateComment)
	api.Put("/comment/:id", controller.UpdateComment)
	api.Delete("/comment/:id", controller.DeleteComment)
}
