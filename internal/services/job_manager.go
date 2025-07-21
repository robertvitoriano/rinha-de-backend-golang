package services

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/robertvitoriano/rinha-de-backend-golang/internal/clients"
	"github.com/robertvitoriano/rinha-de-backend-golang/internal/handlers"
)

type JobManager struct {
	PaymentChannel <-chan *redis.Message
}

func NewJobManager(redisChannel <-chan *redis.Message) *JobManager {

	return &JobManager{
		PaymentChannel: redisChannel,
	}
}

func (jm *JobManager) Run() {
	for msg := range jm.PaymentChannel {
		var paymentData handlers.ReceivePaymentBody
		err := json.Unmarshal([]byte(msg.Payload), &paymentData)
		paymentProcessorClient := clients.NewPaymentProcessor()
		if err == nil {
			go Job(paymentData, paymentProcessorClient)
		}
	}
}
