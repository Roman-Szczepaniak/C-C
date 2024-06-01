package models

type History struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID"`
}
