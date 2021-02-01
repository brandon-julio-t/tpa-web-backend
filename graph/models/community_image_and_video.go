package models

import "time"

type CommunityImageAndVideo struct {
	ID          int64
	CreatedAt   time.Time
	Description string
	FileID      int64
	Name        string
	UserID      int64

	File_ AssetFile `gorm:"foreignKey:FileID"`
	User_ User      `gorm:"foreignKey:UserID"`
}
