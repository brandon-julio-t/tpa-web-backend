package models

type GameGenre struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
}
