package models

type BadgeTransaction struct {
	BadgeID int64 `gorm:"primaryKey"`
	Badge   Badge
	UserID  int64 `gorm:"primaryKey"`
	User    User
}
