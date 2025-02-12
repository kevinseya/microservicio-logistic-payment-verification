package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	StripeSecretKey string
	WebhookURL      string // 🔹 Cambia `webhookURL` a `WebhookURL` para exportarlo
}

// ✅ Definir variable global
var AppConfig *Config

// ✅ Guardar la configuración en `AppConfig`
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = &Config{
		StripeSecretKey: os.Getenv("STRIPE_SECRET_KEY"),
		WebhookURL:      os.Getenv("NOTIFICATION_WEBHOOK_PAYMENT_URL"),
	}

	log.Println("Configuration loaded successfully.")
}
