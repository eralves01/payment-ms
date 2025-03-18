package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

const (
	Pending  PaymentStatus = "pending"
	Approved PaymentStatus = "approved"
	Canceled PaymentStatus = "canceled"
	Refunded PaymentStatus = "refunded"
)

type Payment struct {
	ID              uuid.UUID     `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	UserID          *uuid.UUID    `json:"user_id" gorm:"type:uuid;index;constraint:OnDelete:SET NULL"`
	MerchantID      *uuid.UUID    `json:"merchant_id" gorm:"type:uuid;index;constraint:OnDelete:SET NULL"`
	PaymentMethodID *uuid.UUID    `json:"payment_method_id" gorm:"type:uuid;index;constraint:OnDelete:SET NULL"`
	Amount          float64       `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency        string        `json:"currency" gorm:"type:varchar(3);not null;default:'USD'"`
	Status          PaymentStatus `json:"status" gorm:"type:varchar(50);not null;default:'pending';check:status IN ('pending', 'approved', 'canceled', 'refunded')"`
	CreatedAt       time.Time     `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt       *time.Time    `json:"updated_at" gorm:"type:timestamp"`
}
