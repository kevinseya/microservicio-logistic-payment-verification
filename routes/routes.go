package routes

import (
	"net/http"
	"payment-verification/controller"
)

func RegisterRoutes(paymentController *controller.PaymentController) {
	http.HandleFunc("/api/payment/validation", paymentController.ValidatePayment)
}
