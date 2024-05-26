package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Success(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(Response{
		Success: false,
		Message: err.Error(),
	})
}
