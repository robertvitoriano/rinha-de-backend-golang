package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robertvitoriano/rinha-de-backend-golang/internal/handlers"
)

func main() {

	app := fiber.New()

	app.Post("/payments", handlers.ReceivePayment)

	app.Listen(":4444")
}
