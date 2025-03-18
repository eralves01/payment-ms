package processors

import (
	"fmt"

	"github.com/eralves01/payment-ms/models"
)

type CreditCardProcessor struct {
}

func NewCreditCardProcessor() *CreditCardProcessor {
	return &CreditCardProcessor{}
}

func (c *CreditCardProcessor) ProcessPayment(payment *models.Payment) error {
	fmt.Println("Pagamento com cartão de crédito processado!")
	return nil
}
