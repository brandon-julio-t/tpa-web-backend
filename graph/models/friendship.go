package models

type Friendship struct {
	UserID   int64 `gorm:"primaryKey"`
	User     User
	FriendID int64 `gorm:"primaryKey"`
	Friend   User
}
