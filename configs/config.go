package configs

import (
	"github.com/joho/godotenv"
)

// LoadConfig carrega variáveis de ambiente do arquivo .env
func LoadConfig() {
	log := NewLogger("payment-ms")
	err := godotenv.Load()
	if err != nil {
		log.Info(".env file not found, loading system variables...")
	}
	log.Info("Environment variables were loaded from the .env file...")
}
