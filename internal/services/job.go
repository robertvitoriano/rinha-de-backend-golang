package services

import (
	"os"

	"github.com/robertvitoriano/rinha-de-backend-golang/internal/clients"
	"github.com/robertvitoriano/rinha-de-backend-golang/internal/handlers"
)

func Job(paymentData handlers.ReceivePaymentBody, client *clients.PaymentProcessor) {

	paymentProcessorUrls := []string{
		os.Getenv("PAYMENT_PROCESSOR_URL"),
		os.Getenv("PAYMENT_PROCESSOR_URL"),
	}

	defaultBaseUrlIndex := 0
	fallbackBaseUrlIndex := 1

	client.SetBaseUrl(paymentProcessorUrls[defaultBaseUrlIndex])

	err := client.SendPayment(paymentData)
	if err != nil {
		client.SetBaseUrl(paymentProcessorUrls[fallbackBaseUrlIndex])
		client.SendPayment(paymentData)
	}
}
