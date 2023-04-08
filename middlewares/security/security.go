package security

import (
	"localizer/helper/genericresponse"

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

		// X-XSS-Protection
		c.Response().Header.Add("X-XSS-Protection", "1; mode=block")

		// HTTP Strict Transport Security
		c.Response().Header.Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		// X-Frame-Options
		c.Response().Header.Add("X-Frame-Options", "SAMEORIGIN")

		// X-Content-Type-Options
		c.Response().Header.Add("X-Content-Type-Options", "nosniff")

		// Content Security Policy
		c.Response().Header.Add("Content-Security-Policy", "default-src 'self';")

		// X-Permitted-Cross-Domain-Policies
		c.Response().Header.Add("X-Permitted-Cross-Domain-Policies", "none")

		// Referrer-Policy
		c.Response().Header.Add("Referrer-Policy", "no-referrer")

		// Feature-Policy
		c.Response().Header.Add("Feature-Policy", "microphone 'none'; camera 'none'")

		if c.Method() == "OPTIONS" {
			return genericresponse.Error(c, 204, "")
		}

		return c.Next()
	}
}
