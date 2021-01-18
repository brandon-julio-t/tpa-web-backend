package models

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"gorm.io/gorm"
)

type TopUpCode struct {
	ID     int64 `gorm:"primaryKey"`
	Code   string `gorm:"uniqueIndex"`
	Amount float64
}

func (u *TopUpCode) BeforeCreate(tx *gorm.DB) error {
	u.Code = facades.UseOTP()
	return nil
}
