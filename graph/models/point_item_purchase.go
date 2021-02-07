package models

type PointItemPurchase struct {
	UserID      int64     `gorm:"primaryKey"`
	PointItemID int64     `gorm:"primaryKey"`
	User_       User      `gorm:"foreignKey:UserID"`
	PointItem_  PointItem `gorm:"foreignKey:PointItemID"`
}
