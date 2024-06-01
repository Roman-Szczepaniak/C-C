package models

type Encounter struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Difficulty string    `json:"difficulty"`
	Type       string    `json:"type"`
	TotalXP    *string   `json:"total_xp"`
	PartyID    uint      `json:"party_id"`
	Party      Party     `gorm:"foreignKey:PartyID"`
	Monsters   []Monster `gorm:"many2many:monster_encounters;constraint:OnDelete:CASCADE;" json:"monsters"`
}
