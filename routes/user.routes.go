package routes

import (
	"github.com/eakarach/AuthJWT/handles"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(routes fiber.Router)  {
	routes.Get("/", handles.GetAllUser)
	routes.Get("/:usrId", handles.GetUserProfile)
	routes.Post("/", handles.CreateUser)
	routes.Delete("/:usrId", handles.DeleteUser)
	routes.Put("/:usrId", handles.UpdateUserProfile)
}