package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	Name      string    `json:"name" gorm:"not null" validate:"required"`
	Email     string    `json:"email" gorm:"not null;unique;index" validate:"required,email"`
	Password  string    `json:"password" gorm:"not null" validate:"required"`
}
