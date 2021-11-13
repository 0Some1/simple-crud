package controller

import (
	"github.com/gofiber/fiber/v2"
	"simpleCrud/models"
	"simpleCrud/utilties"
)

func GetComments(c *fiber.Ctx) error {
	var comments []models.Comment
	err := utilties.DB.Find(&comments).Error
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(comments)
}

func CreateComment(c *fiber.Ctx) error {
	var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return fiber.ErrBadRequest
	}
	if err := utilties.DB.Create(&comment).Error; err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(comment)
}

func UpdateComment(c *fiber.Ctx) error {
	id := c.Params("id")
	var comment models.Comment
	if err := utilties.DB.First(&comment, id).Error; err != nil {
		return fiber.ErrNotFound
	}
	if err := c.BodyParser(&comment); err != nil {
		return fiber.ErrBadRequest
	}
	if err := utilties.DB.Save(&comment).Error; err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(comment)
}

func DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")
	var comment models.Comment
	if err := utilties.DB.First(&comment, id).Error; err != nil {
		return fiber.ErrNotFound
	}
	if err := utilties.DB.Delete(&comment).Error; err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(comment)
}
