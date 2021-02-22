package models

type BadgeCardTransaction struct {
	UserID      int64 `gorm:"primaryKey"`
	User        User
	BadgeCardID int64 `gorm:"primaryKey"`
	BadgeCard   BadgeCard
}
