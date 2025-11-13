package main

import (
	"github.com/eakarach/AuthJWT/routes"
	"github.com/eakarach/AuthJWT/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// connect db
	database.ConnectDB()
	
	// fiber instance
	app := fiber.New()

	// routes 
	routes.SetupRoutes(app)

	// app listening at PORT: 3000
	app.Listen(":3000")
}