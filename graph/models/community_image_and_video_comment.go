package models

import "time"

type CommunityImageAndVideoComment struct {
	ID                       int64
	Body                     string
	CommunityImageAndVideoID int64
	CreatedAt                time.Time
	UserID                   int64

	User_                   User                   `gorm:"foreignKey:UserID"`
	CommunityImageAndVideo_ CommunityImageAndVideo `gorm:"foreignKey:CommunityImageAndVideoID"`
}
