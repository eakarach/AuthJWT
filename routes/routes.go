package routes

import (
	"github.com/eakarach/AuthJWT/handles"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", handles.HealthCheck)

	// Auth
	AuthRoutes(api.Group("/auth"))

	// User
	UserRoutes(api.Group("/usr"))
}