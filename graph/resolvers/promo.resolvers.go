package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gorm.io/gorm"
	"math"
	"time"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func (r *mutationResolver) CreatePromo(ctx context.Context, discount float64, endAt time.Time, gameID int64) (*models.Promo, error) {
	game := new(models.Game)
	if err := facades.UseDB().First(game, gameID).Error; err != nil {
		return nil, err
	}

	promo := &models.Promo{
		Discount: 0,
		EndAt:    time.Time{},
		Game_:    *game,
	}

	return promo, facades.UseDB().Transaction(func(tx *gorm.DB) error {
		game.Discount = discount / 100

		if err := tx.Save(game).Error; err != nil {
			return err
		}

		if err := tx.Create(promo).Error; err != nil {
			return err
		}

		wishlists := make([]*models.Wishlist, 0)
		if err := tx.Where("game_id = ?", game.ID).
			Find(&wishlists).
			Error; err != nil {
			return err
		}

		for _, wishlist := range wishlists {
			user := wishlist.User_
			game := wishlist.Game_
			if err := facades.UseMail().SendText(
				"A game in your wishlist, "+game.Title+", is on sale!",
				"Sale Game",
				"SaleGame",
				user.Email,
			); err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *mutationResolver) UpdatePromo(ctx context.Context, id int64, discount float64, endAt time.Time) (*models.Promo, error) {
	var promo models.Promo
	if err := facades.UseDB().First(&promo, id).Error; err != nil {
		return nil, err
	}
	promo.Discount = discount
	promo.EndAt = endAt
	return &promo, facades.UseDB().Save(&promo).Error
}

func (r *mutationResolver) DeletePromo(ctx context.Context, id int64) (*models.Promo, error) {
	var promo models.Promo
	if err := facades.UseDB().First(&promo, id).Error; err != nil {
		return nil, err
	}
	return &promo, facades.UseDB().Delete(&promo).Error
}

func (r *queryResolver) Promos(ctx context.Context, page int64) (*models.PromoPagination, error) {
	var promos []*models.Promo
	var count int64
	pageSize := 5

	if err := facades.UseDB().
		Model(&models.Promo{}).
		Count(&count).
		Scopes(facades.UsePagination(int(page), pageSize)).
		Find(&promos).Error; err != nil {
		return nil, err
	}

	return &models.PromoPagination{
		Data:       promos,
		TotalPages: int64(math.Ceil(float64(count) / float64(pageSize))),
	}, nil
}

func (r *queryResolver) Promo(ctx context.Context, id int64) (*models.Promo, error) {
	var promo models.Promo
	return &promo, facades.UseDB().First(&promo, id).Error
}
