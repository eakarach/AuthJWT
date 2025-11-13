package routes

import (
	"github.com/eakarach/AuthJWT/handles"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(routes fiber.Router)  {
	routes.Get("/", handles.Auth)
	routes.Post("/login", handles.Login)
}