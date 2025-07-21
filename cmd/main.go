package main

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/robertvitoriano/rinha-de-backend-golang/internal/handlers"
	"github.com/robertvitoriano/rinha-de-backend-golang/internal/services"
)

func main() {

	app := fiber.New()

	redisClient := redis.NewClient(&redis.Options{})

	paymentHandlers := handlers.NewPaymentHandlers(redisClient)

	go func() {
		pubsub := redisClient.Subscribe("payments")
		ch := pubsub.Channel()

		workerCount := 4
		jobManager := services.NewJobManager(ch, workerCount)

		jobManager.Run()
	}()
	app.Post("/payments", paymentHandlers.ReceivePayment)
	app.Listen(":4444")
}
