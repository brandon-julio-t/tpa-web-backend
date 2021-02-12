package models

type MarketItem struct {
	ID          int64
	Description string
	GameID      int64
	Game_       Game `gorm:"foreignKey:GameID"`
	ImageID     int64
	ImageRef    AssetFile `gorm:"foreignKey:ImageID"`
	Name        string
}
