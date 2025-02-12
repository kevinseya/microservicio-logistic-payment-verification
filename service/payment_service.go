package service

import (
	"payment-verification/model"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type PaymentService struct {
	StripeSecretKey string
}

func NewPaymentService(secretKey string) *PaymentService {
	return &PaymentService{StripeSecretKey: secretKey}
}

func (s *PaymentService) ValidatePayment(payment *model.Payment) (bool, error) {
	stripe.Key = s.StripeSecretKey
	paymentIntent, err := paymentintent.Get(payment.PaymentIntent, nil)
	if err != nil {
		return false, err
	}
	return paymentIntent.Status == "succeeded", nil
}
