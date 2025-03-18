package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type PaymentMethodType string

const (
	CreditCard PaymentMethodType = "credit_card"
	DebitCard  PaymentMethodType = "debit_card"
	Pix        PaymentMethodType = "pix"
	Ticket     PaymentMethodType = "ticket"
)

type PaymentMethod struct {
	ID        uuid.UUID         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	UserID    uuid.UUID         `json:"user_id" gorm:"type:uuid;not null;index"`
	Type      PaymentMethodType `json:"type" gorm:"type:varchar(50);not null;check:type IN ('credit_card', 'debit_card', 'pix', 'ticket')"`
	Data      json.RawMessage   `json:"data" gorm:"type:jsonb;not null"`
	CreatedAt time.Time         `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt *time.Time        `json:"updated_at" gorm:"type:timestamp"`
}
