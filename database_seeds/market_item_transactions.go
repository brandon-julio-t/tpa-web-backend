package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"syreclabs.com/go/faker"
	"time"
)

func SeedMarketItemTransactions() error {
	if err := randomTransactions("buy"); err != nil {
		return err
	}

	return randomTransactions("sell")
}

func randomTransactions(category string) error {
	for i := 0; i < 100; i++ {
		buyer := new(models.User)
		if err := facades.UseDB().Order("random()").First(buyer).Error; err != nil {
			return err
		}

		seller := new(models.User)
		if err := facades.UseDB().Order("random()").First(seller).Error; err != nil {
			return err
		}

		item := new(models.MarketItem)
		if err := facades.UseDB().Order("random()").First(item).Error; err != nil {
			return err
		}

		now := time.Now()
		before := now.AddDate(0, -1, 0)

		if err := facades.UseDB().Create(&models.MarketItemTransaction{
			Category:    category,
			CreatedAt:   faker.Time().Between(before, now),
			Buyer_:      *buyer,
			MarketItem_: *item,
			Price:       float64(faker.Commerce().Price()),
			Seller_:     *seller,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}
