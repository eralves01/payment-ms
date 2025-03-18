package processors

import (
	"github.com/eralves01/payment-ms/models"
)

type PaymentProcessor interface {
	ProcessPayment(payment *models.Payment) error
}
