package main

import (
	baseRouter "localizer/cmd/baserouter"
	"localizer/cmd/config"
	customLogger "localizer/helper/logger"
	"localizer/middlewares/localizer"
	"localizer/middlewares/security"
	"time"

	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	customLogger.Infof("Starting the application...")

	//  Initialize the app and add middleware
	app := fiber.New(config.FiberConfig)

	// Custom middleware
	app.Use(security.New())
	app.Use(localizer.New())

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: time.DateTime,
	}))

	baseRouter.InitializeRouters(app)

	// Start server
	go func() {

		err := app.Listen(":8000")
		if err != nil {
			customLogger.Fatalf("Error while starting the server", err.Error())
		}
	}()

	GracefulShutdown(app)

}

func GracefulShutdown(app *fiber.App) {

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	customLogger.Infof("Received terminate,graceful shutdown", sig)

	app.Shutdown()
}
