package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
	"math"
	"strconv"
	"strings"
	"syreclabs.com/go/faker"
	"time"
)

func SeedGameGiftTransactions() error {
	return facades.UseDB().Transaction(func(tx *gorm.DB) error {
		for i := 0; i < 20; i++ {
			userIds, err := strconv.ParseInt(faker.Number().Between(2, 10), 10, 64)
			if err != nil {
				return err
			}

			friendIds, err := strconv.ParseInt(faker.Number().Between(2, 10), 10, 64)
			if err != nil {
				return err
			}

			header := &models.GameGiftTransactionHeader{
				CreatedAt:                         faker.Date().Between(time.Now().AddDate(0, 0, faker.Number().NumberInt(1) * -1), time.Now()),
				GameGiftTransactionHeaderUserID:   userIds,
				GameGiftTransactionHeaderFriendID: friendIds,
				Message:                           strings.Join(faker.Lorem().Words(faker.Number().NumberInt(1)), " "),
				Sentiment:                         strings.Join(faker.Lorem().Words(faker.Number().NumberInt(1)), " "),
				Signature:                         strings.Join(faker.Lorem().Words(faker.Number().NumberInt(1)), " "),
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

				if err := tx.Create(&models.GameGiftTransactionDetail{
					GameGiftTransactionHeaderID:     header.ID,
					GameGiftTransactionDetailGameID: ids,
				}).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}
