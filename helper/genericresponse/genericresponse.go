package genericresponse

import (
	"github.com/gofiber/fiber/v2"
)

type GenericResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Success(c *fiber.Ctx, data interface{}, status int, msg ...string) error {
	if len(msg) == 0 {
		msg = append(msg, "Success")
	}
	message := msg[0]

	return c.Status(status).JSON(
		GenericResponse{
			Message: message,
			Data:    data,
			Success: true,
		})
}

func Error(c *fiber.Ctx, status int, msg string, data ...interface{}) error {
	if len(data) == 0 {
		data = append(data, []interface{}{})
	}

	resData := data[0]

	return c.Status(status).JSON(
		GenericResponse{
			Message: msg,
			Data:    resData,
			Success: false,
		})
}
