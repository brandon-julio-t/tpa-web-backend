package models

import "time"

type FriendRequest struct {
	ID int64
	UserID int64
	FriendID int64
	CreatedAt time.Time

	FriendRequestUser User `gorm:"foreignKey:UserID"`
	FriendRequestFriend User `gorm:"foreignKey:FriendID"`
}
