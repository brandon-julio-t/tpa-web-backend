package models

import "time"

type MarketItemTransaction struct {
	ID           int64
	BuyerID      int64
	Buyer_       User `gorm:"foreignKey:BuyerID"`
	Category     string
	CreatedAt    time.Time
	MarketItemID int64
	MarketItem_  MarketItem `gorm:"foreignKey:MarketItemID"`
	Price        float64
	SellerID     int64
	Seller_      User `gorm:"foreignKey:SellerID"`
}
