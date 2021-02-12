package database_seeds

import (
	"errors"
	"github.com/brandon-julio-t/tpa-web-backend/commands"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

func SeedInventory() error {
	users := make([]*models.User, 0)
	if err := facades.UseDB().Find(&users, "id != ?", 1).Error; err != nil {
		return nil
	}

	for _, user := range users {
		gamesCount := faker.Number().NumberInt(1)
		games := make([]*models.Game, 0)
		if err := facades.UseDB().Order("random()").Limit(gamesCount).Find(&games).Error; err != nil {
			return err
		}

		for _, game := range games {
			for i := 0; i < 20; i++ {
				item := new(models.MarketItem)
				if err := facades.UseDB().
					Where("game_id = ?", game.ID).
					Order("random()").
					First(item).Error; err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						continue
					}

					return err
				}

				command := commands.AddToInventoryCommand{
					MarketItem: item,
					User:       user,
				}
				if err := command.Execute(); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
