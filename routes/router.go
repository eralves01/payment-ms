package routes

import (
	"net/http"
	"payment-ms/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Criando um subrouter para a versão da API
	api := router.PathPrefix("/api/v1").Subrouter()

	// Rotas de pagamentos
	payments := api.PathPrefix("/payments").Subrouter()
	payments.HandleFunc("", handlers.ProcessPayment).Methods(http.MethodPost)
	payments.HandleFunc("/{id}", handlers.GetPayment).Methods(http.MethodGet)

	// Rotas de usuários
	users := api.PathPrefix("/users").Subrouter()
	users.HandleFunc("", handlers.CreateUser).Methods(http.MethodPost)
	users.HandleFunc("/{id}", handlers.GetUserByID).Methods(http.MethodGet)
	users.HandleFunc("/email/{email}", handlers.GetUserByEmail).Methods(http.MethodGet)

	return router
}
