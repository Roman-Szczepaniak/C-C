package models

type Card struct {
	ID              uint     `gorm:"primaryKey" json:"id"`
	Description     *string  `json:"description"`
	CA              uint     `json:"ca"`
	PV              uint     `json:"pv"`
	Speed           string   `json:"speed"`
	Strength        string   `json:"strength"`
	Dexterity       string   `json:"dexterity"`
	Constitution    string   `json:"constitution"`
	Intelligence    string   `json:"intelligence"`
	Wisdom          string   `json:"wisdom"`
	Charisma        string   `json:"charisma"`
	SaveThrows      *string  `json:"save_throws"`
	DamageResist    *string  `json:"damage_resist"`
	DamageImmune    *string  `json:"damage_immune"`
	ConditionImmune *string  `json:"condition_immune"`
	Senses          *string  `json:"senses"`
	Languages       *string  `json:"languages"`
	Actions         []Action `gorm:"many2many:card_actions;constraint:OnDelete:CASCADE;" json:"actions"`
}
