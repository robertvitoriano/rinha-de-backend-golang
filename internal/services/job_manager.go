package services

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/robertvitoriano/rinha-de-backend-golang/internal/clients"
	"github.com/robertvitoriano/rinha-de-backend-golang/internal/handlers"
)

type JobManager struct {
	PaymentChannel <-chan *redis.Message
	WorkerCount    int
}

func NewJobManager(redisChannel <-chan *redis.Message, workerCount int) *JobManager {
	return &JobManager{
		PaymentChannel: redisChannel,
		WorkerCount:    workerCount,
	}
}

func (jm *JobManager) Run() {
	jobs := make(chan handlers.ReceivePaymentBody)

	for i := 0; i < jm.WorkerCount; i++ {
		go func() {
			for paymentData := range jobs {
				paymentProcessorClient := clients.NewPaymentProcessor()
				Job(paymentData, paymentProcessorClient)
			}
		}()
	}

	for msg := range jm.PaymentChannel {
		var paymentData handlers.ReceivePaymentBody
		err := json.Unmarshal([]byte(msg.Payload), &paymentData)
		if err == nil {
			jobs <- paymentData
		}
	}
	close(jobs)
}
