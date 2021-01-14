package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func (r *queryResolver) AllCountries(ctx context.Context) ([]*models.Country, error) {
	var countries []*models.Country
	if err := facades.UseDB().Find(&countries).Error; err != nil {
		return nil, err
	}
	return countries, nil
}
