package models

type Card struct {
	ID           uint `gorm:"primaryKey"`
	Description  *string
	CA           *uint
	PV           *uint
	Speed        *string
	Action       *string
	Strength     string
	Dexterity    string
	Constitution string
	Intelligence string
	Wisdom       string
	Charisma     string
}
