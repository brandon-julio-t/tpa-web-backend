package models

import "time"

type GameReviewComment struct {
	ID           int64
	CreatedAt    time.Time
	Body         string
	GameReviewID int64
	UserID       int64

	GameReview_ GameReview `gorm:"foreignKey:GameReviewID"`
	User_       User       `gorm:"foreignKey:UserID"`
}
