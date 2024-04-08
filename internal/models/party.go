package models

import "time"

type Party struct {
	ID         uint       `gorm:"primaryKey"`
	Created_at *time.Time // * ==> can be null
	Updated_at *time.Time
}
