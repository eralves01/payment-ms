package models

import (
	"time"

	"github.com/google/uuid"
)

type UserType string

const (
	Client   UserType = "client"
	Merchant UserType = "merchant"
)

type User struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name         string     `json:"name" gorm:"type:varchar(255);not null"`
	Email        string     `json:"email" gorm:"type:varchar(255);unique;not null"`
	PasswordHash string     `json:"password_hash" gorm:"type:text;not null"`
	UserType     UserType   `json:"user_type" gorm:"type:varchar(50);not null;check:user_type IN ('client', 'merchant')"`
	CreatedAt    time.Time  `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"type:timestamp"`
}
