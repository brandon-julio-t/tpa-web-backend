package models

type UnsuspendRequest struct {
	ID     int64 `gorm:"primaryKey"`
	UserID int64
	User   User
}
