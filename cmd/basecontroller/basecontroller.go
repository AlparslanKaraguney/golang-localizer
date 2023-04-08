package basecontroller

import (
	"fmt"
	"localizer/helper/genericresponse"

	"localizer/internal/localizer"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(app fiber.Router) {

	app.Get("/", func(c *fiber.Ctx) error {
		l := c.Locals("localizer").(localizer.Localizer)

		fmt.Println(l.Translate("Welcome!"))

		return genericresponse.Success(c, l.Translate("System is running"), fiber.StatusOK)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		l := c.Locals("localizer").(localizer.Localizer)

		return genericresponse.Success(c, l.Translate("System is running"), fiber.StatusOK)
	})
}
