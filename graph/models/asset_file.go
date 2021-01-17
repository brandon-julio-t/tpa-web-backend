package models

type AssetFile struct {
	ID          int64 `gorm:"primaryKey;autoIncrement:true"`
	File        []byte
	ContentType string
}
