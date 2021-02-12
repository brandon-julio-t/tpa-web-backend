package models

type MarketItemTransaction struct {
	ID           int64
	BuyerID      int64
	Buyer_       User `gorm:"foreignKey:BuyerID"`
	MarketItemID int64
	MarketItem_  MarketItem `gorm:"foreignKey:MarketItemID"`
	SellerID     int64
	Seller_      User `gorm:"foreignKey:SellerID"`
}
