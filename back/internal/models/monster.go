package models

type Monster struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string  `json:"name"`
	Alignment *string `json:"alignment"`
	Size      string  `json:"size"`
	Type      string  `json:"typz"`
	CR        string  `json:"cr"`
	MonsterID uint    `json:"monster_id"`
}
