package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
	"math"
	"strconv"
	"syreclabs.com/go/faker"
	"time"
)

func SeedGamePurchaseTransactions() error {
	return facades.UseDB().Transaction(func(tx *gorm.DB) error {
		for i := 0; i < 20; i++ {
			ids, err := strconv.ParseInt(faker.Number().Between(2, 10), 10, 64)
			if err != nil {
				return err
			}

			header := &models.GamePurchaseTransactionHeader{
				CreatedAt:                           faker.Date().Between(time.Now().AddDate(0, 0, faker.Number().NumberInt(1) * -1), time.Now()),
				GamePurchaseTransactionHeaderUserID: ids,
				GrandTotal:                          float64(faker.Commerce().Price()),
			}

			if err := tx.Create(header).Error; err != nil {
				return err
			}

			count := int(math.Max(float64(faker.Number().NumberInt(1)), 1))
			for j := 0; j < count; j++ {
				ids, err := strconv.ParseInt(faker.Number().Between(1, 20), 10, 64)
				if err != nil {
					return err
				}

				if err := tx.Create(&models.GamePurchaseTransactionDetail{
					GamePurchaseTransactionHeaderID:     header.ID,
					GamePurchaseTransactionDetailGameID: ids,
				}).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}
