package models

type Account struct {
	ID            uint   `gorm:"primaryKey"` // or gorm.Model
	Email         string `gorm:"uniqueIndex"`
	Login         string
	HashPpassword string
}
