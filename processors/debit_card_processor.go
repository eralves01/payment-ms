package processors

import (
	"fmt"

	"github.com/eralves01/payment-ms/models"
)

type DeditCardProcessor struct {
}

func NewDeditCardProcessor() *DeditCardProcessor {
	return &DeditCardProcessor{}
}

func (c *DeditCardProcessor) ProcessPayment(payment *models.Payment) error {
	fmt.Println("Pagamento com cartão de débito processado!")
	return nil
}
