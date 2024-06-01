package models

type Action struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Cards       []Card `gorm:"many2many:card_actions;constraint:OnDelete:CASCADE;" json:"cards"`
}
