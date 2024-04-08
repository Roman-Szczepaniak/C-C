package models

type Player struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Level      uint
	Initiative *uint
	HP         uint
	HP_max     uint
	XP         *bool
	AccountID  uint
	PartyID    uint
}
