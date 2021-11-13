package utilties

import "github.com/gofiber/fiber/v2"

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	err = ctx.Status(code).JSON(fiber.Map{
		"status":      code,
		"description": err.Error(),
	})
	if err != nil {
		return ctx.Status(code).JSON(fiber.Map{
			"status":      500,
			"description": err.Error(),
		})
	}
	return nil
}
