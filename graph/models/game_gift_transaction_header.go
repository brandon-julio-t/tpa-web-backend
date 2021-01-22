package models

import "time"

type GameGiftTransactionHeader struct {
	ID                                int64 `gorm:"primaryKey"`
	GameGiftTransactionHeaderUserID   int64
	GameGiftTransactionHeaderUser     User
	GameGiftTransactionHeaderFriendID int64
	GameGiftTransactionHeaderFriend   User
	Message                           string
	Sentiment                         string
	Signature                         string
	GameGiftTransactionHeaderDetails  []*GameGiftTransactionDetail
	CreatedAt                         time.Time
}
