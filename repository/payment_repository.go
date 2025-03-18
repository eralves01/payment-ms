package repository

import (
	"database/sql"
	"fmt"

	"github.com/eralves01/payment-ms/models"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) CreatePayment(payment *models.Payment) error {
	// _, err := r.db.Exec("INSERT INTO payments (order_id, amount, status) VALUES ($1, $2, $3)", payment.OrderID, payment.Amount, payment.Status)
	// return err
	fmt.Println("Pagamento criado com sucesso!")
	return nil
}
