package main

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/robertvitoriano/rinha-de-backend-golang/internal/clients"
	"github.com/robertvitoriano/rinha-de-backend-golang/internal/handlers"
)

func main() {

	app := fiber.New()

	redisClient := redis.NewClient(&redis.Options{})

	paymentHandlers := handlers.NewPaymentHandlers(redisClient)

	paymentProcessorUrls := []string{
		os.Getenv("PAYMENT_PROCESSOR_URL"),
		os.Getenv("PAYMENT_PROCESSOR_URL"),
	}
	defaultBaseUrlIndex := 0
	fallbackBaseUrlIndex := 1
	paymentProcessorClient := clients.NewPaymentProcessor()
	paymentProcessorClient.SetBaseUrl(paymentProcessorUrls[defaultBaseUrlIndex])
	paymentData := handlers.ReceivePaymentBody{
		CorrelationId: "asdsa",
		Amount:        15,
	}

	err := paymentProcessorClient.SendPayment(paymentData)

	if err != nil {
		paymentProcessorClient.SetBaseUrl(paymentProcessorUrls[fallbackBaseUrlIndex])
		paymentProcessorClient.SendPayment(paymentData)
	}

	app.Post("/payments", paymentHandlers.ReceivePayment)
	app.Listen(":4444")
}
