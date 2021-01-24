package models

type GameReviewVote struct {
	GameReviewVoteGameReviewID int64 `gorm:"primaryKey"`
	GameReviewVoteGameReview   GameReview
	GameReviewVoteUserID       int64 `gorm:"primaryKey"`
	GameReviewVoteUser         User
	IsUpVote                   bool
}
