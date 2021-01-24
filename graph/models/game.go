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
	Discount           float64
	GameTags           []*GameTag `gorm:"many2many:game_tag_mappings;"`
	GenreID            int64
	Genre              GameGenre
	HoursPlayed        float64
	IsInappropriate    bool
	Price              float64
	GameGameReviews    []*GameReview `gorm:"foreignKey:GameReviewGameID"`
	GameSlideshows     []*GameSlideshow
	SystemRequirements string
	Title              string
}
