package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

type ReceivePaymentBody struct {
	CorrelationId string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
}

type PaymentHandlers struct {
	RedisClient *redis.Client
}

func NewPaymentHandlers(client *redis.Client) *PaymentHandlers {
	return &PaymentHandlers{
		RedisClient: client,
	}
}

type PaymentInfo struct {
	CorrelationId string    `json:"correlationId"`
	Amount        float64   `json:"amount"`
	Time          time.Time `json:"time"`
}

func (ph *PaymentHandlers) ReceivePayment(c *fiber.Ctx) error {
	var bodyParsed ReceivePaymentBody
	body := c.Body()

	err := json.Unmarshal(body, &bodyParsed)
	if err != nil {
		c.App().ErrorHandler(c, fmt.Errorf("error parsing json"))
	}
	if bodyParsed.Amount == 0 || bodyParsed.CorrelationId == "" {
		c.App().ErrorHandler(c, fmt.Errorf("amount and correlation id must be provided"))
		return fmt.Errorf("amount and correlation id must be provided")
	}

	paymentInfo := PaymentInfo{
		CorrelationId: bodyParsed.CorrelationId,
		Amount:        bodyParsed.Amount,
		Time:          time.Now(),
	}

	paymentInfoData, err := json.Marshal(paymentInfo)

	if err != nil {
		c.App().ErrorHandler(c, fmt.Errorf("error storing payment information"))

	}

	paymentKey := fmt.Sprintf("payment:%v", bodyParsed.CorrelationId)

	ph.RedisClient.Set(paymentKey, paymentInfoData, time.Hour)

	fmt.Println("Payment successfully stored")

	//TODO MAKE REQUEST TO PAYMENT PROCCESSOR AND HANDLE FAILURE
	// resp, err := http.Get("https://api.example.com/data")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	return nil
}
