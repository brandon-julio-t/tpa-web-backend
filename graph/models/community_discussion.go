package models

import "time"

type CommunityDiscussion struct {
	ID        int64
	Body      string
	CreatedAt time.Time
	GameID    int64
	Title     string
	UserID    int64

	Game_ Game `gorm:"primaryKey:GameID"`
	User_ User `gorm:"primaryKey:UserID"`
}
