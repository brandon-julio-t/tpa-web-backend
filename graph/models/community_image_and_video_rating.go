package models

import "time"

type CommunityImageAndVideoRating struct {
	CommunityImageAndVideoID int64 `gorm:"primaryKey"`
	CreatedAt                  time.Time
	UserID                     int64 `gorm:"primaryKey"`
	IsLike                     bool

	CommunityImageAndVideo_ CommunityImageAndVideo `gorm:"foreignKey:CommunityImageAndVideoID"`
	User_                     User                   `gorm:"foreignKey:UserID"`
}
