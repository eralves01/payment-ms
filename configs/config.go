package configs

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadConfig carrega variáveis de ambiente do arquivo .env
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, carregando variáveis do sistema...")
	}
}
