package processors

import (
	"payment-ms/models"
)

type PaymentProcessor interface {
	ProcessPayment(payment *models.Payment) error
}
