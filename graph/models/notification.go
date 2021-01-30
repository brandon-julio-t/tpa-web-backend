package models

import (
	"time"
)

// TODO: new item in users inventory

type Notification struct {
	ID               int64
	UserID           int64
	NotificationUser User `gorm:"foreignKey:UserID"`
	Content          string
	CreatedAt        time.Time
}
