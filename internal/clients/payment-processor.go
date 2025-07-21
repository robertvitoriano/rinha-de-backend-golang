package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/robertvitoriano/rinha-de-backend-golang/internal/handlers"
)

type PaymentProcessor struct {
	BaseUrl string
}

func (pp *PaymentProcessor) SendPayment(paymentData handlers.ReceivePaymentBody) error {
	jsonData, err := json.Marshal(paymentData)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("%v/payments", pp.BaseUrl),
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (pp *PaymentProcessor) SetBaseUrl(url string) {
	pp.BaseUrl = url
}
func NewPaymentProcessor() *PaymentProcessor {

	return &PaymentProcessor{
		BaseUrl: "",
	}

}
