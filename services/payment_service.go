package services

import (
	"payment-ms/database"
	"payment-ms/models"
	"payment-ms/processors"
	"payment-ms/repository"
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
