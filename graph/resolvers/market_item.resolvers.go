package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"math"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
)

func (r *marketItemResolver) Image(ctx context.Context, obj *models.MarketItem) (*models.AssetFile, error) {
	return &obj.ImageRef, facades.UseDB().Preload("ImageRef").First(obj).Error
}

func (r *marketItemResolver) StartingPrice(ctx context.Context, obj *models.MarketItem) (float64, error) {
	offer := new(models.MarketItemOffer)

	if err := facades.UseDB().Where("market_item_id = ?", obj.ID).Order("price").First(offer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}

		return 0, err
	}

	return offer.Price, nil
}

func (r *marketItemResolver) TransactionsCount(ctx context.Context, obj *models.MarketItem) (int64, error) {
	count := new(int64)
	return *count, facades.UseDB().
		Model(obj).
		Joins("join market_item_transactions mit on market_items.id = mit.market_item_id").
		Where("market_items.id = ?", obj.ID).
		Group("market_items.id, market_items.game_id, market_items.image_id, market_items.name, market_items.user_id").
		Count(count).
		Error
}

func (r *queryResolver) MarketItems(ctx context.Context, page int64) (*models.MarketItemPagination, error) {
	items := make([]*models.MarketItem, 0)
	count := new(int64)
	perPage := 10

	if err := facades.UseDB().
		Model(new(models.MarketItem)).
		Joins("full join market_item_transactions mit on market_items.id = mit.market_item_id").
		Group("market_items.id, market_items.game_id, market_items.image_id, market_items.name, market_items.user_id").
		Count(count).
		Scopes(facades.UsePagination(int(page), perPage)).
		Order("count(mit.id) desc").
		Find(&items).
		Error; err != nil {
		return nil, err
	}

	return &models.MarketItemPagination{
		Data:       items,
		TotalPages: int64(math.Ceil(float64(*count) / float64(perPage))),
	}, nil
}

func (r *userResolver) GamesByOwnedMarketItems(ctx context.Context, obj *models.User) ([]*models.Game, error) {
	games := make([]*models.Game, 0)
	if err := facades.UseDB().
		Joins("join market_items mi on games.id = mi.game_id").
		Joins("join inventories i on mi.id = i.market_item_id").
		Where("user_id = ?", obj.ID).
		Distinct("games.*").
		Find(&games).
		Error; err != nil {
		return nil, err
	}

	return games, nil
}

func (r *userResolver) MarketItemsByGame(ctx context.Context, obj *models.User, page int64, gameID int64) (*models.MarketItemPagination, error) {
	items := make([]*models.MarketItem, 0)
	count := new(int64)
	perPage := 10

	inventories := make([]*models.Inventory, 0)
	if err := facades.UseDB().
		Model(new(models.Inventory)).
		Joins("join market_items mi on inventories.market_item_id = mi.id").
		Where("game_id = ?", gameID).
		Where("user_id = ?", obj.ID).
		Count(count).
		Scopes(facades.UsePagination(int(page), perPage)).
		Find(&inventories).
		Error; err != nil {
		return nil, err
	}

	for _, inventory := range inventories {
		if err := facades.UseDB().Preload("MarketItem_").First(inventory).Error; err != nil {
			return nil, err
		}

		items = append(items, &inventory.MarketItem_)
	}

	return &models.MarketItemPagination{
		Data:       items,
		TotalPages: int64(math.Ceil(float64(*count) / float64(perPage))),
	}, nil
}

// MarketItem returns generated.MarketItemResolver implementation.
func (r *Resolver) MarketItem() generated.MarketItemResolver { return &marketItemResolver{r} }

type marketItemResolver struct{ *Resolver }
