package models

type Card struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	Description  *string `json:"description"`
	CA           *uint   `json:"ca"`
	PV           *uint   `json:"pv"`
	Speed        *string `json:"speed"`
	Action       *string `json:"action"`
	Strength     string  `json:"strength"`
	Dexterity    string  `json:"dexterity"`
	Constitution string  `json:"constitution"`
	Intelligence string  `json:"intelligence"`
	Wisdom       string  `json:"wisdom"`
	Charisma     string  `json:"charisma"`
}
