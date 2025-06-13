package main

import (
	"log"
	"net/http"
	"payment-verification/config"
	"payment-verification/controller"
	"payment-verification/routes"
	"payment-verification/service"
)

func main() {
	config.LoadConfig()
	paymentService := service.NewPaymentService(config.AppConfig.StripeSecretKey)
	paymentController := controller.NewPaymentController(paymentService)

	routes.RegisterRoutes(paymentController)

	log.Println("Server is running on port 8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
