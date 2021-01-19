package models

import "time"

type ProfileComment struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64
	User      User
	ProfileID int64
	Profile   User
	Comment   string
	CreatedAt time.Time
}
