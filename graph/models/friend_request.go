package models

import (
	"gorm.io/gorm"
	"time"
)

type FriendRequest struct {
	ID        int64
	UserID    int64
	FriendID  int64
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt

	FriendRequestUser   User `gorm:"foreignKey:UserID"`
	FriendRequestFriend User `gorm:"foreignKey:FriendID"`
}
