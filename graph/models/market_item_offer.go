package models

import "time"

type MarketItemOffer struct {
	ID           int64
	Category     string
	CreatedAt    time.Time
	MarketItemID int64
	MarketItem_  MarketItem `gorm:"foreignKey:MarketItemID"`
	Price        float64
	Quantity     int64
	UserID       int64
	User_        User `gorm:"foreignKey:UserID"`
}
