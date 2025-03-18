package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"payment-ms/configs"
	"payment-ms/database"
	"payment-ms/routes"
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
