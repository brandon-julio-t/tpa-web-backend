package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func (r *queryResolver) GetAllGameTags(ctx context.Context) ([]*models.GameTag, error) {
	var tags []*models.GameTag
	if err := facades.UseDB().Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
