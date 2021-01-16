package models

type GameTag struct {
	BaseModel
	Games []Game `gorm:"many2many:game_tag_mappings;"`
	Name  string `gorm:"uniqueIndex"`
}
