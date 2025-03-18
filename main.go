package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eralves01/payment-ms/configs"
	"github.com/eralves01/payment-ms/database"
	"github.com/eralves01/payment-ms/routes"
)

func main() {
	configs.LoadConfig()
	database.RunMigrations()
	database.GetInstance()

	router := routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Servidor rodando na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
