package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func SeedMarketItemTransactions() error {
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

		if err := facades.UseDB().Create(&models.MarketItemTransaction{
			Buyer_:      *buyer,
			MarketItem_: *item,
			Seller_:     *seller,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}
