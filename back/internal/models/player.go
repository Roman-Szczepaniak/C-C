package models

type Player struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Level      uint   `json:"level"`
	Initiative *uint  `json:"initiative"`
	HP         uint   `json:"hp"`
	HPMax      uint   `json:"hp_max"`
	XP         *bool  `json:"xp"`
	UserID     uint   `json:"user_id"`
	User       User   `gorm:"foreignKey:UserID""`
	PartyID    uint   `json:"party_id"`
	Party      Party  `gorm:"foreignKey:PartyID"`
}
