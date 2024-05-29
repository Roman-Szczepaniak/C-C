package models

type Monster struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Alignment   string `json:"alignment"`
	Size        string `json:"size"`
	Type        string `json:"type"`
	Environment string `json:"environment"`
	Status      string `json:"status"`
	CR          string `json:"cr"`
	CardID      uint   `json:"card_id"`
}
