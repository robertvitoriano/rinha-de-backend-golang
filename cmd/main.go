package main

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/robertvitoriano/rinha-de-backend-golang/internal/handlers"
)

func main() {

	app := fiber.New()

	redisClient := redis.NewClient(&redis.Options{})

	paymentHandlers := handlers.NewPaymentHandlers(redisClient)

	app.Post("/payments", paymentHandlers.ReceivePayment)

	app.Listen(":4444")
}
