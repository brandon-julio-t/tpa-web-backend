package models

type Inventory struct {
	UserID       int64      `gorm:"primaryKey"`
	User_        User       `gorm:"foreignKey:UserID"`
	MarketItemID int64      `gorm:"primaryKey"`
	MarketItem_  MarketItem `gorm:"foreignKey:MarketItemID"`
	Quantity     int64
}
