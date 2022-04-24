package response

import (
	"github.com/gofiber/fiber/v2"
)

type CommonResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func Success(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(data)
}

func Error(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(ErrorResponse{
		Message: err.Error(),
	})
}
