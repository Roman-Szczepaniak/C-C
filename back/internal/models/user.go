package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex" json:"email"`
	FirstName    string    `json:"firstname" validate:"required,min=2,max=100"`
	LastName     string    `json:"lastname" validate:"required,min=2,max=100"`
	HashPassword string    `json:"-"`
	Token        *string   `json:"token"`
	RefreshToken *string   `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
