package handlers

import (
	"fmt"
	"net/http"
	"payment-ms/processors"
	"payment-ms/services"

	"payment-ms/models"

	"github.com/gorilla/mux"
)

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pagamento Processado!")
	var processor processors.PaymentProcessor = processors.NewDeditCardProcessor()
	result := services.NewPaymentService(&processor).ProcessPayment(&models.Payment{})
	if result != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Pagamento com ID " + id)
	w.WriteHeader(http.StatusOK)
}
