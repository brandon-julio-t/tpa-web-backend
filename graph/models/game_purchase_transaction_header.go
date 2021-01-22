package models

import "time"

type GamePurchaseTransactionHeader struct {
	ID                                   int64 `gorm:"primaryKey"`
	GamePurchaseTransactionHeaderUserID  int64
	GamePurchaseTransactionHeaderUser    User
	GamePurchaseTransactionHeaderDetails []*GamePurchaseTransactionDetail
	GrandTotal                           float64
	CreatedAt                            time.Time
}
