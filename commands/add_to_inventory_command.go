package commands

import (
	"errors"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
)

type AddToInventoryCommand struct {
	MarketItem *models.MarketItem
	User       *models.User
}

func (c AddToInventoryCommand) Execute() error {
	inventory := new(models.Inventory)
	if err := facades.UseDB().
		Where("user_id = ?", c.User.ID).
		Where("market_item_id = ?", c.MarketItem.ID).
		First(inventory).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			facades.UseDB().Create(&models.Inventory{
				User_:       *c.User,
				MarketItem_: *c.MarketItem,
				Quantity:    1,
			})

			return nil
		}

		return err
	}

	inventory.Quantity++
	return facades.UseDB().Save(inventory).Error
}
