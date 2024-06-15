package router

import (
	"github.com/dmdhrumilmistry/gofilecheck/service/check"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// check routes
	checkHandler := check.NewCheckHandler()
	checkHandler.SetupRoutes(api)
}
