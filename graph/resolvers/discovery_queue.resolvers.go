package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func (r *discoveryQueueResolver) NewReleases(ctx context.Context, obj *models.DiscoveryQueue) ([]*models.Game, error) {
	obj.DiscoveryQueueNewReleased = make([]*models.Game, 0)
	return obj.DiscoveryQueueNewReleased, facades.UseDB().
		Order("created_at desc").
		Limit(10).
		Find(&obj.DiscoveryQueueNewReleased).
		Error
}

func (r *queryResolver) DiscoverQueue(ctx context.Context) (*models.DiscoveryQueue, error) {
	return new(models.DiscoveryQueue), nil
}

// DiscoveryQueue returns generated.DiscoveryQueueResolver implementation.
func (r *Resolver) DiscoveryQueue() generated.DiscoveryQueueResolver {
	return &discoveryQueueResolver{r}
}

type discoveryQueueResolver struct{ *Resolver }
