package models

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	Id           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Country      string    `json:"country" gorm:"not null" validate:"required"`
	State        string    `json:"state" gorm:"not null" validate:"required"`
	City         string    `json:"city" gorm:"not null" validate:"required"`
	District     string    `json:"district" gorm:"not null" validate:"required"`
	ZipCode      string    `json:"zipCode" gorm:"not null" validate:"required"`
	Street       string    `json:"street" gorm:"not null" validate:"required"`
	StreetNumber string    `json:"streetNumber" gorm:"not null" validate:"required"`
	Complement   string    `json:"complement" gorm:"not null" validate:"required"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	CostomerId   uuid.UUID `json:"customer_id" gorm:"type:uuid;index"`
}

type Document struct {
	Id         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Number     int       `json:"number" gorm:"not null" validate:"required"`
	Doc_type   string    `json:"doc_type" gorm:"not null" validate:"required"`
	Country    string    `json:"country" gorm:"default:'BRL'" validate:"required"`
	CostomerId uuid.UUID `json:"customer_id" gorm:"type:uuid;index"`
}

type Costumer struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	Name        string    `json:"name" gorm:"not null" validate:"required"`
	Email       string    `json:"email" gorm:"not null;unique;index" validate:"required,email"`
	Password    string    `json:"-" gorm:"not null" validate:"required"`
	PhoneNumber string    `json:"phoneNumber" gorm:"not null" validate:"required"`

	Address  Address  `json:"address" gorm:"foreignkey:ID;references:id"`
	Document Document `json:"document" gorm:"foreignkey:ID;references:id"`
}
