package models

type Encounter struct {
	ID         uint `gorm:"primaryKey"`
	Difficulty string
	Type       string
	Total_XP   *string
	PartyID    uint
}
