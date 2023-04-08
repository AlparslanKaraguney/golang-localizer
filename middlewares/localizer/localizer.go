package localizer

import (
	"localizer/internal/localizer"

	"github.com/gofiber/fiber/v2"
)

/*
REQUIRED(Any middleware must have this)

For every middleware we need a config.
In config we also need to define a function which allows us to skip the middleware if return true.
By convention it should be named as "Filter" but any other name will work too.
*/
type Config struct {
}

/*
Middleware specific

Our middleware's config default values if not passed
*/
var ConfigDefault = Config{}

/*
Middleware specific

Function for generating default config
*/
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	return cfg

}

/*
	    REQUIRED(Any middleware must have this)

		Our main middleware function used to initialize our middleware.
		By convention we name it "New" but any other name will work too.
*/
func New(config ...Config) fiber.Handler {

	// For setting default config
	_ = configDefault(config...)

	return func(c *fiber.Ctx) error {

		// get header Accept-Language
		locale := c.Get("Accept-Language")

		l := localizer.Get(locale)

		c.Locals("localizer", l)

		return c.Next()
	}
}
