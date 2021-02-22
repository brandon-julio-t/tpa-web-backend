package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/brandon-julio-t/tpa-web-backend/facades"

	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func (r *badgeResolver) Game(ctx context.Context, obj *models.Badge) (*models.Game, error) {
	game := new(models.Game)
	return game, facades.UseDB().First(game, obj.GameID).Error
}

func (r *badgeResolver) IsOwned(ctx context.Context, obj *models.Badge) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *badgeCardResolver) Badge(ctx context.Context, obj *models.BadgeCard) (*models.Badge, error) {
	return &obj.Badge, facades.UseDB().Preload("Badge").First(obj).Error
}

func (r *badgeCardResolver) IsOwned(ctx context.Context, obj *models.BadgeCard) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *gameResolver) Badges(ctx context.Context, obj *models.Game) ([]*models.Badge, error) {
	badges := make([]*models.Badge, 0)
	return badges, facades.UseDB().Where("game_id = ?", obj.ID).Find(&badges).Error
}

// Badge returns generated.BadgeResolver implementation.
func (r *Resolver) Badge() generated.BadgeResolver { return &badgeResolver{r} }

// BadgeCard returns generated.BadgeCardResolver implementation.
func (r *Resolver) BadgeCard() generated.BadgeCardResolver { return &badgeCardResolver{r} }

type badgeResolver struct{ *Resolver }
type badgeCardResolver struct{ *Resolver }
