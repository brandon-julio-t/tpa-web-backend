package models

import (
	"time"
)

type Notification struct {
	ID               int64
	UserID           int64
	NotificationUser User `gorm:"foreignKey:UserID"`
	Content          string
	CreatedAt        time.Time
}
