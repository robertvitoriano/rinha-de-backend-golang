package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ReceivePaymentBody struct {
	CorrelationId string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
}

func ReceivePayment(c *fiber.Ctx) error {
	var bodyParsed ReceivePaymentBody
	body := c.Body()

	err := json.Unmarshal(body, &bodyParsed)
	if err != nil {
		c.App().ErrorHandler(c, fmt.Errorf("error parsing json"))

	}
	return nil
}
