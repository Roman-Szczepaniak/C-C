package models

type Monster struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Alignment *string
	Size      string
	Type      string
	CR        string
	MonsterID uint
}
