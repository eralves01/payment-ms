package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/eralves01/payment-ms/models"
	"github.com/eralves01/payment-ms/services"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = nil

	service := services.NewUserServices()
	result := service.CreateUser(&user)
	if result != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	service := services.NewUserServices()
	users, err := service.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	service := services.NewUserServices()
	user, err := service.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	email := vars["email"]

	fmt.Println("Email: " + email)

	service := services.NewUserServices()
	user, err := service.GetUserByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUsersByFilters(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	service := services.NewUserServices()
	users, err := service.GetUsersByFilters(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
