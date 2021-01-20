package models

import "time"

type PrivateMessage struct {
	ID         int64 `gorm:"primaryKey"`
	Text       string
	SenderID   int64
	Sender     User
	ReceiverID int64
	Receiver   User
	CreatedAt  time.Time
}
