package models

type AssetFile struct {
	ID          int64 `gorm:"primaryKey"`
	File        []byte
	ContentType string
}
