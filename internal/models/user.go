package models

type User struct {
	ID           uint   `gorm:"primaryKey"` // or gorm.Model
	Email        string `gorm:"uniqueIndex"`
	Login        string
	HashPassword string
}
