package routes

import (
	"net/http"

	"github.com/eralves01/payment-ms/database"
	"github.com/eralves01/payment-ms/handlers"
	"github.com/eralves01/payment-ms/repository"
	"github.com/eralves01/payment-ms/services"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	db := database.NewDatadase().GetInstance()
	repository := repository.NewUserRepository(db)
	service := services.NewUserServices(repository)
	userHandler := handlers.NewUserHandler(service)

	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users/filters", userHandler.GetUsersByFilters).Methods(http.MethodGet)

	// payments := api.PathPrefix("/payments").Subrouter()
	// payments.HandleFunc("", userHandler.ProcessPayment).Methods(http.MethodPost)
	// payments.HandleFunc("/{id}", userHandler.GetPayment).Methods(http.MethodGet)

	users := api.PathPrefix("/users").Subrouter()
	users.HandleFunc("", userHandler.CreateUser).Methods(http.MethodPost)
	users.HandleFunc("", userHandler.GetUsers).Methods(http.MethodGet)
	users.HandleFunc("/{id}", userHandler.GetUserByID).Methods(http.MethodGet)
	users.HandleFunc("/email/{email}", userHandler.GetUserByEmail).Methods(http.MethodGet)
	users.HandleFunc("/filters", userHandler.GetUsersByFilters).Methods(http.MethodGet)

	return router
}
