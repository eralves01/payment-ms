package routes

import (
	"net/http"

	"github.com/eralves01/payment-ms/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users/filters", handlers.GetUsersByFilters).Methods(http.MethodGet)

	payments := api.PathPrefix("/payments").Subrouter()
	payments.HandleFunc("", handlers.ProcessPayment).Methods(http.MethodPost)
	payments.HandleFunc("/{id}", handlers.GetPayment).Methods(http.MethodGet)

	users := api.PathPrefix("/users").Subrouter()
	users.HandleFunc("", handlers.CreateUser).Methods(http.MethodPost)
	users.HandleFunc("", handlers.GetUsers).Methods(http.MethodGet)
	users.HandleFunc("/{id}", handlers.GetUserByID).Methods(http.MethodGet)
	users.HandleFunc("/email/{email}", handlers.GetUserByEmail).Methods(http.MethodGet)
	users.HandleFunc("/filters", handlers.GetUsersByFilters).Methods(http.MethodGet)

	return router
}
