package config

import (
	GenericResponse "localizer/helper/genericresponse"
	"localizer/helper/logger"

	"github.com/gofiber/fiber/v2"
)

var FiberConfig = fiber.Config{
	AppName:   "TGA Minio Service",
	BodyLimit: 50 * 1024 * 1024, // 50 MB

	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's an fiber.*Error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		logger.Errorf("Unexpected error is occured ", err.Error())
		// logger.CreateLog(ctx, time.Since(start).String(), err.Error())

		return GenericResponse.Error(ctx, code, "Unexpected error is occured")
	},
}
