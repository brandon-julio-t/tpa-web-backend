package models

type GameSlideshow struct {
	ID     int64 `gorm:"primaryKey;autoIncrement:true"`
	GameID int64
	Game   Game
	FileID int64
	File   AssetFile `gorm:"foreignKey:ID"`
}
