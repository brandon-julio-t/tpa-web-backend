package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int64 `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
