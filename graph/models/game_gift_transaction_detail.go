package models

type GameGiftTransactionDetail struct {
	GameGiftTransactionHeaderID     int64
	GameGiftTransactionDetailGameID int64
	GameGiftTransactionDetailGame   Game
}
