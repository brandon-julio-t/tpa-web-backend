package models

type Country struct {
	ID   int64  `gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"uniqueIndex"`
}
