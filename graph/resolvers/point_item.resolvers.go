package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/repositories"
)

func (r *mutationResolver) PurchasePointItem(ctx context.Context, id int64) (*models.PointItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PointItemProfileBackgrounds(ctx context.Context) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetProfileBackgrounds()
}

func (r *queryResolver) PointItemAvatarBorders(ctx context.Context) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetAvatarBorders()
}

func (r *queryResolver) PointItemAnimatedAvatars(ctx context.Context) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetAnimatedAvatars()
}

func (r *queryResolver) PointItemChatStickers(ctx context.Context) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetChatStickers()
}

func (r *queryResolver) PointItemMiniProfileBackgrounds(ctx context.Context) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetMiniProfileBackgrounds()
}
