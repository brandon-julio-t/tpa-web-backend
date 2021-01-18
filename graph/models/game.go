package models

import (
	"time"
)

type Game struct {
	ID                 int64 `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt          time.Time
	Title              string
	Description        string
	Price              float64
	BannerID           int64
	Banner             AssetFile `gorm:"foreignKey:ID"`
	Slideshows         []*GameSlideshow
	GameTags           []*GameTag `gorm:"many2many:game_tag_mappings;"`
	SystemRequirements string
}
