package models

type GameTag struct {
	ID    int64  `gorm:"primaryKey;autoIncrement:true"`
	Games []Game `gorm:"many2many:game_tag_mappings;"`
	Name  string `gorm:"uniqueIndex"`
}
