package models

import (
	"time"
)

type GameReview struct {
	ID                       int64 `gorm:"primaryKey"`
	CreatedAt                time.Time
	GameReviewGameID         int64
	GameReviewGame           Game `gorm:"foreignKey:GameReviewGameID"`
	GameReviewUserID         int64
	GameReviewUser           User              `gorm:"foreignKey:GameReviewUserID"`
	GameReviewGameReviewVote []*GameReviewVote `gorm:"foreignKey:GameReviewVoteGameReviewID"`
	Content                  string
	IsRecommended            bool
}
