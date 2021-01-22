package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
)

func (r *mutationResolver) AddToWishlist(ctx context.Context, gameID int64) (*models.Game, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	var game models.Game
	if err := facades.UseDB().First(&game, gameID).Error; err != nil {
		return nil, err
	}

	return &game, facades.UseDB().Model(user).Association("UserWishlist").Append(&game)
}

func (r *mutationResolver) RemoveFromWishlist(ctx context.Context, gameID int64) (*models.Game, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	var game models.Game
	if err := facades.UseDB().First(&game, gameID).Error; err != nil {
		return nil, err
	}

	return &game, facades.UseDB().Model(user).Association("UserWishlist").Delete(&game)
}
