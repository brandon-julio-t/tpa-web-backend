package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math"
	"strings"
	"time"

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

func (r *queryResolver) CommunityRecommended(ctx context.Context) ([]*models.Game, error) {
	games := make([]*models.Game, 0)
	now := time.Now()
	aWeekAgo := now.AddDate(0, 0, -7)

	return games, facades.UseDB().
		Joins(
			"join (?) as recommended_games on games.id = recommended_games.id",
			facades.UseDB().
				Select("g.id, count(game_reviews.id) as recommendations_count").
				Model(new(models.GameReview)).
				Joins("join games g on game_reviews.game_review_game_id = g.id").
				Where("is_recommended = ?", true).
				Where("game_reviews.created_at >= ?", aWeekAgo).
				Where("game_reviews.created_at <= ?", now).
				Group("g.id"),
		).
		Joins(
			"join (?) as unrecommended_games on games.id = unrecommended_games.id",
			facades.UseDB().
				Select("g.id, count(game_reviews.id) as recommendations_count").
				Model(new(models.GameReview)).
				Joins("join games g on game_reviews.game_review_game_id = g.id").
				Where("is_recommended = ?", false).
				Where("game_reviews.created_at >= ?", aWeekAgo).
				Where("game_reviews.created_at <= ?", now).
				Group("g.id"),
		).
		Order("recommended_games.recommendations_count - unrecommended_games.recommendations_count desc").
		Limit(12).
		Find(&games).
		Error
}

func (r *queryResolver) FeaturedAndRecommendedGames(ctx context.Context) ([]*models.Game, error) {
	return new(repositories.GameRepository).GetFeaturedAndRecommendedGames()
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

func (r *queryResolver) SearchGames(ctx context.Context, page int64, keyword string) (*models.GamePagination, error) {
	games := make([]*models.Game, 0)
	total := new(int64)

	if err := facades.UseDB().Debug().
		Model(new(models.Game)).
		Count(total).
		Scopes(facades.UsePagination(int(page), 10)).
		Where(
			"games.id in (?)",
			facades.UseDB().
				Select("gtm.game_id").
				Model(new(models.GameTag)).
				Joins("join game_tag_mappings gtm on game_tags.id = gtm.game_tag_id").
				Where("lower(game_tags.name) like ?", "%"+strings.ToLower(keyword)+"%"),
		).
		Or("lower(games.title) like ?", "%"+strings.ToLower(keyword)+"%").
		Find(&games).
		Error; err != nil {
		return nil, err
	}

	return &models.GamePagination{
		Data:       games,
		TotalPages: int64(math.Ceil(float64(*total) / float64(10))),
	}, nil
}

func (r *queryResolver) SpecialOffersGame(ctx context.Context) ([]*models.Game, error) {
	games := make([]*models.Game, 0)
	return games, facades.UseDB().
		Where("discount > 0").
		Order("discount desc").
		Limit(24).
		Find(&games).
		Error
}

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// GameSlideshow returns generated.GameSlideshowResolver implementation.
func (r *Resolver) GameSlideshow() generated.GameSlideshowResolver { return &gameSlideshowResolver{r} }

type gameResolver struct{ *Resolver }
type gameSlideshowResolver struct{ *Resolver }
