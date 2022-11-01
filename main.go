package main

import (
	"fiberTest/database"
	"fiberTest/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

)

func main() {

	database.Connect()
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Listen(":3000")
}