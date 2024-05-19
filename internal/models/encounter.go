package models

type Encounter struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Difficulty string  `json:"difficulty"`
	Type       string  `json:"type"`
	Total_XP   *string `json:"total_xp"`
	PartyID    uint    `json:"party_id"`
}
