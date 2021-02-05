package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math"
	"strings"

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

func (r *queryResolver) AllGames(ctx context.Context) ([]*models.Game, error) {
	games := make([]*models.Game, 0)
	return games, facades.UseDB().Find(&games).Error
}

func (r *queryResolver) CommunityRecommended(ctx context.Context) ([]*models.Game, error) {
	games := make([]*models.Game, 0)
	return games, new(repositories.GameRepository).
		GetCommunityRecommends().
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

func (r *queryResolver) NewAndTrending(ctx context.Context) ([]*models.Game, error) {
	games := make([]*models.Game, 0)
	return games, facades.UseDB().
		Order("created_at desc").
		Limit(10).
		Find(&games).
		Error
}

func (r *queryResolver) SearchGames(ctx context.Context, page int64, keyword string, price int64, genres []int64, category string) (*models.GamePagination, error) {
	games := make([]*models.Game, 0)
	total := new(int64)
	db := facades.UseDB()
	repo := new(repositories.GameRepository)

	switch category {
	case "community_recommends":
		db = repo.GetCommunityRecommends()
	case "special_offers":
		db = repo.GetSpecialOffers()
	case "top_sellers":
		db = repo.GetTopSellers()
	}

	query := db.Model(new(models.Game)).
		Scopes(facades.UsePagination(int(page), 10))

	if price > 0 {
		query = query.Where("games.price <= ?", price)
	}

	if len(genres) > 0 {
		query = query.Where(
			"games.id in (?)",
			facades.UseDB().
				Model(new(models.GameTag)).
				Select("gtm.game_id").
				Joins("join game_tag_mappings gtm on game_tags.id = gtm.game_tag_id").
				Where("game_tags.id in ?", genres),
		)
	}

	if err := query.
		Where("lower(games.title) like ?", "%"+strings.ToLower(keyword)+"%").
		Count(total).
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
	return games, new(repositories.GameRepository).
		GetSpecialOffers().
		Limit(24).
		Find(&games).
		Error
}

func (r *queryResolver) Specials(ctx context.Context) ([]*models.Game, error) {
	games := make([]*models.Game, 0)
	return games, facades.UseDB().
		Where("discount >= ?", 0.5).
		Order("discount desc").
		Limit(10).
		Find(&games).
		Error
}

func (r *queryResolver) TopSellers(ctx context.Context) ([]*models.Game, error) {
	games := make([]*models.Game, 0)
	return games, new(repositories.GameRepository).
		GetTopSellers().
		Limit(10).
		Find(&games).
		Error
}

func (r *userResolver) Games(ctx context.Context, obj *models.User) ([]*models.Game, error) {
	games := make([]*models.Game, 0)

	if err := facades.UseDB().
		Where(
			"games.id in (?)",
			facades.UseDB().
				Model(new(models.GamePurchaseTransactionDetail)).
				Select("game_purchase_transaction_detail_game_id").
				Joins("join game_purchase_transaction_headers gpth on game_purchase_transaction_details.game_purchase_transaction_header_id = gpth.id").
				Where("game_purchase_transaction_header_user_id = ?", obj.ID),
		).
		Or(
			"games.id in (?)",
			facades.UseDB().
				Model(new(models.GameGiftTransactionDetail)).
				Select("game_gift_transaction_detail_game_id").
				Joins("join game_gift_transaction_headers ggth on game_gift_transaction_details.game_gift_transaction_header_id = ggth.id").
				Where("game_gift_transaction_header_user_id = ?", obj.ID),
		).
		Find(&games).
		Error; err != nil {
		return nil, err
	}

	return games, nil
}

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// GameSlideshow returns generated.GameSlideshowResolver implementation.
func (r *Resolver) GameSlideshow() generated.GameSlideshowResolver { return &gameSlideshowResolver{r} }

type gameResolver struct{ *Resolver }
type gameSlideshowResolver struct{ *Resolver }
