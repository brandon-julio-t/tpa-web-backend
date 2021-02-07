package models

type PointItem struct {
	ID       int64
	Name     string
	Category string
	Price    int
	ImageID  int64
	Image_   AssetFile `gorm:"foreignKey:ImageID"`
}
