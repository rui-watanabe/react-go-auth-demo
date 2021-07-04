package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"react-go-auth/database"
	"react-go-auth/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":3000")
}
