package models

type Player struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Level      uint   `json:"level"`
	Initiative *uint  `json:"initiative"`
	HP         uint   `json:"hp"`
	HP_max     uint   `json:"hp_max"`
	XP         *bool  `json:"xp"`
	AccountID  uint   `json:"account_id"`
	PartyID    uint   `json:"party_id"`
}
