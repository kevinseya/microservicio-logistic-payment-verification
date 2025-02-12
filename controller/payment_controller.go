package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"payment-verification/config"
	"payment-verification/model"
	"payment-verification/service"
)

type PaymentController struct {
	PaymentService *service.PaymentService
	Config         *config.Config //  We added the configuration here
}

func NewPaymentController(paymentService *service.PaymentService) *PaymentController {
	return &PaymentController{
		PaymentService: paymentService,
		Config:         config.AppConfig, // We use the global config
	}
}

func (pc *PaymentController) ValidatePayment(w http.ResponseWriter, r *http.Request) {
	var payment model.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isValid, err := pc.PaymentService.ValidatePayment(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]bool{"valid": isValid}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	if isValid {
		pc.sendWebhook(payment.PaymentIntent)
	}
}

func (pc *PaymentController) sendWebhook(paymentIntent string) {
	webhookURL := pc.Config.WebhookURL //  Now use the configuration URL
	if webhookURL == "" {
		log.Println("Webhook URL is not set in the configuration.")
		return
	}

	payload := map[string]string{
		"payment_intent": paymentIntent,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON payload: %v", err)
		return
	}

	// Send the POST request to the webhook
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Error sending POST request to webhook: %v", err)
		return
	}
	defer resp.Body.Close()

	// Handle the webhook response
	if resp.StatusCode != http.StatusOK {
		log.Printf("Webhook responded with status: %s", resp.Status)
	} else {
		log.Println("Webhook notification sent successfully.")
	}
}
