package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/repositories"
)

func (r *queryResolver) AllCountries(ctx context.Context) ([]*models.Country, error) {
	return new(repositories.CountryRepository).GetAll()
}
