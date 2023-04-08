package baserouter

import (
	BaseController "localizer/cmd/basecontroller"

	"github.com/gofiber/fiber/v2"
)

func InitializeRouters(app fiber.Router) {

	root := app.Group("/")
	BaseController.HealthCheck(root)

	// api := app.Group("api")

}
