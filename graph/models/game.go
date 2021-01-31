package models

import (
	"time"
)

type Game struct {
	ID                 int64 `gorm:"primaryKey;autoIncrement:true"`
	Banner             AssetFile
	BannerID           int64
	CreatedAt          time.Time
	Description        string
	Developer          string
	Discount           float64
	GameGameReviews    []*GameReview `gorm:"foreignKey:GameReviewGameID"`
	GameSlideshows     []*GameSlideshow
	GameTags           []*GameTag `gorm:"many2many:game_tag_mappings;"`
	Genre              GameGenre
	GenreID            int64
	HoursPlayed        float64
	IsInappropriate    bool
	Price              float64
	Publisher          string
	SystemRequirements string
	Title              string
}
