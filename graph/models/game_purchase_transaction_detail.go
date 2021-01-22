package models

type GamePurchaseTransactionDetail struct {
	GamePurchaseTransactionHeaderID     int64
	GamePurchaseTransactionDetailGameID int64
	GamePurchaseTransactionDetailGame   Game
}
