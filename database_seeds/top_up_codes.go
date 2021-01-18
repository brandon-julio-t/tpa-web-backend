package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
)

func SeedTopUpCodes() error {
	return facades.UseDB().Transaction(func(tx *gorm.DB) error {
		for i := 0; i < 10; i++ {
			tx.Create(&models.TopUpCode{Amount: 10000, Code: facades.UseOTP()})
		}
		return nil
	})
}
