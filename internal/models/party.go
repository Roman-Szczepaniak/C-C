package models

import "time"

type Party struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Created_at *time.Time `json:"created_at"` // * ==> can be null
	Updated_at *time.Time `json:"updated_at"`
}
