package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/eralves01/payment-ms/configs"
	"github.com/eralves01/payment-ms/database"
	"github.com/eralves01/payment-ms/routes"
)

func main() {
	log := configs.NewLogger("payment-ms")
	configs.LoadConfig()
	database.RunMigrations()
	database.GetInstance()

	router := routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Infof("Servidor rodando na porta %s...\n", port)
	log.Error(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
