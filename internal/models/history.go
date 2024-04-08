package models

type History struct {
	ID        uint `gorm:"primaryKey"`
	AccountID uint
}
