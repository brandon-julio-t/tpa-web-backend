package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"github.com/brandon-julio-t/tpa-web-backend/repositories"
)

func (r *gameResolver) IsInCart(ctx context.Context, obj *models.Game) (bool, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return false, nil
	}

	return facades.UseDB().
		Model(&user).
		Where("game_id = ?", obj.ID).
		Association("UserCart").
		Count() > 0, nil
}

func (r *gameResolver) IsInWishlist(ctx context.Context, obj *models.Game) (bool, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return false, nil
	}

	return facades.UseDB().
		Model(&user).
		Where("game_id = ?", obj.ID).
		Association("UserWishlist").
		Count() > 0, nil
}

func (r *gameResolver) Slideshows(ctx context.Context, obj *models.Game) ([]*models.GameSlideshow, error) {
	return obj.GameSlideshows, facades.UseDB().
		Preload("GameSlideshows.GameSlideshowFile").
		First(obj).
		Error
}

func (r *gameResolver) Tags(ctx context.Context, obj *models.Game) ([]*models.GameTag, error) {
	return obj.GameTags, nil
}

func (r *gameSlideshowResolver) File(ctx context.Context, obj *models.GameSlideshow) (*models.AssetFile, error) {
	return &obj.GameSlideshowFile, facades.UseDB().First(obj).Error
}

func (r *mutationResolver) CreateGame(ctx context.Context, input models.CreateGame) (*models.Game, error) {
	return new(repositories.GameRepository).Create(input)
}

func (r *mutationResolver) UpdateGame(ctx context.Context, input models.UpdateGame) (*models.Game, error) {
	return new(repositories.GameRepository).Update(input)
}

func (r *mutationResolver) DeleteGame(ctx context.Context, id int64) (*models.Game, error) {
	return new(repositories.GameRepository).Delete(id)
}

func (r *queryResolver) Games(ctx context.Context, page int64) (*models.GamePagination, error) {
	return new(repositories.GameRepository).GetAll(int(page))
}

func (r *queryResolver) Genres(ctx context.Context) ([]*models.GameGenre, error) {
	var genres []*models.GameGenre
	return genres, facades.UseDB().Find(&genres).Error
}

func (r *queryResolver) GetGameByID(ctx context.Context, id int64) (*models.Game, error) {
	return new(repositories.GameRepository).GetById(id)
}

func (r *queryResolver) FeaturedAndRecommendedGames(ctx context.Context) ([]*models.Game, error) {
	return new(repositories.GameRepository).GetFeaturedAndRecommendedGames()
}

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// GameSlideshow returns generated.GameSlideshowResolver implementation.
func (r *Resolver) GameSlideshow() generated.GameSlideshowResolver { return &gameSlideshowResolver{r} }

type gameResolver struct{ *Resolver }
type gameSlideshowResolver struct{ *Resolver }
