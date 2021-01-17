package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/repositories"
)

func (r *gameResolver) Tags(ctx context.Context, obj *models.Game) ([]*models.GameTag, error) {
	return obj.GameTags, nil
}

func (r *mutationResolver) CreateGame(ctx context.Context, input models.CreateGame) (*models.Game, error) {
	return new(repositories.GameRepository).Create(input)
}

func (r *mutationResolver) UpdateGame(ctx context.Context, input models.UpdateGame) (*models.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteGame(ctx context.Context, id int64) (*models.Game, error) {
	return new(repositories.GameRepository).Delete(id)
}

func (r *queryResolver) GetAllGames(ctx context.Context, page int) ([]*models.Game, error) {
	return new(repositories.GameRepository).GetAll(page)
}

func (r *queryResolver) GetGameByID(ctx context.Context, id int64) (*models.Game, error) {
	return new(repositories.GameRepository).GetById(id)
}

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

type gameResolver struct{ *Resolver }
