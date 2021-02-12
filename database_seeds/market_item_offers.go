package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"syreclabs.com/go/faker"
	"time"
)

func SeedMarketItemOffers() error {
	if err := randomOffers("buy"); err != nil {
		return err
	}

	return randomOffers("sell")
}

func randomOffers(category string) error {
	now := time.Now()
	before := now.AddDate(0, 0, faker.Number().NumberInt(1) * -1)

	for i := 0; i < 20; i++ {
		item := new(models.MarketItem)
		if err := facades.UseDB().Order("random()").First(item).Error; err != nil {
			return err
		}

		user := new(models.User)
		if err := facades.UseDB().Order("random()").First(user).Error; err != nil {
			return err
		}

		if err := facades.UseDB().Create(&models.MarketItemOffer{
			Category:     category,
			CreatedAt:    faker.Time().Between(before, now),
			MarketItem_:  *item,
			Price:        float64(faker.Commerce().Price()),
			User_:        *user,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}
