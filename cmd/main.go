package main

import (
	"encoding/json"
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

	go func() {
		pubsub := redisClient.Subscribe("payments")
		ch := pubsub.Channel()
		for msg := range ch {
			var paymentData handlers.ReceivePaymentBody
			err := json.Unmarshal([]byte(msg.Payload), &paymentData)
			if err == nil {
				err = paymentProcessorClient.SendPayment(paymentData)
				if err != nil {
					paymentProcessorClient.SetBaseUrl(paymentProcessorUrls[fallbackBaseUrlIndex])
					paymentProcessorClient.SendPayment(paymentData)
				}
			}
		}
	}()

	app.Post("/payments", paymentHandlers.ReceivePayment)
	app.Listen(":4444")
}
