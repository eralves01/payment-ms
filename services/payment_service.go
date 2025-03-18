package services

import (
	"github.com/eralves01/payment-ms/database"
	"github.com/eralves01/payment-ms/models"
	"github.com/eralves01/payment-ms/processors"
	"github.com/eralves01/payment-ms/repository"
)

type PaymentService struct {
	processor processors.PaymentProcessor
}

func NewPaymentService(processor *processors.PaymentProcessor) *PaymentService {
	return &PaymentService{processor: *processor}
}

func (s *PaymentService) ProcessPayment(payment *models.Payment) error {
	db := database.GetInstance()
	repo := repository.NewPaymentRepository(db)
	err := repo.CreatePayment(payment)
	if err != nil {
		return err
	}

	err = s.processor.ProcessPayment(payment)
	if err != nil {
		return err
	}
	return nil
}
